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
}