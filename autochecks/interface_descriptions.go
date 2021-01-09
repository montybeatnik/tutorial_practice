package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/web_server_for_network_devices/devcon"
)

const (
	interfaceDescriptionCMD = "show interface descriptions | display xml"
)

type InterfaceDescription struct {
	InterfaceInformation struct {
		Text             string `xml:",chardata"`
		Xmlns            string `xml:"xmlns,attr"`
		Style            string `xml:"style,attr"`
		LogicalInterface []struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name"`
			AdminStatus string `xml:"admin-status"`
			OperStatus  string `xml:"oper-status"`
			Description string `xml:"description"`
		} `xml:"logical-interface"`
		PhysicalInterface []struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name"`
			AdminStatus string `xml:"admin-status"`
			OperStatus  string `xml:"oper-status"`
			Description string `xml:"description"`
		} `xml:"physical-interface"`
	} `xml:"interface-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

func (it *InterfaceDescription) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &it)
}

func (it *InterfaceDescription) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: interfaceDescriptionCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return it, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	it.Mapper(buf)
	return it, nil
}
