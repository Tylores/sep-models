package sep

import (
	"encoding/xml"
	"github.com/terminalstatic/go-xsd-validate"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stringify(A any) []byte {
	out, err := xml.MarshalIndent(A, " ", "  ")
	check(err)
	return out
}

func newValidator() *xsdvalidate.XsdHandler {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("sep.xsd", xsdvalidate.ParsErrDefault)
	check(err)
	defer xsdhandler.Free()
	return xsdhandler
}

func TestDeviceCapability(t *testing.T) {
	xsdhandler := newValidator()
	dcap := &DeviceCapability{
		Href:      "/dcap",
		Poll_rate: 900,
		Time: TimeLink{
			Href: "/tm",
		},
		SelfDevice: SelfDeviceLink{
			Href: "/sdev",
		},
		EndDevices: EndDeviceListLink{
			Href: "/edev",
			All:  1,
		},
	}
	dcap_xml := stringify(dcap)
	xmlhandler, err := xsdvalidate.NewXmlHandlerMem(dcap_xml, xsdvalidate.ParsErrDefault)
	check(err)

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
