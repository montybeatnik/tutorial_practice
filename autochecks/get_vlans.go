package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/web_server_for_network_devices/devcon"
)

const (
	getVlansCMD = "show vlans | display xml"
)

// Vlans represents the XML formated "show vlans" output
type Vlans struct {
	L2ngL2aldVlanInstanceInformation struct {
		Text                       string `xml:",chardata"`
		Xmlns                      string `xml:"xmlns,attr"`
		Style                      string `xml:"style,attr"`
		L2ngL2aldVlanInstanceGroup []struct {
			Text                  string `xml:",chardata"`
			L2ngL2rtbBriefSummary struct {
				Text string `xml:",chardata"`
			} `xml:"l2ng-l2rtb-brief-summary"`
			L2ngL2rtbName struct {
				Text string `xml:",chardata"`
			} `xml:"l2ng-l2rtb-name"`
			L2ngL2rtbVlanName struct {
				Text string `xml:",chardata"`
			} `xml:"l2ng-l2rtb-vlan-name"`
			L2ngL2rtbVlanTag struct {
				Text string `xml:",chardata"`
			} `xml:"l2ng-l2rtb-vlan-tag"`
			L2ngL2rtbVlanMember []struct {
				Text                         string `xml:",chardata"`
				L2ngL2rtbVlanMemberInterface struct {
					Text string `xml:",chardata"`
				} `xml:"l2ng-l2rtb-vlan-member-interface"`
			} `xml:"l2ng-l2rtb-vlan-member"`
		} `xml:"l2ng-l2ald-vlan-instance-group"`
	} `xml:"l2ng-l2ald-vlan-instance-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner struct {
			Text string `xml:",chardata"`
		} `xml:"banner"`
	} `xml:"cli"`
}

// Mapper takes in a bytes.Buffer and returns a pointer to an
// MacTable object (struct)
func (v *Vlans) mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &v)
}

// Run Issues the command against the device and returns the object
func (v *Vlans) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: getVlansCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return v, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	v.mapper(buf)
	return v, nil
}
