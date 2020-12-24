package models

type ExtensionInfo struct {
	ExtensionID int `json:"extensionId,omitempty"`
	Name string `json:"extensionName"`
	Description string `json:"description"`

	LanguageID int `json:"languageId"`

	AllowSaves bool `json:"allowSaves"`

	StartingThemeID int `json:"startingThemeId,omitempty"`
	StartingMusic int `json:"startingMusic,omitempty"`
	StartingMissionID int `json:"startingMissionId,omitempty"`

	WorkshopLanguage string `json:"workshop_language"`
	WorkshopDescription string `json:"workshop_description"`
	WorkshopTags []string `json:"workshop_tags"`
	WorkshopImg string `json:"workshop_img"`
	WorkshopID int `json:"workshop_id,omitempty"`

	StartingNodes []int `json:"startingNodes,omitempty"`
}

type ExtensionLanguage struct {
	LanguageID int `json:"langId"`
	ShortCode string `json:"lang"`
	Language string `json:"Language"`
}

type MusicInfo struct {
	MusicID int `json:"musicId"`
	OwnerID int `json:"ownerId"`
	Title string `json:"title"`
}