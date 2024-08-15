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

func TestDeviceCapabilityXmlOpt(t *testing.T) {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("sep.xsd", xsdvalidate.ParsErrDefault)
	check("newValidator", err)
	defer xsdhandler.Free()

	dcap := DeviceCapability{}
	//dcap.Href = "/dcap"
	//dcap.PollRate = 900
	dcap.SelfDeviceLink = &SelfDeviceLink{}

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
func TestDeviceCapabilityXml(t *testing.T) {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("sep.xsd", xsdvalidate.ParsErrDefault)
	check("newValidator", err)
	defer xsdhandler.Free()
	dcap := DeviceCapability{
		FunctionSetAssignmentsBase: FunctionSetAssignmentsBase{
			Resource: Resource{
				Href: "/dcap",
			},
			CustomerAccountListLink: &CustomerAccountListLink{
				ListLink: ListLink{
					Link: Link{
						Href: "/ca",
					},
					All: 1,
				},
			},
			DemandResponseProgramListLink: &DemandResponseProgramListLink{
				ListLink: ListLink{
					Link: Link{
						Href: "/dr",
					},
					All: 1,
				},
			},
			DERProgramListLink: &DERProgramListLink{
				ListLink: ListLink{
					Link: Link{
						Href: "/derp",
					},
					All: 1,
				},
			},
			FileListLink: &FileListLink{
				ListLink: ListLink{
					Link: Link{
						Href: "/fs",
					},
					All: 1,
				},
			},
			MessagingProgramListLink: &MessagingProgramListLink{
				ListLink: ListLink{
					Link: Link{
						Href: "/msg",
					},
					All: 1,
				},
			},
			PrepaymentListLink: &PrepaymentListLink{
				ListLink: ListLink{
					Link: Link{
						Href: "/ppy",
					},
					All: 1,
				},
			},
			ResponseSetListLink: &ResponseSetListLink{
				ListLink: ListLink{
					Link: Link{
						Href: "/rsps",
					},
					All: 1,
				},
			},
			TariffProfileListLink: &TariffProfileListLink{
				ListLink: ListLink{
					Link: Link{
						Href: "/tp",
					},
					All: 1,
				},
			},
			TimeLink: &TimeLink{
				Link: Link{
					Href: "/tm",
				},
			},
			UsagePointListLink: &UsagePointListLink{
				ListLink: ListLink{
					Link: Link{
						Href: "/up",
					},
					All: 1,
				},
			},
		},
		PollRate: 900,
		EndDeviceListLink: &EndDeviceListLink{
			ListLink: ListLink{
				Link: Link{
					Href: "/edev",
				},
				All: 1,
			},
		},
		MirrorUsagePointListLink: &MirrorUsagePointListLink{
			ListLink: ListLink{
				Link: Link{
					Href: "/mup",
				},
				All: 1,
			},
		},
		SelfDeviceLink: &SelfDeviceLink{
			Link: Link{
				Href: "/sdev",
			},
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

func TestTimeXML(t *testing.T) {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("sep.xsd", xsdvalidate.ParsErrDefault)
	check("newValidator", err)
	defer xsdhandler.Free()

	model := NewTime("/tm")

	xml_str := stringify(model)
	xmlhandler, err := xsdvalidate.NewXmlHandlerMem(xml_str, xsdvalidate.ParsErrDefault)
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
