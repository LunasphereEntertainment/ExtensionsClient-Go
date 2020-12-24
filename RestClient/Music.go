package RestClient

import (
	"HNExtensionConverter/RestClient/models"
	"fmt"
)

func (c *APIClient) listMusic(filter string) ([]models.MusicInfo, error) {
	req, err := c.Build("GET", fmt.Sprintf("music/list/%s", filter), nil)

	if err != nil {
		return nil, err
	}

	var music []models.MusicInfo
	_, err = c.Do(req, music)

	return music, err
}

func (c *APIClient) GetAllMusic() ([]models.MusicInfo, error) {
	return c.listMusic("all")
}

func (c *APIClient) GetUserMusic() ([]models.MusicInfo, error) {
	return c.listMusic("custom")
}

func (c *APIClient) GetExtensionMusic() ([]models.MusicInfo, error) {
	return c.listMusic("extension")
}


/*func (c *APIClient) CreateTrack(tInfo models.MusicInfo, buffer bytes.Buffer) (*models.MusicInfo, error) {
	req, err := c.Build("POST", "music/new", tInfo)

	if err != nil {
		return nil, err
	}

	music := new(models.MusicInfo)
	_, err = c.Do(req, music)

	return music, err
}*/

func (c *APIClient) UpdateTrack(trackId int, tInfo models.MusicInfo) error {
	req, err := c.Build("PUT", fmt.Sprintf("music/%d", trackId), tInfo)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)

	return err
}

func (c *APIClient) DeleteTrack(trackId int) error {
	req, err := c.Build("DELETE", fmt.Sprintf("music/%d", trackId), nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)

	return err
}