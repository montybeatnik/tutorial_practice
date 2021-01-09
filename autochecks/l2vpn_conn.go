package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/tutorial_practice/devcon"
)

const (
	l2vpnCMD = "show l2vpn connections | display xml"
)

type L2VPNConn struct {
	L2vpnConnectionInformation struct {
		Text     string `xml:",chardata"`
		Xmlns    string `xml:"xmlns,attr"`
		Instance []struct {
			Text           string `xml:",chardata"`
			Xmlns          string `xml:"xmlns,attr"`
			Style          string `xml:"style,attr"`
			InstanceName   string `xml:"instance-name"`
			EdgeProtection string `xml:"edge-protection"`
			ReferenceSite  struct {
				Text        string `xml:",chardata"`
				LocalSiteID string `xml:"local-site-id"`
				Connection  struct {
					Text             string `xml:",chardata"`
					Heading          string `xml:"heading,attr"`
					ConnectionID     string `xml:"connection-id"`
					ConnectionType   string `xml:"connection-type"`
					ConnectionStatus string `xml:"connection-status"`
					LastChange       string `xml:"last-change"`
					UpTransitions    string `xml:"up-transitions"`
					RemotePe         string `xml:"remote-pe"`
					ControlWord      string `xml:"control-word"`
					ControlWordType  string `xml:"control-word-type"`
					InboundLabel     string `xml:"inbound-label"`
					OutboundLabel    string `xml:"outbound-label"`
					LocalInterface   struct {
						Text                   string `xml:",chardata"`
						InterfaceName          string `xml:"interface-name"`
						InterfaceStatus        string `xml:"interface-status"`
						InterfaceEncapsulation string `xml:"interface-encapsulation"`
					} `xml:"local-interface"`
					VcFlowLabelTransmit string `xml:"vc-flow-label-transmit"`
					VcFlowLabelReceive  string `xml:"vc-flow-label-receive"`
				} `xml:"connection"`
			} `xml:"reference-site"`
		} `xml:"instance"`
	} `xml:"l2vpn-connection-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

func (l2vpn *L2VPNConn) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &l2vpn)
}

// Run Issues the command against the device and returns the object
func (l2vpn *L2VPNConn) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: l2vpnCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return l2vpn, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	l2vpn.Mapper(buf)
	return l2vpn, nil
}
