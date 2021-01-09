package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/tutorial_practice/devcon"
)

var (
	softwareVerCMD = "show version | display xml"
)

type SoftwareVersion struct {
	SoftwareInformation struct {
		Text               string `xml:",chardata"`
		HostName           string `xml:"host-name"`
		ProductModel       string `xml:"product-model"`
		ProductName        string `xml:"product-name"`
		JunosVersion       string `xml:"junos-version"`
		PackageInformation []struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name"`
			Comment string `xml:"comment"`
		} `xml:"package-information"`
	} `xml:"software-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

// Mapper takes in a bytes buffer and unmarshals the buffer into the sv object
func (sv *SoftwareVersion) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &sv)
}

// Run Issues the command against the device and returns the object
func (sv *SoftwareVersion) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: softwareVerCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return sv, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	sv.Mapper(buf)
	return sv, nil
}
