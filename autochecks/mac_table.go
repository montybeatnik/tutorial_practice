package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/tutorial_practice/devcon"
)

// MacTable represents the XML format of the 'show ethernet-switching table {VLAN}' output
type MacTable struct {
	L2ngL2aldRtbMacdb struct {
		Text                  string `xml:",chardata"`
		L2ngL2aldMacEntryVlan struct {
			Text           string `xml:",chardata"`
			Style          string `xml:"style,attr"`
			MacCountGlobal struct {
				Text string `xml:",chardata"`
			} `xml:"mac-count-global"`
			LearntMacCount struct {
				Text string `xml:",chardata"`
			} `xml:"learnt-mac-count"`
			L2ngL2MacRoutingInstance struct {
				Text string `xml:",chardata"`
			} `xml:"l2ng-l2-mac-routing-instance"`
			L2ngL2VlanID struct {
				Text string `xml:",chardata"`
			} `xml:"l2ng-l2-vlan-id"`
			L2ngMacEntry []struct {
				Text              string `xml:",chardata"`
				L2ngL2MacVlanName struct {
					Text string `xml:",chardata"`
				} `xml:"l2ng-l2-mac-vlan-name"`
				L2ngL2MacAddress struct {
					Text string `xml:",chardata"`
				} `xml:"l2ng-l2-mac-address"`
				L2ngL2MacFlags struct {
					Text string `xml:",chardata"`
				} `xml:"l2ng-l2-mac-flags"`
				L2ngL2MacAge struct {
					Text string `xml:",chardata"`
				} `xml:"l2ng-l2-mac-age"`
				L2ngL2MacLogicalInterface struct {
					Text string `xml:",chardata"`
				} `xml:"l2ng-l2-mac-logical-interface"`
				L2ngL2MacFwdNextHop struct {
					Text string `xml:",chardata"`
				} `xml:"l2ng-l2-mac-fwd-next-hop"`
				L2ngL2MacRtrID struct {
					Text string `xml:",chardata"`
				} `xml:"l2ng-l2-mac-rtr-id"`
			} `xml:"l2ng-mac-entry"`
		} `xml:"l2ng-l2ald-mac-entry-vlan"`
	} `xml:"l2ng-l2ald-rtb-macdb"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner struct {
			Text string `xml:",chardata"`
		} `xml:"banner"`
	} `xml:"cli"`
}

// Mapper takes in a bytes.Buffer and returns a pointer to an
// MacTable object (struct)
func (mt *MacTable) mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &mt)
}

// Run Issues the command against the device and returns the object
func (mt *MacTable) Run(p Params) (interface{}, error) {
	getVLANCMD := fmt.Sprintf("show ethernet-switching table vlan-name %v | display xml", p.Args[0])
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: getVLANCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return mt, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	mt.mapper(buf)
	return mt, nil
}
