package models

type ActionSetDefinition struct {
	ActionSetID int `json:"actionSetId"`
	ExtensionID int `json:"extensionId"`
	Name string `json:"name"`
}

type ActionSet struct {
	*ActionSetDefinition

	Conditions []Condition `json:"conditions"`
}

type Condition struct {
	ConditionID int `json:"conditionId"`
	ActionSetID int `json:"actionSetId"`
	TypeID int `json:"typeId"`

	NeedsMissionComplete bool `json:"needsMissionComplete"`
	RequiredFlags string `json:"requiredFlags"`

	TargetNodeID int `json:"targetNodeId"`
	TypeText string `json:"typeText"`

	Actions []Action `json:"actions"`
}

type Action struct {
	ActionID int `json:"actionId"`
	TypeID int `json:"typeId"`

	LoadActionSetId int `json:"loadActionSetId"`
	LoadMissionID int `json:"loadMissionId"`

	SwitchThemeID int `json:"switchThemeId"`

	FileID int `json:"fileId"`

	IRCMessageID int `json:"ircMessageId"`

	DelayCompID int `json:"delayCompId"`
	Delay float32 `json:"delay"`

	TargetCompID int `json:"targetCompId"`

	FunctionID int `json:"functionId"`
	FunctionValue int `json:"functionValue"`

	ConditionID int `json:"conditionId"`
	TypeText string `json:"typeText"`
}