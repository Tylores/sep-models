package sep

import (
	"encoding/xml"
	"fmt"
	"github.com/terminalstatic/go-xsd-validate"
	"math"
	"testing"
)

func check(ctx string, e error) {
	if e != nil {
		fmt.Printf("Error thrown during: %s\n", ctx)
		panic(e)
	}
}

func stringify(A any) []byte {
	out, err := xml.MarshalIndent(A, " ", "  ")
	check("stringify", err)
	fmt.Println(string(out))
	return out
}

func NewHandler(xsd string) *xsdvalidate.XsdHandler {
	handler, err := xsdvalidate.NewXsdHandlerUrl(xsd, xsdvalidate.ParsErrDefault)
	check("newValidator", err)
	return handler
}

func Validate(xml []byte, t *testing.T) {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()

	handler := NewHandler("sep.xsd")
	defer handler.Free()

	xmlhandler, err := xsdvalidate.NewXmlHandlerMem(xml, xsdvalidate.ParsErrDefault)
	check("xmlhandler", err)

	err = handler.Validate(xmlhandler, xsdvalidate.ValidErrDefault)
	if err != nil {
		switch err.(type) {
		case xsdvalidate.ValidationError:
			t.Errorf("Validation Error: %v\n", err)
			t.Errorf("Error in line: %d\n", err.(xsdvalidate.ValidationError).Errors[0].Line)
			t.Errorf("Message %s\n", err.(xsdvalidate.ValidationError).Errors[0].Message)
		default:
			t.Errorf("Error: %v\n", err)
		}
	}
}

func TestDeviceCapabilityXmlOpt(t *testing.T) {
	obj := NewDeviceCapability("/dcap")
	xml := stringify(&obj)
	Validate(xml, t)
}

func TestDeviceCapabilityXml(t *testing.T) {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("sep.xsd", xsdvalidate.ParsErrDefault)
	check("newValidator", err)
	defer xsdhandler.Free()
	obj := DeviceCapability{
		FunctionSetAssignmentsBase: FunctionSetAssignmentsBase{
			Resource: Resource{Href: "/dcap"},
			CustomerAccountListLink: &CustomerAccountListLink{
				ListLink: *NewListLink("/ca", 1),
			},
			DemandResponseProgramListLink: &DemandResponseProgramListLink{
				ListLink: *NewListLink("/dr", 1),
			},
			DERProgramListLink: &DERProgramListLink{
				ListLink: *NewListLink("/derp", 1),
			},
			FileListLink: &FileListLink{
				ListLink: *NewListLink("/fs", 1),
			},
			MessagingProgramListLink: &MessagingProgramListLink{
				ListLink: *NewListLink("/msg", 1),
			},
			PrepaymentListLink: &PrepaymentListLink{
				ListLink: *NewListLink("/ppy", 1),
			},
			ResponseSetListLink: &ResponseSetListLink{
				ListLink: *NewListLink("/rsps", 1),
			},
			TariffProfileListLink: &TariffProfileListLink{
				ListLink: *NewListLink("/tp", 1),
			},
			TimeLink: &TimeLink{Link: Link{Href: "/tm"}},
			UsagePointListLink: &UsagePointListLink{
				ListLink: *NewListLink("/up", 1),
			},
		},
		PollRate: 900,
		EndDeviceListLink: &EndDeviceListLink{
			ListLink: *NewListLink("/edev", 1),
		},
		MirrorUsagePointListLink: &MirrorUsagePointListLink{
			ListLink: *NewListLink("/mup", 1),
		},
		SelfDeviceLink: &SelfDeviceLink{Link: Link{Href: "/sdev"}},
	}

	xml := stringify(obj)
	Validate(xml, t)
}

func TestTimeXML(t *testing.T) {
	obj := NewTime("/tm")
	xml := stringify(&obj)
	Validate(xml, t)
}

func TestEndDeviceXMLOpt(t *testing.T) {
	adev := NewAbstractDevice("/edev", 1, 1234)
	obj := NewEndDevice(*adev, 1111)
	xml := stringify(&obj)
	Validate(xml, t)
}

func TestEndDeviceXML(t *testing.T) {
	adev := NewAbstractDevice("/edev", 1, 1234)
	adev.ConfigurationLink = &ConfigurationLink{Link: Link{Href: "/cfg"}}
	adev.DERListLink = &DERListLink{ListLink: *NewListLink("/der", 1)}
	category := []byte("12345678")
	dcat := DeviceCategoryType(category)
	adev.DeviceCategory = &dcat
	adev.DeviceInformationLink = &DeviceInformationLink{Link: Link{Href: "/dinf"}}
	adev.DeviceStatusLink = &DeviceStatusLink{Link: Link{Href: "/dst"}}
	adev.FileStatusLink = &FileStatusLink{Link: Link{Href: "/fst"}}
	adev.IPInterfaceListLink = &IPInterfaceListLink{ListLink: *NewListLink("/ipi", 1)}
	lfdi_bs := []byte("aaaaabbbbbcccccdddddeeeeefffff1111122222")
	hb := HexBinary160(lfdi_bs)
	adev.LFDI = &hb
	adev.LoadShedAvailabilityListLink = &LoadShedAvailabilityListLink{ListLink: *NewListLink("/ls", 1)}
	adev.LogEventListLink = &LogEventListLink{ListLink: *NewListLink("/le", 1)}
	adev.PowerStatusLink = &PowerStatusLink{Link: Link{Href: "/ps"}}
	obj := NewEndDevice(*adev, 1111)
	enabled := true
	obj.Enabled = &enabled
	obj.FlowReservationRequestListLink = &FlowReservationRequestListLink{ListLink: *NewListLink("/frq", 1)}
	obj.FlowReservationResponseListLink = &FlowReservationResponseListLink{ListLink: *NewListLink("/frp", 1)}
	obj.FunctionSetAssignmentsListLink = &FunctionSetAssignmentsListLink{ListLink: *NewListLink("/fsa", 1)}
	xml := stringify(&obj)
	Validate(xml, t)
}

func TestEndDeviceListXML(t *testing.T) {
	all := uint32(4)
	sub_res := NewSubscribableResource("/edev", SubscribableType(1))
	sub_list := NewSubscribableList(*sub_res, all, all)
	list := EndDeviceList{}
	list.SubscribableList = *sub_list
	for i := uint32(0); i < all; i++ {
		path := fmt.Sprintf("%s/%d", sub_res.Href, i)
		adev := NewAbstractDevice(path, 1, SFDIType(1000+i))
		obj := NewEndDevice(*adev, 1111)
		list.EndDevice = append(list.EndDevice, obj)
	}
	xml := stringify(&list)
	Validate(xml, t)
}

func TestSelfDeviceXMLOpt(t *testing.T) {
	adev := NewAbstractDevice("/sdev", 1, 1234)
	obj := NewSelfDevice(*adev)
	xml := stringify(&obj)
	Validate(xml, t)
}

func TestRegistrationXML(t *testing.T) {
	obj := NewRegistration("/rg", 1234, 777)
	xml := stringify(&obj)
	Validate(xml, t)
}

func TestFlowReservationRequestXMLOpt(t *testing.T) {
	mrid_bs := []byte("aaaaabbbbbcccccdddddeeeeefffff11")
	mrid := MRIDType(mrid_bs)
	id_obj := NewIdentifiedObject("/rg", mrid)
	energy := NewSignedRealEnergy(0, 100)
	power := NewActivePower(0, 10)
	dt := NewDateTimeInterval(1000, 12345)
	status := NewRequestStatus(12345, 1)
	obj := NewFlowReservationRequest(*id_obj, 12345, *energy, *dt, *power, *status)
	xml := stringify(&obj)
	Validate(xml, t)
}
func TestFlowReservationRequestXML(t *testing.T) {
	mrid_bs := []byte("aaaaabbbbbcccccdddddeeeeefffff11")
	mrid := MRIDType(mrid_bs)
	id_obj := NewIdentifiedObject("/rg", mrid)
	energy := NewSignedRealEnergy(0, 100)
	power := NewActivePower(0, 10)
	dt := NewDateTimeInterval(1000, 12345)
	status := NewRequestStatus(12345, 1)
	obj := NewFlowReservationRequest(*id_obj, 12345, *energy, *dt, *power, *status)
	e := float64(energy.Value) * math.Pow(10.0, float64(energy.Multiplier))
	p := float64(power.Value) * math.Pow(10.0, float64(power.Multiplier))
	d := uint16(e / p)
	obj.DurationRequested = &d
	xml := stringify(&obj)
	Validate(xml, t)
}

func TestFlowReservationResponseXMLOpt(t *testing.T) {
	mrid_bs := []byte("aaaaabbbbbcccccdddddeeeeefffff11")
	mrid := MRIDType(mrid_bs)

	dt := NewDateTimeInterval(1000, 12345)

	rsio := NewRespondableSubscribableIdentifiedObject("/rsps", mrid)
	xml := stringify(&rsio)
	Validate(xml, t)

	status := EventStatus{}
	status.CurrentStatus = 1
	status.DateTime = 12345
	status.Reason = "because I said so."
	status.PotentiallySuperseded = false
	status.PotentiallySupersededTime = 0

	event := NewEvent(*rsio, 12345, status, *dt)

	energy := NewSignedRealEnergy(0, 100)
	power := NewActivePower(0, 10)
	obj := NewFlowReservationResponse(*event, *energy, *power, mrid)
	xml = stringify(&obj)
	Validate(xml, t)
}

func TestFlowReservationResponseXML(t *testing.T) {
	mrid_bs := []byte("aaaaabbbbbcccccdddddeeeeefffff11")
	mrid := MRIDType(mrid_bs)

	dt := NewDateTimeInterval(1000, 12345)

	reply := "/me"
	description := "Testing xml conversion"
	version := VersionType(111)
	rsio := NewRespondableSubscribableIdentifiedObject("/rsps", mrid)
	rsio.ReplyTo = &reply
	rsio.Description = &description
	rsio.Version = &version
	xml := stringify(&rsio)
	Validate(xml, t)

	status := EventStatus{}
	status.CurrentStatus = 1
	status.DateTime = 12345
	status.Reason = "because I said so."
	status.PotentiallySuperseded = false
	status.PotentiallySupersededTime = 0

	event := NewEvent(*rsio, 12345, status, *dt)

	energy := NewSignedRealEnergy(0, 100)
	power := NewActivePower(0, 10)
	obj := NewFlowReservationResponse(*event, *energy, *power, mrid)
	xml = stringify(&obj)
	Validate(xml, t)
}

func TestResponseXMLOpt(t *testing.T) {
	mrid_bs := []byte("aaaaabbbbbcccccdddddeeeeefffff11")
	mrid := MRIDType(mrid_bs)

	lfdi_bs := []byte("aaaaabbbbbcccccdddddeeeeefffff1111122222")
	lfdi := HexBinary160(lfdi_bs)

	obj := NewResponse("/rsps", lfdi, mrid)
	xml := stringify(&obj)
	Validate(xml, t)
}

func TestResponseXML(t *testing.T) {
	mrid_bs := []byte("aaaaabbbbbcccccdddddeeeeefffff11")
	mrid := MRIDType(mrid_bs)

	lfdi_bs := []byte("aaaaabbbbbcccccdddddeeeeefffff1111122222")
	lfdi := HexBinary160(lfdi_bs)

	obj := NewResponse("/rsps", lfdi, mrid)
	var created TimeType = 12345
	obj.CreatedDateTime = &created
	var status uint8 = 1
	obj.Status = &status
	xml := stringify(&obj)
	Validate(xml, t)
}

func TestPowerStatusXMLOpt(t *testing.T) {
	obj := NewPowerStatus("/rsps", 1, 12345, 1)
	xml := stringify(&obj)
	Validate(xml, t)
}

func TestPowerStatusXML(t *testing.T) {
	power := NewActivePower(1, 100)
	max_power := NewActivePower(1, 150)
	energy := NewRealEnergy(1, 1000)
	info := NewPEVInfo(*power, *energy, *max_power, 50, 9000, 23456, 12345)
	obj := NewPowerStatus("/rsps", 1, 12345, 1)
	var charge_remaining PerCent = 10
	obj.EstimatedChargeRemaining = &charge_remaining
	var time_remaining uint32 = 123
	obj.EstimatedTimeRemaining = &time_remaining
	obj.PEVInfo = info
	var session_time uint32 = 12
	obj.SessionTimeOnBattery = &session_time
	var total_time uint32 = 1
	obj.TotalTimeOnBattery = &total_time
	xml := stringify(&obj)
	Validate(xml, t)
}
