package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/tutorial_practice/devcon"
)

const (
	interfaceTerseCMD = "show interfaces terse | display xml"
)

type InterfaceTerse struct {
	InterfaceInformation struct {
		Text              string `xml:",chardata"`
		Xmlns             string `xml:"xmlns,attr"`
		Style             string `xml:"style,attr"`
		PhysicalInterface []struct {
			Text             string `xml:",chardata"`
			Name             string `xml:"name"`
			AdminStatus      string `xml:"admin-status"`
			OperStatus       string `xml:"oper-status"`
			LogicalInterface []struct {
				Text              string `xml:",chardata"`
				Name              string `xml:"name"`
				AdminStatus       string `xml:"admin-status"`
				OperStatus        string `xml:"oper-status"`
				FilterInformation string `xml:"filter-information"`
				AddressFamily     []struct {
					Text              string `xml:",chardata"`
					AddressFamilyName struct {
						Text string `xml:",chardata"`
						Emit string `xml:"emit,attr"`
					} `xml:"address-family-name"`
					InterfaceAddress []struct {
						Text     string `xml:",chardata"`
						IfaLocal struct {
							Text string `xml:",chardata"`
							Emit string `xml:"emit,attr"`
						} `xml:"ifa-local"`
						IfaDestination struct {
							Text string `xml:",chardata"`
							Emit string `xml:"emit,attr"`
						} `xml:"ifa-destination"`
					} `xml:"interface-address"`
				} `xml:"address-family"`
			} `xml:"logical-interface"`
		} `xml:"physical-interface"`
	} `xml:"interface-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

func (it *InterfaceTerse) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &it)
}

func (it *InterfaceTerse) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: interfaceTerseCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return it, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	it.Mapper(buf)
	return it, nil
}
