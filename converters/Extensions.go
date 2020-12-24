package converters

import (
	dModels "HNExtensionConverter/ExtensionData"
	"HNExtensionConverter/RestClient"
	"HNExtensionConverter/RestClient/models"
	rModels "HNExtensionConverter/RestClient/models"
)

var (
	languages []models.ExtensionLanguage
)

type ExtensionConverter struct{
	client RestClient.APIClient
}

func NewExtensionsConverter(client RestClient.APIClient) ExtensionConverter {
	return ExtensionConverter{client}
}

func (c *ExtensionConverter) FetchPrelims() {
	var err error
	languages, err = c.client.GetLanguages()

	if err != nil {
		panic(err)
	}
}

func findLanguage(langCode string) int {
	for _, lang := range languages {
		if lang.ShortCode == langCode {
			return lang.LanguageID
		}
	}

	return 0
}

func (c *ExtensionConverter) ConvertExtension(extInfo *dModels.ExtensionInfo) rModels.ExtensionInfo {
	/* MISSING:
	StartingTheme

	StartingMissionID
	StartingMusicID
	StartingNodeID's
	*/

	langId := findLanguage(extInfo.Language)
	if langId == 0 {
		langId = 1
	}

	rExtension := rModels.ExtensionInfo{
		Name:                extInfo.Name,
		Description:         extInfo.Description,
		LanguageID:          langId,
		AllowSaves:          extInfo.AllowSaves,
		StartingThemeID:     0,
		StartingMusic:       0,
		StartingMissionID:   0,
		WorkshopLanguage:    extInfo.WorkshopLanguage,
		WorkshopDescription: extInfo.WorkshopDescription,
		WorkshopTags:        nil,
		WorkshopImg:         extInfo.WorkshopPreviewImagePath,
		WorkshopID:          0,
		StartingNodes:       nil,
	}

	return rExtension
}