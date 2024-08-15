// Sep is a libary of IEEE 2030.5-2018 modes for clients and servers
package sep

import (
	"time"
	_ "time/tzdata"
)

const (
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

type TimeType int64
type TimeOffsetType int32
type HexBinary160 [20]byte
type HexBinary32 [4]byte
type SFDIType uint64
type SubscribableType uint
type DeviceCategoryType HexBinary32

// Links provide a reference, via URI, to another resource.
type Link struct {
	Href string `xml:"href,attr"`
}
type SelfDeviceLink struct{ Link }
type FileListLink struct{ ListLink }
type TimeLink struct{ Link }

// ListLinks provide a reference, via URI, to a List.
type ListLink struct {
	Link
	All uint32 `xml:"all,attr"`
}
type EndDeviceListLink struct{ ListLink }
type MirrorUsagePointListLink struct{ ListLink }
type CustomerAccountListLink struct{ ListLink }
type DemandResponseProgramListLink struct{ ListLink }
type DERProgramListLink struct{ ListLink }
type MessagingProgramListLink struct{ ListLink }
type PrepaymentListLink struct{ ListLink }
type ResponseSetListLink struct{ ListLink }
type TariffProfileListLink struct{ ListLink }
type UsagePointListLink struct{ ListLink }

// A resource is an addressable unit of information, either a collection (List)
// or instance of an object (identifiedObject, or simply, Resource)
type Resource struct {
	Href string `xml:"href,attr,omitempty"`
}

// A Resource to which a Subscription can be requested.
type SubscribableResource struct {
	Resource
	Subscribable SubscribableType `xml:"subscribable,attr"`
}

func NewSubscribableResource(href string, subscribable SubscribableType) *SubscribableResource {
	sub_res := SubscribableResource{}
	sub_res.Href = href
	sub_res.Subscribable = subscribable
	return &sub_res
}

// A List to which a Subscription can be requested.
type SubscribableList struct {
	SubscribableResource
	All     uint32 `xml:"all,attr"`
	Results uint32 `xml:"results,attr"`
}

func NewSubscribableList(resource SubscribableResource, all uint32, results uint32) *SubscribableList {
	sub_list := SubscribableList{SubscribableResource: resource}
	sub_list.All = all
	sub_list.Results = results
	return &sub_list
}

// The EndDevice providing the resources available within the DeviceCapabilities.
type AbstractDevice struct {
	SubscribableResource
	ConfigurationLink            *string             `xml:"ConfigurationLink,omitempty"`
	DERListLink                  *string             `xml:"DERListLink,omitempty"`
	DeviceCategory               *DeviceCategoryType `xml:"deviceCategory,omitempty"`
	DeviceInformationLink        *string             `xml:"DeviceInformationLink,omitempty"`
	DeviceStatusLink             *string             `xml:"DeviceStatusLink,omitempty"`
	FileStatusLink               *string             `xml:"FileStatusLink,omitempty"`
	IPInterfaceListLink          *string             `xml:"IPInterfaceListLink,omitempty"`
	LFDI                         *HexBinary160       `xml:"lFDI,omitempty"`
	LoadShedAvailabilityListLink *string             `xml:"LoadShedAvailabilityListLink,omitempty"`
	LogEventListLink             *string             `xml:"LogEventListLink,omitempty"`
	PowerStatusLink              *string             `xml:"PowerStatusLink,omitempty"`
	SFDI                         SFDIType            `xml:"sFDI"`
}

func NewAbstractDevice(href string, subscribable SubscribableType, sfdi SFDIType) *AbstractDevice {
	adev := AbstractDevice{}
	adev.Href = href
	adev.Subscribable = subscribable
	adev.SFDI = sfdi
	return &adev
}

// Asset container that performs one or more end device functions. Contains information
// about individual devices in the network.
type EndDevice struct {
	AbstractDevice
	ChangedTime                     TimeType `xml:"changedTime"`
	Enabled                         *bool    `xml:"enabled,omitempty"`
	FlowReservationRequestListLink  *string  `xml:"FlowReservationRequestListLink,omitempty"`
	FlowReservationResponseListLink *string  `xml:"FlowReservationResponseListLink,omitempty"`
	PostRate                        *uint32  `xml:"postRate,omitempty"`
	RegistrationLink                *string  `xml:"RegistrationLink,omitempty"`
	SubscriptionListLink            *string  `xml:"SubscriptionListLink,omitempty"`
}

func NewEndDevice(adev AbstractDevice, time TimeType) *EndDevice {
	edev := EndDevice{AbstractDevice: adev}
	edev.ChangedTime = time
	return &edev
}

// A List element to hold EndDevice objects.
type EndDeviceList struct {
	SubscribableList
	EndDevice []*EndDevice `xml:"EndDevice,omitempty"`
	PollRate  uint32       `xml:"pollRate,attr"`
}

// A DeviceCapability is returned by the URI provided by DNS-SD, to allow clients to find the URIs
// to the resources in which they are interested.
type DeviceCapability struct {
	FunctionSetAssignmentsBase
	PollRate                 uint32                    `xml:"pollRate,attr"`
	EndDeviceListLink        *EndDeviceListLink        `xml:"EndDeviceListLink,omitempty"`
	MirrorUsagePointListLink *MirrorUsagePointListLink `xml:"MirrorUsagePointListLink,omitempty"`
	SelfDeviceLink           *SelfDeviceLink           `xml:"SelfDeviceLink,omitempty"`
}

func NewDeviceCapability(href string) *DeviceCapability {
	dcap := DeviceCapability{}
	dcap.Href = href
	dcap.PollRate = 900
	return &dcap
}

// The EndDevice providing the resources available within the DeviceCapabilities.
type SelfDevice struct {
	AbstractDevice
	PollRate uint32 `xml:"pollRate,attr"`
}

// Defines a collection of function set instances that are to be used by one
// or more devices as indicated by the EndDevice object(s) of the server.
type FunctionSetAssignmentsBase struct {
	Resource
	CustomerAccountListLink       *CustomerAccountListLink       `xml:"CustomerAccountListLink,omitempty"`
	DemandResponseProgramListLink *DemandResponseProgramListLink `xml:"DemandResponseProgramListLink,omitempty"`
	DERProgramListLink            *DERProgramListLink            `xml:"DERProgramListLink,omitempty"`
	FileListLink                  *FileListLink                  `xml:"FileListLink,omitempty"`
	MessagingProgramListLink      *MessagingProgramListLink      `xml:"MessagingProgramListLink,omitempty"`
	PrepaymentListLink            *PrepaymentListLink            `xml:"PrepaymentListLink,omitempty"`
	ResponseSetListLink           *ResponseSetListLink           `xml:"ResponseSetListLink,omitempty"`
	TariffProfileListLink         *TariffProfileListLink         `xml:"TariffProfileListLink,omitempty"`
	TimeLink                      *TimeLink                      `xml:"TimeLink,omitempty"`
	UsagePointListLink            *UsagePointListLink            `xml:"UsagePointListLink,omitempty"`
}

// Contains the representation of time, constantly updated.
type Time struct {
	Resource
	CurrentTime  TimeType       `xml:"currentTime"`
	DstEndTime   TimeType       `xml:"dstEndTime"`
	DstOffset    TimeType       `xml:"dstOffset"`
	DstStartTime TimeType       `xml:"dstStartTime"`
	LocalTime    TimeType       `xml:"localTime,omitempty"`
	Quality      uint8          `xml:"quality"`
	TzOffset     TimeOffsetType `xml:"tzOffset"`
	PollRate     uint32         `xml:"pollRate,attr"`
}

func NewTime(href string) *Time {
	t := time.Now()
	_, offset := t.Local().Zone()
	start, end := t.Local().ZoneBounds()

	tm := Time{}
	tm.Href = href
	tm.CurrentTime = TimeType(t.Unix())
	tm.DstEndTime = TimeType(end.Unix())
	tm.DstOffset = 3600
	tm.DstStartTime = TimeType(start.Unix())
	tm.LocalTime = tm.CurrentTime + TimeType(offset)
	tm.Quality = 3
	tm.TzOffset = TimeOffsetType(offset)
	tm.PollRate = 900
	return &tm
}
