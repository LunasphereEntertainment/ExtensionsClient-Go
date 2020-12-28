package ExtensionData

type Node struct {
	ID string `xml:"id,attr"`
	Name string `xml:"name,attr"`
	IP string `xml:"ip,attr"`
	Icon string `xml:"icon,attr"`
	Ports string `xml:"ports"`
	PortRemap string `xml:"portRemap"`


	SecurityLevel int `xml:"security,attr"`
	SysType string `xml:"type,attr"`
	PortsForCrack int `xml:"portsForCrack"`

	AllowsDefaultBootModule bool `xml:"allowsDefaultBootModule,attr"`

	Files []File `xml:"file"`
	Links []DLink `xml:"link"`
}

type File struct {
	Path string `xml:"path,attr"`
	Name string `xml:"name,attr"`
	Content string `xml:",chardata"`
}

type DLink struct {
	Target string `xml:"target"`
}