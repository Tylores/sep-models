// Sep is a libary of IEEE 2030.5-2018 modes for clients and servers
package sep

import (
	"encoding/xml"
)

const (
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

// A EndDeviceListLink is a Link to a List of EndDevice instances
type EndDeviceListLink struct {
	XMLName xml.Name `xml:"EndDeviceListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

// A MirrorUsagePointListLink is a Link to a List of MirrorUsagePoint instances
type MirrorUsagePointListLink struct {
	XMLName xml.Name `xml:"MirrorUsagePointListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

// A SelfDeviceLink is a Link to a SelfDevice instance
type SelfDeviceLink struct {
	XMLName xml.Name `xml:"SelfDeviceLink"`
	Href    string   `xml:"href,attr"`
}

// A DeviceCapability is returned by the URI provided by DNS-SD, to allow clients to find the URIs
// to the resources in which they are interested.
type DeviceCapability struct {
	XMLName  xml.Name `xml:"DeviceCapability"`
	Href     string   `xml:"href,attr"`
	PollRate int      `xml:"pollRate,attr"`
	FunctionSetAssignmentsBase
	EndDevices        *EndDeviceListLink
	MirrorUsagePoints *MirrorUsagePointListLink
	SelfDevice        *SelfDeviceLink
}

type CustomerAccountListLink struct {
	XMLName xml.Name `xml:"CustomerAccountListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

type DemandResponseProgramListLink struct {
	XMLName xml.Name `xml:"DemandResponseProgramListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

type DERProgramListLink struct {
	XMLName xml.Name `xml:"DERProgramListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

type FileListLink struct {
	XMLName xml.Name `xml:"FileListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

type MessagingProgramListLink struct {
	XMLName xml.Name `xml:"MessagingProgramListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

type PrepaymentListLink struct {
	XMLName xml.Name `xml:"PrepaymentListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

type ResponseSetListLink struct {
	XMLName xml.Name `xml:"ResponseSetListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

type TariffProfileListLink struct {
	XMLName xml.Name `xml:"TariffProfileListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

type TimeLink struct {
	XMLName xml.Name `xml:"TimeLink"`
	Href    string   `xml:"href,attr"`
}

type UsagePointListLink struct {
	XMLName xml.Name `xml:"UsagePointListLink"`
	All     uint     `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

type FunctionSetAssignmentsBase struct {
	CustomerAccounts       *CustomerAccountListLink
	DemandResponsePrograms *DemandResponseProgramListLink
	DERPrograms            *DERProgramListLink
	Files                  *FileListLink
	MessagingPrograms      *MessagingProgramListLink
	Prepayments            *PrepaymentListLink
	ResponseSets           *ResponseSetListLink
	TariffProfiles         *TariffProfileListLink
	Time                   *TimeLink
	UsagePoints            *UsagePointListLink
}
