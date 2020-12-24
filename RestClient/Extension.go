package RestClient

import (
	"HNExtensionConverter/RestClient/models"
	"fmt"
)

func (c *APIClient) GetExtension(extensionId int) (*models.ExtensionInfo, error) {

	path := fmt.Sprintf("extensions/%d", extensionId)

	req, err := c.Build("GET", path, nil)

	if err != nil {
		return nil, err
	}

	extension := new(models.ExtensionInfo)
	_, err = c.Do(req, extension)

	return extension, err
}

func (c *APIClient) CreateExtension(extension models.ExtensionInfo) (*models.ExtensionInfo, error) {
	req, err := c.Build("POST", "extensions/new", extension)

	if err != nil {
		return nil, err
	}

	ext := new(models.ExtensionInfo)
	_, err = c.Do(req, ext)

	return ext, err
}

func (c *APIClient) GetLanguages() ([]models.ExtensionLanguage, error) {
	req, err := c.Build("GET", "extensions/languages", nil)

	if err != nil {
		return nil, err
	}

	var langs []models.ExtensionLanguage
	_, err = c.Do(req, &langs)

	return langs, err
}