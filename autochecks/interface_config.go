package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/web_server_for_network_devices/devcon"
)

// InterfaceConfig represents the xml formatted 'show configuration interfaces {INTERFACE}'
type InterfaceConfig struct {
	Configuration struct {
		Text            string `xml:",chardata"`
		CommitSeconds   string `xml:"commit-seconds,attr"`
		CommitLocaltime string `xml:"commit-localtime,attr"`
		CommitUser      string `xml:"commit-user,attr"`
		Comment         struct {
			Text string `xml:",chardata"`
		} `xml:"comment"`
		Interfaces struct {
			Text      string `xml:",chardata"`
			Interface struct {
				Text string `xml:",chardata"`
				Name struct {
					Text string `xml:",chardata"`
				} `xml:"name"`
				FlexibleVlanTagging struct {
					Text string `xml:",chardata"`
				} `xml:"flexible-vlan-tagging"`
				Unit []struct {
					Text string `xml:",chardata"`
					Name struct {
						Text string `xml:",chardata"`
					} `xml:"name"`
					Family struct {
						Text              string `xml:",chardata"`
						EthernetSwitching struct {
							Text          string `xml:",chardata"`
							InterfaceMode struct {
								Text string `xml:",chardata"`
							} `xml:"interface-mode"`
							Vlan struct {
								Text    string `xml:",chardata"`
								Members struct {
									Text string `xml:",chardata"`
								} `xml:"members"`
							} `xml:"vlan"`
						} `xml:"ethernet-switching"`
						Inet struct {
							Text    string `xml:",chardata"`
							Address struct {
								Text string `xml:",chardata"`
								Name struct {
									Text string `xml:",chardata"`
								} `xml:"name"`
							} `xml:"address"`
						} `xml:"inet"`
					} `xml:"family"`
					Description struct {
						Text string `xml:",chardata"`
					} `xml:"description"`
					VlanID struct {
						Text string `xml:",chardata"`
					} `xml:"vlan-id"`
				} `xml:"unit"`
			} `xml:"interface"`
		} `xml:"interfaces"`
	} `xml:"configuration"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner struct {
			Text string `xml:",chardata"`
		} `xml:"banner"`
	} `xml:"cli"`
}

// Mapper takes in a bytes.Buffer and returns a pointer to an
// InterfaceConfig object (struct)
func (cfg *InterfaceConfig) mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &cfg)
}

// Run Issues the command against the device and returns the object
func (cfg *InterfaceConfig) Run(p Params) (interface{}, error) {
	fwPolConfig := fmt.Sprintf("show configuration interfaces %v | display xml", p.Args[0])
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
