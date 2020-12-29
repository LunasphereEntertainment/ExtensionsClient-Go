package converters

import (
	dModels "HNExtensionConverter/ExtensionData"
	"HNExtensionConverter/RestClient"
	rModels "HNExtensionConverter/RestClient/models"
	"fmt"
	"os"
	"strings"
)

type MissionConverter struct {
	client RestClient.APIClient
}

var (
	missionMap     map[int]rModels.Mission
	missionPathMap map[string]int

	pendingMissionLinks    map[int]string
	pendingMissionBranches map[int][]string

	goalTypeMap map[string]int

	functionMap map[string]int
)

func NewMissionConverter(client RestClient.APIClient) MissionConverter {
	missionMap = make(map[int]rModels.Mission)
	missionPathMap = make(map[string]int)

	pendingMissionLinks = make(map[int]string)
	pendingMissionBranches = make(map[int][]string)

	goalTypeMap = make(map[string]int)
	gTypes, err := client.GetGoalTypes()
	if err != nil {
		panic(err)
	}
	for _, gType := range *gTypes {
		goalTypeMap[gType.Name] = gType.TypeID
	}

	functionMap = make(map[string]int)
	funcs, err := client.GetMissionFunctions()
	if err != nil {
		panic(err)
	}
	for _, mFunc := range *funcs {
		functionMap[mFunc.DataName] = mFunc.FunctionID
	}

	return MissionConverter{client}
}

// TODO: Process Mission Goals
// DONE: Process Next Mission(s)
// DONE: Process Mission Branches
// DONE: Process Email Attachments

func resolveFunction(fullFunction string) (string, int) {
	for funcName, funcId := range functionMap {
		if strings.HasPrefix(fullFunction, funcName) {
			return strings.ReplaceAll(fullFunction, funcName+":", ""), funcId
		}
	}
	return fullFunction, 0
}

func (c *MissionConverter) ConvertMission(path string, mission *dModels.Mission) rModels.Mission {

	goals := make([]string, 0)

	mStartMeta, mStartID := resolveFunction(mission.MissionStart.Function)
	mEndMeta, mEndId := resolveFunction(mission.MissionEnd.Function)

	rMission := rModels.Mission{
		//MissionID:                      0,
		//ExtensionID: c.client.ExtensionID,
		ID:                             mission.ID,
		ActiveCheck:                    mission.ActiveCheck,
		ShouldIgnoreSenderVerification: mission.ShouldIgnoreSenderVerification,
		MissionStart: rModels.MissionActions{
			FunctionID: mStartID,
			Meta:       mStartMeta,
			Value:      mission.MissionStart.Value,
			Suppress:   mission.MissionStart.Suppress,
		},
		MissionEnd: rModels.MissionActions{
			FunctionID: mEndId,
			Meta:       mEndMeta,
			Value:      mission.MissionEnd.Value,
			Suppress:   mission.MissionEnd.Suppress,
		},
		NextMission: 0,
		Goals:       goals,
	}

	newMission, err := c.client.CreateMission(rMission)
	if err != nil {
		panic(err)
	}

	rEmail := rModels.Email{
		Sender:  mission.Email.Sender,
		Subject: mission.Email.Subject,
		Body:    mission.Email.Body,
	}

	newEmail, err := c.client.CreateEmail(&rEmail)
	if err != nil {
		panic(err)
	}
	err = c.client.LinkMissionEmail(newEmail.EmailID, newMission.MissionID)
	if err != nil {
		panic(err)
	}

	c.processEmailAttachments(newEmail.EmailID, mission.Email.Attachments)

	if &mission.Posting != nil && len(mission.Posting.Title) > 0 {
		rPosting := rModels.BoardPost{
			//PostingID:    0,
			Title:        mission.Posting.Title,
			Reqs:         mission.Posting.Requirements,
			RequiredRank: mission.Posting.RequiredRank,
			Content:      mission.Posting.Body,
		}

		newPost, err := c.client.CreateBoardPost(&rPosting)

		if err != nil {
			panic(err)
		}

		c.client.LinkMissionPost(newPost.PostingID, newMission.MissionID)
	}

	// Update maps for reference later
	missionMap[newMission.MissionID] = *newMission
	missionPathMap[path] = newMission.MissionID

	if &mission.NextMission != nil && mission.NextMission.MissionPath != "NONE" {
		rMission.IsSilent = mission.NextMission.IsSilent

		// Prep the link ready for creation at the end
		pendingMissionLinks[newMission.MissionID] = mission.NextMission.MissionPath
	}

	if &mission.BranchMissions != nil && len(mission.BranchMissions.Branch) > 0 {

		pendingMissionBranches[newMission.MissionID] = make([]string, len(mission.BranchMissions.Branch))

		for i, branch := range mission.BranchMissions.Branch {
			pendingMissionBranches[newMission.MissionID][i] = branch.MissionPath
		}
	}

	return *newMission
}

func (c *MissionConverter) processMissionGoals(missionId int, goals dModels.MissionGoalCollection) {
	for _, goal := range goals.Goal {
		rGoal := &rModels.MissionGoal{
			TypeID: goalTypeMap[goal.Type],
		}

		fmt.Sprintf("%+v\n", rGoal)
	}
}

func (c *MissionConverter) processEmailAttachments(emailId int, attachments dModels.EmailAttachmentCollection) {
	for _, note := range attachments.Notes {
		attachment := &rModels.EmailAttachment{
			TypeID:  1,
			Title:   note.Title,
			Content: note.Content,
		}

		c.createAttachment(emailId, attachment)
	}

	for _, link := range attachments.Links {
		attachment := &rModels.EmailAttachment{
			TypeID:   2,
			Computer: link.Comp,
		}

		c.createAttachment(emailId, attachment)
	}

	for _, account := range attachments.Accounts {
		attachment := &rModels.EmailAttachment{
			TypeID:   3,
			Computer: account.Comp,
			Username: account.User,
			Password: account.Pass,
		}

		c.createAttachment(emailId, attachment)
	}
}

func (c *MissionConverter) createAttachment(emailId int, attachment *rModels.EmailAttachment) {
	newAttachment, err := c.client.CreateAttachment(attachment)
	if err != nil {
		panic(err)
	}

	err = c.client.LinkEmailAttachment(emailId, newAttachment.AttachmentID)
	if err != nil {
		panic(err)
	}
}

func (c *MissionConverter) CreateMissionLinks() {
	// Process the next mission links...
	for parentId, path := range pendingMissionLinks {
		childMissionId, ok := missionPathMap[normalisePath(path)]

		if ok {
			err := c.client.SetNextMission(parentId, childMissionId)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("Could not create NextMission link between Mission ID %d and PATH %s \n", parentId, path)
		}
	}

	// Process the branch mission links...
	for parentId, paths := range pendingMissionBranches {
		for _, path := range paths {
			childMissionId, ok := missionPathMap[normalisePath(path)]

			if ok {
				_, err := c.client.CreateMissionBranch(parentId, childMissionId)
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Printf("Could not create Mission Branch between Mission ID %d AND PATH %s \n", parentId, path)
			}
		}
	}
}

func normalisePath(path string) string {
	return strings.ReplaceAll(path, "/", string(os.PathSeparator))
}
