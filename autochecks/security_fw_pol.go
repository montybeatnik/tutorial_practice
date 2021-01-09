package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/web_server_for_network_devices/devcon"
)

const (
	fwPolConfig = "show configuration security policies | display xml"
)

// SecFWPol represents the xml formatted 'show configuration security policies'
type SecFWPol struct {
	Configuration struct {
		Text            string `xml:",chardata"`
		CommitSeconds   string `xml:"commit-seconds,attr"`
		CommitLocaltime string `xml:"commit-localtime,attr"`
		CommitUser      string `xml:"commit-user,attr"`
		Security        struct {
			Text     string `xml:",chardata"`
			Policies struct {
				Text   string `xml:",chardata"`
				Policy []struct {
					Text         string `xml:",chardata"`
					FromZoneName struct {
						Text string `xml:",chardata"`
					} `xml:"from-zone-name"`
					ToZoneName struct {
						Text string `xml:",chardata"`
					} `xml:"to-zone-name"`
					Policy []struct {
						Text     string `xml:",chardata"`
						Inactive string `xml:"inactive,attr"`
						Name     struct {
							Text string `xml:",chardata"`
						} `xml:"name"`
						Match struct {
							Text          string `xml:",chardata"`
							SourceAddress []struct {
								Text string `xml:",chardata"`
							} `xml:"source-address"`
							DestinationAddress []struct {
								Text string `xml:",chardata"`
							} `xml:"destination-address"`
							Application []struct {
								Text string `xml:",chardata"`
							} `xml:"application"`
						} `xml:"match"`
						Then struct {
							Text   string `xml:",chardata"`
							Permit struct {
								Text   string `xml:",chardata"`
								Tunnel struct {
									Text     string `xml:",chardata"`
									IpsecVpn struct {
										Text string `xml:",chardata"`
									} `xml:"ipsec-vpn"`
									PairPolicy struct {
										Text string `xml:",chardata"`
									} `xml:"pair-policy"`
								} `xml:"tunnel"`
							} `xml:"permit"`
							Deny struct {
								Text string `xml:",chardata"`
							} `xml:"deny"`
							Log struct {
								Text        string `xml:",chardata"`
								SessionInit struct {
									Text string `xml:",chardata"`
								} `xml:"session-init"`
							} `xml:"log"`
							Count struct {
								Text string `xml:",chardata"`
							} `xml:"count"`
							Reject struct {
								Text string `xml:",chardata"`
							} `xml:"reject"`
						} `xml:"then"`
					} `xml:"policy"`
				} `xml:"policy"`
				PolicyRematch struct {
					Text string `xml:",chardata"`
				} `xml:"policy-rematch"`
			} `xml:"policies"`
		} `xml:"security"`
	} `xml:"configuration"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner struct {
			Text string `xml:",chardata"`
		} `xml:"banner"`
	} `xml:"cli"`
}

// Mapper takes in a bytes.Buffer and returns a pointer to an
// SecFWPol object (struct)
func (cfg *SecFWPol) mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &cfg)
}

// Run Issues the command against the device and returns the object
func (cfg *SecFWPol) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: fwPolConfig,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return cfg, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	cfg.mapper(buf)
	return cfg, nil
}
