package RestClient

import (
	"HNExtensionConverter/RestClient/models"
	"fmt"
	"strconv"
)

// METHOD OPS for accessing computer nodes in the extension.
func (c *APIClient) GetNode(nodeId int) (*models.Node, error) {
	path := fmt.Sprintf("nodes/%d", nodeId)

	req, err := c.Build("GET", path, nil)

	if err != nil {
		return nil, err
	}

	node := new(models.Node)
	_, err = c.Do(req, node)

	return node, err
}

func (c *APIClient) UpdateNode(nodeId string, info models.Node) (*models.Node, error) {
	method := "PUT"
	if nodeId == "new" {
		method = "POST"
	}

	req, err := c.Build(method, fmt.Sprintf("nodes/%s", nodeId), info)

	if err != nil {
		return nil, err
	}

	ret := new(models.Node)
	_, err = c.Do(req, ret)

	return ret, err
}

func (c *APIClient) CreateNode(node models.Node) (*models.Node, error) {
	return c.UpdateNode("new", node)
}


// PORT OPS for accessing and manipulating ports within a defined computer.
func (c *APIClient) GetAllPorts() ([]models.Port, error) {
	req, err := c.Build("GET", "nodes/ports/all", nil)

	if err != nil {
		return nil, err
	}

	var ports []models.Port
	_, err = c.Do(req, &ports)

	return ports, err
}

func (c *APIClient) GetPorts(nodeId int) ([]models.Port, error) {
	req, err := c.Build("GET", fmt.Sprintf("nodes/ports/list/%d", nodeId), nil)

	if err != nil {
		return nil, err
	}

	var ports []models.Port
	_, err = c.Do(req, ports)

	return ports, err
}

func (c *APIClient) AddPort(port models.Port, nodeId int) error {
	path := fmt.Sprintf("nodes/ports/map?node=%d&port=%d", nodeId, port.PortID)

	req, err := c.Build("GET", path, nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

func (c *APIClient) RemovePort(port models.Port, nodeId int) error {
	path := fmt.Sprintf("nodes/ports/unmap?node=%d&port=%d", nodeId, port.PortID)

	req, err := c.Build("GET", path, nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

// FILE OPS for accessing and manipulating files within a defined computer
func (c *APIClient) listFiles(nodeId string) ([]models.File, error) {
	req, err := c.Build("GET", fmt.Sprintf("nodes/files/list/%s", nodeId), nil)

	if err != nil {
		return nil, err
	}

	var files []models.File
	_, err = c.Do(req, files)

	return files, err
}

func (c *APIClient) GetAllFiles() ([]models.File, error) {
	return c.listFiles("all")
}

func (c *APIClient) GetFiles(nodeId int) ([]models.File, error) {
	return c.listFiles(strconv.Itoa(nodeId))
}

func (c *APIClient) CreateFile(file *models.File) (*models.File, error) {
	req, err := c.Build("POST", "nodes/files/new", file)

	if err != nil {
		return nil, err
	}

	newFile := new(models.File)
	_, err = c.Do(req, newFile)

	return newFile, err
}

func (c *APIClient) UpdateFile(fileId int, info *models.File) error {
	req, err := c.Build("PUT", fmt.Sprintf("nodes/files/%d", fileId), nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)

	return err
}

func (c *APIClient) DeleteFile(fileId int) error {
	req, err := c.Build("DELETE", fmt.Sprintf("nodes/files/%d", fileId), nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)

	return err
}

func (c *APIClient) AddFile(fileID int, nodeId int) error {
	path := fmt.Sprintf("nodes/files/map?file=%d&node=%d", fileID, nodeId)

	req, err := c.Build("GET", path, nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

func (c *APIClient) RemoveFile(fileID int, nodeId int) error {
	path := fmt.Sprintf("nodes/files/unmap?file=%d&node=%d", fileID, nodeId)

	req, err := c.Build("GET", path, nil)

	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}