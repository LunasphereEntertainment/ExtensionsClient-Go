package ExtensionData

type Mission struct {
	ID string `xml:"id,attr"`
	ActiveCheck bool `xml:"activeCheck,attr"`
	ShouldIgnoreSenderVerification bool `xml:"shouldIgnoreSenderVerification,attr"`

	Goals MissionGoalCollection `xml:"goals"`

	MissionStart MissionActions `xml:"missionStart,omitempty"`
	MissionEnd MissionActions `xml:"missionEnd,omitempty"`

	NextMission MissionLink `xml:"nextMission"`

	BranchMissions BranchCollection `xml:"branchMissions"`

	Posting BoardPost `xml:"posting"`
	Email MissionEmail `xml:"email"`
}

type MissionGoalCollection struct {
	Goal []MissionGoal `xml:"goal"`
}

type MissionGoal struct {
	Type string `xml:"type,attr"`

	Target string `xml:"target,attr,omitempty"`
	File string `xml:"file,attr,omitempty"`
	Path string `xml:"path,attr,omitempty"`

	Keyword string `xml:"keyword,attr,omitempty"`
	Removal bool `xml:"removal,attr,omitempty"`
	CaseSensitive bool `xml:"caseSensitive,attr,omitempty"`
	Owner string `xml:"owner,attr,omitempty"`
	Degree string `xml:"degree,attr,omitempty"`
	Uni string `xml:"uni,attr,omitempty"`
	Gpa string `xml:"gpa,attr,omitempty"`

	MailServer string `xml:"mailServer,attr,omitempty"`
	Recipient string `xml:"recipient,attr,omitempty"`
	Subject string `xml:"subject,attr,omitempty"`
}

type MissionActions struct {
	Function string `xml:",chardata"`
	Value int `xml:"val,attr"`
	Suppress bool `xml:"suppress,attr,omitempty"`
}

type MissionLink struct {
	MissionPath string `xml:",chardata"`
	IsSilent bool `xml:"IsSilent,attr,omitempty"`
}

type BranchCollection struct {
	Branch []MissionLink `xml:"branch"`
}

type BoardPost struct {
	Title string `xml:"title,attr"`
	Requirements string `xml:"reqs,attr"`
	RequiredRank int `xml:"requiredRank,attr"`
	Body string `xml:",chardata"`
}

type MissionEmail struct {
	Sender string `xml:"sender"`
	Subject string `xml:"subject"`
	Body string `xml:"body"`

	Attachments EmailAttachmentCollection `xml:"attachments"`
}

type EmailAttachmentCollection struct {
	Notes []EmailNoteAttach `xml:"note"`
	Links []EmailLinkAttach `xml:"link"`
	Accounts []EmailAccountAttach `xml:"account"`
}

type EmailNoteAttach struct {
	Title string `xml:"title,attr"`
	Content string `xml:",chardata"`
}

type EmailLinkAttach struct {
	Comp string `xml:"comp,attr"`
}

type EmailAccountAttach struct {
	*EmailLinkAttach
	User string `xml:"user,attr"`
	Pass string `xml:"pass,attr"`
}