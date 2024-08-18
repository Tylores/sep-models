// Sep is a libary of IEEE 2030.5-2018 modes for clients and servers
package sep

import (
	"fmt"
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
type HexBinary160 [40]byte
type HexBinary128 [32]byte
type HexBinary32 [8]byte
type HexBinary8 [2]byte
type SFDIType uint64
type SubscribableType uint
type DeviceCategoryType HexBinary32
type PINType uint32
type MRIDType HexBinary128
type VersionType uint16
type PowerOfTenMultiplierType int8

// Links provide a reference, via URI, to another resource.
type Link struct {
	Href string `xml:"href,attr"`
}
type SelfDeviceLink struct{ Link }
type FileListLink struct{ ListLink }
type TimeLink struct{ Link }
type ConfigurationLink struct{ Link }
type DeviceInformationLink struct{ Link }
type DeviceStatusLink struct{ Link }
type FileStatusLink struct{ Link }
type PowerStatusLink struct{ Link }
type RegistrationLink struct{ Link }

// Container to hold a collection of object instances or references.
// See Design Pattern section for additional details.
type List struct {
	Resource
	All     uint32 `xml:"all,attr"`
	Results uint32 `xml:"results,attr"`
}

// ListLinks provide a reference, via URI, to a List.
type ListLink struct {
	Link
	All uint32 `xml:"all,attr"`
}

func NewListLink(href string, all uint32) *ListLink {
	ll := ListLink{}
	ll.Href = href
	ll.All = all
	return &ll
}

type EndDeviceListLink struct{ ListLink }
type MirrorUsagePointListLink struct{ ListLink }
type CustomerAccountListLink struct{ ListLink }
type DemandResponseProgramListLink struct{ ListLink }
type DERListLink struct{ ListLink }
type DERProgramListLink struct{ ListLink }
type MessagingProgramListLink struct{ ListLink }
type PrepaymentListLink struct{ ListLink }
type ResponseSetListLink struct{ ListLink }
type TariffProfileListLink struct{ ListLink }
type UsagePointListLink struct{ ListLink }
type IPInterfaceListLink struct{ ListLink }
type LoadShedAvailabilityListLink struct{ ListLink }
type LogEventListLink struct{ ListLink }
type FlowReservationRequestListLink struct{ ListLink }
type FlowReservationResponseListLink struct{ ListLink }
type FunctionSetAssignmentsListLink struct{ ListLink }
type SubscriptionListLink struct{ ListLink }

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

// A Resource to which a Response can be requested.
type RespondableResource struct {
	Resource
	ReplyTo          *string    `xml:"replyTo,attr,omitempty"`
	ResponseRequired HexBinary8 `xml:"responseRequired,attr"`
}

func NewRespondableResource(href string) *RespondableResource {
	rr := RespondableResource{}
	rr.Href = href
	rr.ResponseRequired = HexBinary8([]byte("00"))
	fmt.Println(rr.ResponseRequired)
	return &rr
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
	ConfigurationLink            *ConfigurationLink            `xml:"ConfigurationLink,omitempty"`
	DERListLink                  *DERListLink                  `xml:"DERListLink,omitempty"`
	DeviceCategory               *DeviceCategoryType           `xml:"deviceCategory,omitempty"`
	DeviceInformationLink        *DeviceInformationLink        `xml:"DeviceInformationLink,omitempty"`
	DeviceStatusLink             *DeviceStatusLink             `xml:"DeviceStatusLink,omitempty"`
	FileStatusLink               *FileStatusLink               `xml:"FileStatusLink,omitempty"`
	IPInterfaceListLink          *IPInterfaceListLink          `xml:"IPInterfaceListLink,omitempty"`
	LFDI                         *HexBinary160                 `xml:"lFDI,omitempty"`
	LoadShedAvailabilityListLink *LoadShedAvailabilityListLink `xml:"LoadShedAvailabilityListLink,omitempty"`
	LogEventListLink             *LogEventListLink             `xml:"LogEventListLink,omitempty"`
	PowerStatusLink              *PowerStatusLink              `xml:"PowerStatusLink,omitempty"`
	SFDI                         SFDIType                      `xml:"sFDI"`
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
	ChangedTime                     TimeType                         `xml:"changedTime"`
	Enabled                         *bool                            `xml:"enabled,omitempty"`
	FlowReservationRequestListLink  *FlowReservationRequestListLink  `xml:"FlowReservationRequestListLink,omitempty"`
	FlowReservationResponseListLink *FlowReservationResponseListLink `xml:"FlowReservationResponseListLink,omitempty"`
	FunctionSetAssignmentsListLink  *FunctionSetAssignmentsListLink  `xml:"FunctionSetAssignmentsListLink,omitempty"`
	PostRate                        *uint32                          `xml:"postRate,omitempty"`
	RegistrationLink                *RegistrationLink                `xml:"RegistrationLink,omitempty"`
	SubscriptionListLink            *SubscriptionListLink            `xml:"SubscriptionListLink,omitempty"`
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

// Registration represents an authorization to access the resources on a host.
type Registration struct {
	Resource
	DateTimeRegistered TimeType `xml:"dateTimeRegistered"`
	PIN                PINType  `xml:"pIN"`
	PollRate           uint32   `xml:"pollRate,attr"`
}

func NewRegistration(href string, dt TimeType, pin PINType) *Registration {
	rg := Registration{
		Resource:           Resource{Href: href},
		DateTimeRegistered: dt,
		PIN:                pin,
		PollRate:           900,
	}
	return &rg
}

// This is a root class to provide common naming attributes for all classes needing naming attributes
type IdentifiedObject struct {
	Resource
	MRID        MRIDType     `xml:"mRID"`
	Description *string      `xml:"description,omitempty"`
	Version     *VersionType `xml:"version,omitempty"`
}

func NewIdentifiedObject(href string, mrid MRIDType) *IdentifiedObject {
	obj := IdentifiedObject{}
	obj.Href = href
	obj.MRID = mrid
	return &obj
}

// Real electrical energy, signed.
type SignedRealEnergy struct {
	Multiplier PowerOfTenMultiplierType `xml:"multiplier"`
	Value      int64                    `xml:"value"`
}

func NewSignedRealEnergy(scale PowerOfTenMultiplierType, value int64) *SignedRealEnergy {
	energy := SignedRealEnergy{
		Multiplier: scale,
		Value:      value,
	}
	return &energy
}

// Interval of date and time.
type DateTimeInterval struct {
	Duration uint32   `xml:"duration"`
	Start    TimeType `xml:"start"`
}

func NewDateTimeInterval(duration uint32, start TimeType) *DateTimeInterval {
	interval := DateTimeInterval{
		Duration: duration,
		Start:    start,
	}
	return &interval
}

// The active (real) power P (in W) is the product of root-mean-square (RMS) voltage,
// RMS current, and cos(theta) where theta is the phase angle of current relative to voltage.
// It is the primary measure of the rate of flow of energy.
type ActivePower struct {
	Multiplier PowerOfTenMultiplierType `xml:"multiplier"`
	Value      int16                    `xml:"value"`
}

func NewActivePower(scale PowerOfTenMultiplierType, value int16) *ActivePower {
	power := ActivePower{
		Multiplier: scale,
		Value:      value,
	}
	return &power
}

// The RequestStatus object is used to indicate the current status of a Flow Reservation Request.
type RequestStatus struct {
	DateTime TimeType `xml:"dateTime"`
	Status   uint8    `xml:"requestStatus"`
}

func NewRequestStatus(dt TimeType, status uint8) *RequestStatus {
	rs := RequestStatus{
		DateTime: dt,
		Status:   status,
	}
	return &rs
}

// Used to request flow transactions. Client EndDevices submit a request for charging or discharging
// from the server. The server creates an associated FlowReservationResponse containing the charging
// parameters and interval to provide a lower aggregated demand at the premises, or within a larger
// part of the distribution system.
type FlowReservationRequest struct {
	IdentifiedObject
	CreationTime      TimeType         `xml:"creationTime"`
	DurationRequested *uint16          `xml:"durationRequested,omitempty"`
	EnergyRequested   SignedRealEnergy `xml:"energyRequested"`
	IntervalRequested DateTimeInterval `xml:"intervalRequested"`
	PowerRequested    ActivePower      `xml:"powerRequested"`
	RequestStatus     RequestStatus    `xml:"RequestStatus"`
}

func NewFlowReservationRequest(
	obj IdentifiedObject,
	creation TimeType,
	energy SignedRealEnergy,
	interval DateTimeInterval,
	power ActivePower,
	status RequestStatus) *FlowReservationRequest {
	frq := FlowReservationRequest{IdentifiedObject: obj}
	frq.CreationTime = creation
	frq.EnergyRequested = energy
	frq.IntervalRequested = interval
	frq.PowerRequested = power
	frq.RequestStatus = status
	return &frq
}

// A List element to hold FlowReservationRequest objects.
type FlowReservationRequestList struct {
	List
	FlowReservationRequest []*FlowReservationRequest `xml:"FlowReservationRequest,omitempty"`
	PollRate               uint32                    `xml:"pollRate,attr"`
}

// Current status information relevant to a specific object. The Status object is used
// to indicate the current status of an Event. Devices can read the containing resource
// (e.g. TextMessage) to get the most up to date status of the event.  Devices can also
// subscribe to a specific resource instance to get updates when any of its attributes change,
// including the Status object.
type EventStatus struct {
	CurrentStatus             uint8    `xml:"currentStatus"`
	DateTime                  TimeType `xml:"dateTime"`
	PotentiallySuperseded     bool     `xml:"potentiallySuperseded"`
	PotentiallySupersededTime TimeType `xml:"potentiallySupersededTime"`
	Reason                    string   `xml:"reason,omitempty"`
}

// An IdentifiedObject to which a Response can be requested.
type RespondableSubscribableIdentifiedObject struct {
	RespondableResource
	MRID         MRIDType         `xml:"mRID"`
	Description  *string          `xml:"description,omitempty"`
	Version      *VersionType     `xml:"version,omitempty"`
	Subscribable SubscribableType `xml:"subscribable,attr"`
}

func NewRespondableSubscribableIdentifiedObject(
	href string,
	mrid MRIDType) *RespondableSubscribableIdentifiedObject {
	rsio := RespondableSubscribableIdentifiedObject{
		RespondableResource: *NewRespondableResource(href),
	}
	rsio.MRID = mrid
	rsio.Subscribable = SubscribableType(0)
	return &rsio
}

// An Event indicates information that applies to a particular period of time.
// Events SHALL be executed relative to the time of the server, as described
// in the Time function set section 11.1.
type Event struct {
	RespondableSubscribableIdentifiedObject
	CreationTime TimeType         `xml:"creationTime"`
	EventStatus  EventStatus      `xml:"EventStatus"`
	Interval     DateTimeInterval `xml:"interval"`
}

func NewEvent(
	rsio RespondableSubscribableIdentifiedObject,
	creation TimeType,
	status EventStatus,
	interval DateTimeInterval) *Event {
	event := Event{RespondableSubscribableIdentifiedObject: rsio}
	event.CreationTime = creation
	event.EventStatus = status
	event.Interval = interval
	return &event
}

// The server may modify the charging or discharging parameters and interval to provide
// a lower aggregated demand at the premises, or within a larger part of the distribution system.
type FlowReservationResponse struct {
	Event
	EnergyAvailable SignedRealEnergy `xml:"energyAvailable"`
	PowerAvailable  ActivePower      `xml:"powerAvailable"`
	Subject         MRIDType         `xml:"subject"`
}

func NewFlowReservationResponse(
	event Event,
	energy SignedRealEnergy,
	power ActivePower,
	subject MRIDType) *FlowReservationResponse {
	frp := FlowReservationResponse{Event: event}
	frp.EnergyAvailable = energy
	frp.PowerAvailable = power
	frp.Subject = subject
	return &frp
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

func NewSelfDevice(adev AbstractDevice) *SelfDevice {
	sdev := SelfDevice{AbstractDevice: adev}
	return &sdev
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
