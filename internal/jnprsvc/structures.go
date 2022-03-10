package jnprsvc

import (
	"encoding/xml"
)

// Nh Next-hop is part of the RouteTables struct
//
type Nh struct {
	Text             string `xml:",chardata"`
	SelectedNextHop  string `xml:"selected-next-hop"`
	Via              string `xml:"via"`
	NhLocalInterface string `xml:"nh-local-interface"`
	To               string `xml:"to"`
	MplsLabel        string `xml:"mpls-label"`
	NhTable          string `xml:"nh-table"`
}

// Age is part of the RouteTables struct
//
type Age struct {
	Text    string `xml:",chardata"`
	Seconds string `xml:"seconds,attr"`
}

// RtEntry is part of the RouteTables struct
//
type RtEntry struct {
	Text            string `xml:",chardata"`
	ActiveTag       string `xml:"active-tag"`
	CurrentActive   string `xml:"current-active"`
	LastActive      string `xml:"last-active"`
	ProtocolName    string `xml:"protocol-name"`
	Preference      string `xml:"preference"`
	Age             Age    `xml:"age"`
	Nh              []Nh   `xml:"nh"`
	NhType          string `xml:"nh-type"`
	Metric          string `xml:"metric"`
	RtTag           string `xml:"rt-tag"`
	LocalPreference string `xml:"local-preference"`
	AsPath          string `xml:"as-path"`
	ValidationState string `xml:"validation-state"`
	LearnedFrom     string `xml:"learned-from"`
}

// Rt is part of the RouteTables struct
//
type Rt struct {
	Text           string         `xml:",chardata"`
	Style          string         `xml:"style,attr"`
	RtDestination  string         `xml:"rt-destination"`
	RtEntry        []RtEntry      `xml:"rt-entry"`
	RtPrefixLength RtPrefixLength `xml:"rt-prefix-length"`
}

// RtPrefixLength is part of the RouteTables struct
//
type RtPrefixLength struct {
	Text string `xml:",chardata"`
	Emit string `xml:"emit,attr"`
}

// RouteTable is part of the RouteTables struct
//
type RouteTable struct {
	Text               string `xml:",chardata"`
	TableName          string `xml:"table-name"`
	DestinationCount   string `xml:"destination-count"`
	TotalRouteCount    string `xml:"total-route-count"`
	ActiveRouteCount   string `xml:"active-route-count"`
	HolddownRouteCount string `xml:"holddown-route-count"`
	HiddenRouteCount   string `xml:"hidden-route-count"`
	Rt                 []Rt   `xml:"rt"`
}

// RouteInformation is part of the RouteTables struct
//
type RouteInformation struct {
	Text       string       `xml:",chardata"`
	Xmlns      string       `xml:"xmlns,attr"`
	RouteTable []RouteTable `xml:"route-table"`
}

// RouteTables contains all the routing tables in a junos device
//
type RouteTables struct {
	XMLName          xml.Name         `xml:"rpc-reply"`
	Text             string           `xml:",chardata"`
	Junos            string           `xml:"junos,attr"`
	Output           []string         `xml:"output"`
	RouteInformation RouteInformation `xml:"route-information"`
}

func NewRTs() *RouteTables {
	return &RouteTables{}
}

// InstanceRib is part of the InstanceInformation struct
//
type InstanceRib struct {
	Text              string `xml:",chardata"`
	IribName          string `xml:"irib-name"`
	IribActiveCount   string `xml:"irib-active-count"`
	IribHolddownCount string `xml:"irib-holddown-count"`
	IribHiddenCount   string `xml:"irib-hidden-count"`
}

// InstanceCore is part of the InstanceInformation struct
//
type InstanceCore struct {
	Text         string        `xml:",chardata"`
	InstanceName string        `xml:"instance-name"`
	InstanceType string        `xml:"instance-type"`
	InstanceRib  []InstanceRib `xml:"instance-rib"`
}

// InstanceInfo is part of the InstanceInformation struct
//
type InstanceInfo struct {
	Text         string         `xml:",chardata"`
	Xmlns        string         `xml:"xmlns,attr"`
	Style        string         `xml:"style,attr"`
	InstanceCore []InstanceCore `xml:"instance-core"`
}

// InstanceInformation contains the complete structure that holds the routing instances of a
// junos device
//
type InstanceInformation struct {
	XMLName             xml.Name     `xml:"rpc-reply"`
	Text                string       `xml:",chardata"`
	Junos               string       `xml:"junos,attr"`
	InstanceInformation InstanceInfo `xml:"instance-information"`
}

func NewInstanceInformation() *InstanceInformation {
	return &InstanceInformation{}
}
