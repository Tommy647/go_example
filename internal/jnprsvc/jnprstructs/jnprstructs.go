package jnprstructs

import "encoding/xml"

// RouteTable corresponds to the juniper structure of a routing table, using composition to
// express the different components of this object
type RouteTable struct {
	XMLName            xml.Name `xml:"route-table"`
	Text               string   `xml:",chardata"`
	TableName          string   `xml:"table-name"`
	DestinationCount   string   `xml:"destination-count"`
	TotalRouteCount    string   `xml:"total-route-count"`
	ActiveRouteCount   string   `xml:"active-route-count"`
	HolddownRouteCount string   `xml:"holddown-route-count"`
	HiddenRouteCount   string   `xml:"hidden-route-count"`
	Rt
}

type NextHop struct {
	Text             string `xml:",chardata"`
	SelectedNextHop  string `xml:"selected-next-hop"`
	To               string `xml:"to"`
	Via              string `xml:"via"`
	NhLocalInterface string `xml:"nh-local-interface"`
	MplsLabel        string `xml:"mpls-label"`
}

type Age struct {
	Text    string `xml:",chardata"`
	Seconds string `xml:"seconds,attr"`
}

type RtEntry struct {
	Text          string `xml:",chardata"`
	ActiveTag     string `xml:"active-tag"`
	CurrentActive string `xml:"current-active"`
	LastActive    string `xml:"last-active"`
	ProtocolName  string `xml:"protocol-name"`
	Preference    string `xml:"preference"`
	Age
	Metric string `xml:"metric"`
	NextHop
	NhType          string `xml:"nh-type"`
	LocalPreference string `xml:"local-preference"`
	LearnedFrom     string `xml:"learned-from"`
	AsPath          string `xml:"as-path"`
	ValidationState string `xml:"validation-state"`
	RtTag           string `xml:"rt-tag"`
}

type Rt struct {
	Text          string `xml:",chardata"`
	Style         string `xml:"style,attr"`
	RtDestination string `xml:"rt-destination"`
	RtEntry
}

func NewRT() *RouteTable {
	return &RouteTable{}
}
