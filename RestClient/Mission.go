package RestClient

import (
	"HNExtensionConverter/RestClient/models"
	"fmt"
)

func (c *APIClient) GetMission(missionId int) (*models.Mission, error) {
	path := fmt.Sprintf("missions/%d", missionId)

	req, err := c.Build("GET", path, nil)

	if err != nil {
		return nil, err
	}

	mission := new(models.Mission)
	_, err = c.Do(req, mission)

	return mission, err
}

func (c *APIClient) UpdateMission(missionId string, mission models.Mission) (*models.Mission, error) {
	path := fmt.Sprintf("missions/%s", missionId)

	method := "PUT"
	if (missionId == "new") {
		method = "POST"
	}

	req, err := c.Build(method, path, mission)

	if err != nil {
		return nil, err
	}

	ret := new(models.Mission)
	_, err = c.Do(req, ret)

	return ret, err
}

func (c *APIClient) CreateMission(mission models.Mission) (*models.Mission, error) {
	return c.UpdateMission("new", mission)
}

func (c *APIClient) GetEmail(emailId int) (*models.Email, error) {
	path := fmt.Sprintf("missions/email/%d", emailId)

	req, err := c.Build("GET", path, nil)

	if err != nil {
		return nil, err
	}

	email := new(models.Email)
	_, err = c.Do(req, email)

	return email, err
}

func (c *APIClient) CreateEmail(email *models.Email) (*models.Email, error) {
	req, err := c.Build("POST", "missions/email/new", email)

	if err != nil {
		return nil, err
	}

	newMail := new(models.Email)
	_, err = c.Do(req, newMail)

	return newMail, err
}

func (c *APIClient) UpdateEmail(emailId int, email *models.Email) (*models.Email, error) {
	req, err := c.Build("PUT", fmt.Sprintf("missions/email/%d", emailId), email)

	if err != nil {
		return nil, err
	}

	newMail := new(models.Email)
	_, err = c.Do(req, newMail)

	return newMail, err
}

func (c *APIClient) DeleteEmail(emailId int) error {
	req, err := c.Build("DELETE", fmt.Sprintf("missions/email/%d", emailId), nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

func (c *APIClient) GetBoardPost(postId int) (*models.BoardPost, error) {
	req, err := c.Build("GET", fmt.Sprintf("missions/postings/%d", postId), nil)

	if err != nil {
		return nil, err
	}

	post := new(models.BoardPost)
	_, err = c.Do(req, post)

	return post, err
}

func (c *APIClient) CreateBoardPost(post *models.BoardPost) (*models.BoardPost, error) {
	req, err := c.Build("POST", "missions/postings/new", post)

	if err != nil {
		return nil, err
	}

	newPost := new(models.BoardPost)
	_, err = c.Do(req, newPost)

	return newPost, err
}

func (c *APIClient) UpdateBoardPost(postId int, post *models.BoardPost) (*models.BoardPost, error) {
	req, err := c.Build("PUT", fmt.Sprintf("missions/postings/%d", postId), post)

	if err != nil {
		return nil, err
	}

	newPost := new(models.BoardPost)
	_, err = c.Do(req, newPost)

	return newPost, err
}

func (c *APIClient) DeleteBoardPost(postId int) error {
	req, err := c.Build("DELETE", fmt.Sprintf("missions/postings/%d", postId), nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

func (c *APIClient) LinkMissionPost(postId int, missionId int) error {
	req, err := c.Build("GET", fmt.Sprintf("missions/linkPosting?mission=%d&posting=%d", missionId, postId), nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

func (c *APIClient) LinkMissionEmail(emailId int, missionId int) error {
	req, err := c.Build("GET", fmt.Sprintf("missions/linkEmail?mission=%d&email=%d", missionId, emailId), nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

func (c *APIClient) CreateMissionBranch(parent, child int) (*models.BranchLink, error) {
	brnch := models.BranchLink{
		MissionOne: parent,
		MissionTwo: child,
	}

	req, err := c.Build("POST", "missions/branch", brnch)

	if err != nil {
		return nil, err
	}

	newBranch := new(models.BranchLink)
	_, err = c.Do(req, newBranch)

	return newBranch, err
}

func (c *APIClient) SetNextMission(parent, child int) error {
	path := fmt.Sprintf("missions/nextMission?m1=%d&m2=%d", parent, child)

	req, err := c.Build("GET", path, nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

func (c *APIClient) CreateAttachment(attachment *models.EmailAttachment) (*models.EmailAttachment, error) {
	req, err := c.Build("POST", "missions/email/attachment/new", attachment)

	if err != nil {
		return nil, err
	}

	newAttachment := new(models.EmailAttachment)
	_, err = c.Do(req, newAttachment)

	return newAttachment, err
}

func (c *APIClient) DeleteAttachment(attachmentId int) error {
	req, err := c.Build("DELETE", fmt.Sprintf("missions/email/attachment/%d", attachmentId), nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

func (c *APIClient) LinkEmailAttachment(emailId, attachmentId int) error {
	req, err := c.Build("GET", fmt.Sprintf("missions/email/attachment/link/%d/%d", emailId, attachmentId), nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

func (c *APIClient) RemoveEmailAttachment(emailId, attachmentId int) error {
	req, err := c.Build("DELETE", fmt.Sprintf("missions/email/attachment/link/%d/%d", emailId, attachmentId), nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

func (c *APIClient) GetGoalTypes() (*[]models.TypeDefinition, error) {
	req, err := c.Build("GET", "missions/goals/types/list", nil)

	if err != nil {
		return nil, err
	}

	types := make([]models.TypeDefinition, 0)
	_, err = c.Do(req, types)

	return &types, err
}

//func (c *APIClient) CreateMissionGoal(goal )