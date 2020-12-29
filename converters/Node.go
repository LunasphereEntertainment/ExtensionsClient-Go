package converters

import (
	dModels "HNExtensionConverter/ExtensionData"
	"HNExtensionConverter/RestClient"
	rModels "HNExtensionConverter/RestClient/models"
	"fmt"
	"strconv"
	"strings"
)

var (
	portMap map[int]rModels.Port
	/*sysTypes map[string]int{
		"Corporate": 1,
		"Home": 2,
		"Server": 3,
		"Empty": 4}*/

	nodeMap map[string]int
	pendingLinks map[int]string
)

type NodeConverter struct {
	client RestClient.APIClient
}

func NewNodeConverter(client RestClient.APIClient) NodeConverter {
	nodeMap = make(map[string]int)
	pendingLinks = make(map[int]string)

	return NodeConverter{client}
}

func (c *NodeConverter) FetchPrelims() {
	portMap = make(map[int]rModels.Port)

	ports, err := c.client.GetAllPorts()
	if err != nil {
		panic(err)
	}

	for _, port := range ports {
		portMap[port.Port] = port
	}
}

// TODO: Apply Proxy and Firewall Settings (if any)

// ConvertNode - Converts XML node data into a Data model that can be submitted to the RESTful API
/* Arguments:
*	client - A RESTClient object
*	nInfo - The parsed XML data object from Hacknet

*   Returns: Node - A converted Node object for the API.
 */
func (c *NodeConverter) ConvertNode(nInfo *dModels.Node) rModels.Node {
	// API Node Translation
	ports := make([]rModels.Port, 0)
	files := make([]rModels.File, 0)

	rNode := rModels.Node{
		ID:                      nInfo.ID,
		Name:                    nInfo.Name,
		IP:                      nInfo.IP,
		SecurityLevel:           nInfo.SecurityLevel,
		Icon:                    nInfo.Icon,
		AllowsDefaultBootModule: nInfo.AllowsDefaultBootModule,
		PortsForCrack:           nInfo.PortsForCrack,
		Ports:                   ports,
		Files:                   files,
		//TraceTime:               nInfo.traceTime,
	}

	// Create the computer
	newNode, err := c.client.CreateNode(rNode)
	if err != nil {
		panic(err)
	}

	fmt.Println("Parsing Node: " + nInfo.Name)

	// Resolve ports from the XML data into database entities.
	portNames := strings.Split(nInfo.Ports, ",")

	for _, prt := range portNames {
		prt = strings.TrimSpace(prt)

		pNum, err := strconv.Atoi(prt)

		if err != nil {
			fmt.Println("Invalid Port Number Received: " + prt)
		}

		port, ok := portMap[pNum]
		if ok {
			c.client.AddPort(port, newNode.NodeID)

			fmt.Printf("Resolved Port %d  ==>  %+v\n", pNum, port)
		} else {
			fmt.Printf("No port in Map for %d \n", pNum)
		}
	}

	// Resolve and Create files for this node.
	for _, file := range nInfo.Files {
		rFile := &rModels.File{
			//FileID:  0,
			Path:    file.Path,
			Name:    file.Name,
			Content: file.Content,
		}

		newFile, err := c.client.CreateFile(rFile)
		if err != nil {
			panic(err)
		}

		// Map link between this file and this node.
		c.client.AddFile(newFile.FileID, newNode.NodeID)
	}

	nodeMap[rNode.ID] = rNode.NodeID

	for _, link := range nInfo.Links {
		pendingLinks[newNode.NodeID] = link.Target
		//c.client.CreateLink()
	}

	return rNode
}

func (c *NodeConverter) CreateNodeLinks() {
	for parentID, child := range pendingLinks {
		childNodeId, ok := nodeMap[child]

		if ok {
			c.client.CreateLink(parentID, childNodeId)
		}
	}
}
