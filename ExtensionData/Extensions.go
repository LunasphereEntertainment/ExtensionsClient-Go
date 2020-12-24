package ExtensionData

type ExtensionInfo struct {
	Language string `xml:"Language"`
	Name string `xml:"Name"`
	Description string `xml:"Description"`

	StartingNodes string `xml:"StartingVisibleNodes"`
	StartingActions string `xml:"StartingActions"`
	StartingMission string `xml:"StartingMission"`

	AllowSaves bool `xml:"AllowSaves"`
	StartsWithTutorial bool `xml:"StartsWithTutorial"`
	HasIntro bool `xml:"HasIntroStartup"`

	StartingTheme string `xml:"StartingTheme"`
	IntroSong string `xml:"IntroStartupSong"`

	WorkshopDescription string `xml:"WorkshopDescription"`
	WorkshopLanguage string `xml:"WorkshopLanguage"`
	WorkshopVisibility int `xml:"WorkshopVisibility"`
	WorkshopTags string `xml:"WorkshopTags"`
	WorkshopPreviewImagePath string `xml:"WorkshopPreviewImagePath"`
	WorkshopPublishID string `xml:"WorkshopPublishID"`

	Factions []Faction `xml:"Faction"`
}

type Faction string