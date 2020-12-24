package models

// Node - a Hacknet Computer Node in an extension.
type Node struct {
	NodeID int `json:"nodeId"`
	ID string `json:"id"`
	Name string `json:"name"`
	IP string `json:"ip"`

	SecurityLevel int            `json:"securityLevel"`
	Icon string                  `json:"icon"`
	AllowsDefaultBootModule bool `json:"allowsDefaultBootModule"`
	Ports []Port                 `json:"ports"`

	PortsForCrack int `json:"portsForCrack"`
	TraceTime float32 `json:"traceTime"`

	Files []File `json:"files"`

	HasTracker bool `json:"tracker"`
	ProxyTime float32 `json:"proxyTime"`
	FirewallLevel int `json:"fwall_Level"`
	FirewallSolution string `json:"fwall_solution"`
	FirewallAdditionalTime float32 `json:"fwall_additional"`
}

// Port - Definition for a security port accessible on a particular Computer Node.
type Port struct {
	PortID int `json:"portId"`
	Type string `json:"portType"`
	Port int `json:"port"`
}

// File - Description of a file, it's path and contents existing on one or more nodes throughout an extension.
type File struct {
	FileID int `json:"fileId"`
	Path string `json:"path"`
	Name string `json:"name"`
	Content string `json:"contents"`
}