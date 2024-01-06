// Sep is a libary of IEEE 2030.5-2018 modes for clients and servers
package sep

import (
	"encoding/xml"
)

// A EndDeviceListLink is a Link to a List of EndDevice instances
type EndDeviceListLink struct {
	XMLName xml.Name `xml:"EndDeviceListLink"`
	All     int      `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

// A TimeLink is a Link to a Time instance
type TimeLink struct {
	XMLName xml.Name `xml:"TimeLink"`
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
	XMLName    xml.Name          `xml:"DeviceCapability"`
	Poll_rate  int               `xml:"pollRate,attr"`
	Href       string            `xml:"href,attr"`
	Time       TimeLink          `xml:"TimeLink"`
	SelfDevice SelfDeviceLink    `xml:"SelfDeviceLink"`
	EndDevices EndDeviceListLink `xml:"EndDeviceListLink"`
}
