package main

import (
	"HNExtensionConverter/RestClient"
	rModels "HNExtensionConverter/RestClient/models"
	"HNExtensionConverter/converters"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	dModels "HNExtensionConverter/ExtensionData"
)

func parseFile(path string, v interface{}) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	return xml.Unmarshal(data, v)
}

func isXML(info os.FileInfo) bool {
	return strings.HasSuffix(info.Name(), ".xml")
}

func main() {
	// SET-UP BaseURL
	//baseUrl, err := url.Parse("https://hn.lunasphere.co.uk/api")
	baseUrl, err := url.Parse("http://dev.lunasphere.co.uk/api")

	if err != nil {
		panic(err)
	}

	// SET-UP API CLIENT.
	client := RestClient.APIClient{
		BaseURL: baseUrl,
		//AuthToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEsImlhdCI6MTYwODU5NTEyOSwiZXhwIjoxNjA4NjgxNTI5fQ.6cX0A5H_J4K8WK1ynNjETRrm6q3ojQ00bORIDd3Zu4Q",
		AuthToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEsImlhdCI6MTYwODU5NzAyMCwiZXhwIjoxNjA4NjgzNDIwfQ.hPsfZsWxLKo9A4qm7tqu1FkwPtPOI5eH_rTbYQ-gnQQ",
		Client: &http.Client{}}

	// Set-up BASE Extension Path
	//extensionPath := "E:\\SteamLibrary\\steamapps\\common\\Hacknet\\Extensions\\IntroExtension"
	extensionPath := "E:\\SteamLibrary\\steamapps\\workshop\\content\\365450\\937939681"

	// Start on ExtensionInfo.xml
	xmlLoc := path.Join(extensionPath, "ExtensionInfo.xml")

	// Unmarshal XML file into Data Model.
	extInfo := new(dModels.ExtensionInfo)
	err = parseFile(xmlLoc, extInfo)

	if err != nil {
		panic(err)
	}

	// Run through Extension Converter to convert to API element.
	extConverter := converters.NewExtensionsConverter(client)
	extConverter.FetchPrelims()

	rExtension := extConverter.ConvertExtension(extInfo)
	// Create the extension to receive an extension ID for further work...
	newExtension, err := client.CreateExtension(rExtension)

	//newExtension.ExtensionID =


	// Update the client to use this extension for further requests.
	client.ExtensionID = newExtension.ExtensionID

	nodePath := path.Join(extensionPath, "Nodes")

	// Prepare the converter for Node processing
	nodeConverter := converters.NewNodeConverter(client)
	nodeConverter.FetchPrelims()

	// Create a map of Hacknet ID to DB ID
	nmap := make(map[string]int)

	// LOAD NODE DATA
	err = filepath.Walk(nodePath, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			// Double-check it's XML
			if strings.HasSuffix(info.Name(), ".xml") {
				nInfo := new(dModels.Node)

				fmt.Println("Loading Node Info from: " + path)
				err := parseFile(path, nInfo)

				if err != nil {
					panic(err)
				}

				// Convert and create the Node.
				rNode := nodeConverter.ConvertNode(nInfo)

				nmap[rNode.ID] = rNode.NodeID
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}


	// Load Missions
	missionPath := path.Join(extensionPath, "Missions")

	// Set-up converter
	mConverter := converters.NewMissionConverter(client)

	// Map for references later...
	// Map ID to Mission
	missionMap := make(map[int]rModels.Mission)
	// Map PATH to ID
	missionPathMap := make(map[string]int)

	err = filepath.Walk(missionPath, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			if isXML(info) {
				mInfo := new(dModels.Mission)

				fmt.Println("Loading Mission Info from: " + path)
				err := parseFile(path, mInfo)

				if err != nil {
					panic(err)
				}

				partialPath := strings.Replace(path, extensionPath + string(os.PathSeparator), "", 1)
				newMission := mConverter.ConvertMission(partialPath, mInfo)
				missionMap[newMission.MissionID] = newMission
				missionPathMap[strings.Replace(path, extensionPath, "", 1)] = newMission.MissionID
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	mConverter.CreateMissionLinks()
}