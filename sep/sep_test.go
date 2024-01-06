package sep

import (
	"encoding/xml"
	"fmt"
	"github.com/terminalstatic/go-xsd-validate"
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

func TestDeviceCapabilityXml(t *testing.T) {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("sep.xsd", xsdvalidate.ParsErrDefault)
	check("newValidator", err)
	defer xsdhandler.Free()

	dcap := DeviceCapability{
		Href:     "/dcap",
		PollRate: 900,
		FunctionSetAssignmentsBase: FunctionSetAssignmentsBase{
			CustomerAccounts: &CustomerAccountListLink{
				Href: "/ca",
				All:  1,
			},
			DemandResponsePrograms: &DemandResponseProgramListLink{
				Href: "/dr",
				All:  1,
			},
			DERPrograms: &DERProgramListLink{
				Href: "/derp",
				All:  1,
			},
			Files: &FileListLink{
				Href: "/fs",
				All:  1,
			},
			MessagingPrograms: &MessagingProgramListLink{
				Href: "/msg",
				All:  1,
			},
			Prepayments: &PrepaymentListLink{
				Href: "/ppy",
				All:  1,
			},
			ResponseSets: &ResponseSetListLink{
				Href: "/rsps",
				All:  1,
			},
			TariffProfiles: &TariffProfileListLink{
				Href: "/tp",
				All:  1,
			},
			Time: &TimeLink{
				Href: "/tm",
			},
			UsagePoints: &UsagePointListLink{
				Href: "/up",
				All:  1,
			},
		},
		EndDevices: &EndDeviceListLink{
			Href: "/edev",
			All:  1,
		},
		MirrorUsagePoints: &MirrorUsagePointListLink{
			Href: "/mup",
			All:  1,
		},
		SelfDevice: &SelfDeviceLink{
			Href: "/sdev",
		},
	}
	dcap_xml := stringify(dcap)
	xmlhandler, err := xsdvalidate.NewXmlHandlerMem(dcap_xml, xsdvalidate.ParsErrDefault)
	check("xmlhandler", err)

	err = xsdhandler.Validate(xmlhandler, xsdvalidate.ValidErrDefault)
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
