package models

type Mission struct {
	MissionID int `json:"missionId,omitempty"`
	ExtensionID int `json:"extensionId,omitempty"`

	ID string `json:"id"`
	ActiveCheck bool `json:"activeCheck"`
	ShouldIgnoreSenderVerification bool `json:"shouldIgnoreSenderVerification"`

	MissionStart MissionActions `json:"missionStart,omitempty"`
	MissionEnd MissionActions `json:"missionEnd,omitempty"`

	NextMission int `json:"nextMission,omitempty"`

	IsSilent bool `json:"isSilent,omitempty"`

	EmailID int `json:"emailId,omitempty"`
	PostingID int `json:"postingId,omitempty"`

	Goals []string `json:"goals"`
}

type Email struct {
	EmailID int `json:"emailId,omitempty"`
	Sender string `json:"sender"`
	Subject string `json:"subject"`
	Body string `json:"body"`
}

type BoardPost struct {
	PostingID int `json:"postingId,omitempty"`
	Title string `json:"title"`
	Reqs string `json:"reqs"`
	RequiredRank int `json:"requiredRank"`
	Content string `json:"content"`
}

type BranchLink struct {
	MissionOne int `json:"missionOne"`
	MissionTwo int `json:"missionTwo"`
}

type EmailAttachment struct {
	AttachmentID int `json:"attachmentId,omitempty"`
	TypeID int `json:"typeId"`

	Title string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`

	Computer string `json:"comp,omitempty"`

	Username string `json:"user,omitempty"`
	Password string `json:"pass,omitempty"`
}

type MissionGoal struct {
	TypeID int `json:"typeId"`

	TargetNodeID int `json:"targetNodeId"`

	CaseSensitive bool `json:"caseSensitive"`
	Degree string `json:"degree"`
	Delay float32 `json:"delay"`
	File string `json:"file"`
	GPA string `json:"gpa"`
	Keyword string `json:"keyword"`
	MailServer string `json:"mailServer"`
	Owner string `json:"owner"`
	Path string `json:"path"`
	Recipient string `json:"recipient"`
	Removal bool `json:"removal"`
	Subject string `json:"subject"`
	Target string `json:"target"`

	//TypeText string `json:"typeText"`
	Uni string `json:"uni"`
}

type MissionFunctionDef struct {
	FunctionID int `json:"functionId"`
	DisplayName string `json:"funcDisplayName"`
	DataName string `json:"funcName"`
}

type MissionActions struct {
	FunctionID int `json:"functionId,omitempty"`
	Meta string `json:"meta,omitempty"`
	Value int `json:"value,omitempty"`
	Suppress bool `json:"suppress,omitempty"`
}