package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/web_server_for_network_devices/devcon"
)

//BGPSummary (master)
type BGPSummary struct {
	BgpInformation struct {
		Text          string `xml:",chardata"`
		Xmlns         string `xml:"xmlns,attr"`
		GroupCount    string `xml:"group-count"`
		PeerCount     string `xml:"peer-count"`
		DownPeerCount string `xml:"down-peer-count"`
		BgpRib        []struct {
			Text                          string `xml:",chardata"`
			Style                         string `xml:"style,attr"`
			Name                          string `xml:"name"`
			TotalPrefixCount              string `xml:"total-prefix-count"`
			ReceivedPrefixCount           string `xml:"received-prefix-count"`
			AcceptedPrefixCount           string `xml:"accepted-prefix-count"`
			ActivePrefixCount             string `xml:"active-prefix-count"`
			SuppressedPrefixCount         string `xml:"suppressed-prefix-count"`
			HistoryPrefixCount            string `xml:"history-prefix-count"`
			DampedPrefixCount             string `xml:"damped-prefix-count"`
			TotalExternalPrefixCount      string `xml:"total-external-prefix-count"`
			ActiveExternalPrefixCount     string `xml:"active-external-prefix-count"`
			AcceptedExternalPrefixCount   string `xml:"accepted-external-prefix-count"`
			SuppressedExternalPrefixCount string `xml:"suppressed-external-prefix-count"`
			TotalInternalPrefixCount      string `xml:"total-internal-prefix-count"`
			ActiveInternalPrefixCount     string `xml:"active-internal-prefix-count"`
			AcceptedInternalPrefixCount   string `xml:"accepted-internal-prefix-count"`
			SuppressedInternalPrefixCount string `xml:"suppressed-internal-prefix-count"`
			PendingPrefixCount            string `xml:"pending-prefix-count"`
			BgpRibState                   string `xml:"bgp-rib-state"`
			VpnRibState                   string `xml:"vpn-rib-state"`
		} `xml:"bgp-rib"`
		BgpPeer []struct {
			Text            string `xml:",chardata"`
			Style           string `xml:"style,attr"`
			Heading         string `xml:"heading,attr"`
			PeerAddress     string `xml:"peer-address"`
			PeerAs          string `xml:"peer-as"`
			InputMessages   string `xml:"input-messages"`
			OutputMessages  string `xml:"output-messages"`
			RouteQueueCount string `xml:"route-queue-count"`
			FlapCount       string `xml:"flap-count"`
			ElapsedTime     struct {
				Text    string `xml:",chardata"`
				Seconds string `xml:"seconds,attr"`
			} `xml:"elapsed-time"`
			Description string `xml:"description"`
			PeerState   struct {
				Text   string `xml:",chardata"`
				Format string `xml:"format,attr"`
			} `xml:"peer-state"`
			BgpRib []struct {
				Text                  string `xml:",chardata"`
				Style                 string `xml:"style,attr"`
				Name                  string `xml:"name"`
				ActivePrefixCount     string `xml:"active-prefix-count"`
				ReceivedPrefixCount   string `xml:"received-prefix-count"`
				AcceptedPrefixCount   string `xml:"accepted-prefix-count"`
				SuppressedPrefixCount string `xml:"suppressed-prefix-count"`
			} `xml:"bgp-rib"`
		} `xml:"bgp-peer"`
	} `xml:"bgp-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

// Mapper takes in a bytes.Buffer and returns a pointer to an RouteSum object (struct)
func (bgps *BGPSummary) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &bgps)
}

// Run constructs a NewConfig and uses it to
// issue the command against the device.
// It returns a pointer to a BGPSummary struct and a nil error if all goes
// as planned. Else, it returns an empty BGPSummary
func (bgps *BGPSummary) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: "show bgp summary | display xml",
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return bgps, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	bgps.Mapper(buf)
	return bgps, nil
}

// // AdvRouteBGP runs "show route advertising-protocol bgp {prefix} | display xml"
// type AdvRouteBGP struct {
// 	RouteInformation struct {
// 		Text       string `xml:",chardata"`
// 		Xmlns      string `xml:"xmlns,attr"`
// 		RouteTable []struct {
// 			Text               string `xml:",chardata"`
// 			TableName          string `xml:"table-name"`
// 			DestinationCount   string `xml:"destination-count"`
// 			TotalRouteCount    string `xml:"total-route-count"`
// 			ActiveRouteCount   string `xml:"active-route-count"`
// 			HolddownRouteCount string `xml:"holddown-route-count"`
// 			HiddenRouteCount   string `xml:"hidden-route-count"`
// 			Rt                 []struct {
// 				Text          string `xml:",chardata"`
// 				Style         string `xml:"style,attr"`
// 				RtDestination string `xml:"rt-destination"`
// 				RtEntry       struct {
// 					Text            string `xml:",chardata"`
// 					ActiveTag       string `xml:"active-tag"`
// 					ProtocolName    string `xml:"protocol-name"`
// 					BgpMetricFlags  string `xml:"bgp-metric-flags"`
// 					LocalPreference string `xml:"local-preference"`
// 					AsPath          string `xml:"as-path"`
// 					Nh              struct {
// 						Text string `xml:",chardata"`
// 						To   string `xml:"to"`
// 					} `xml:"nh"`
// 					Med string `xml:"med"`
// 				} `xml:"rt-entry"`
// 				RtPrefixLength struct {
// 					Text string `xml:",chardata"`
// 					Emit string `xml:"emit,attr"`
// 				} `xml:"rt-prefix-length"`
// 			} `xml:"rt"`
// 		} `xml:"route-table"`
// 	} `xml:"route-information"`
// 	Cli struct {
// 		Text   string `xml:",chardata"`
// 		Banner string `xml:"banner"`
// 	} `xml:"cli"`
// }

// // Mapper takes in a bytes.Buffer and returns a pointer to an RouteSum object (struct)
// func (advRt *AdvRouteBGP) mapper(b bytes.Buffer) {
// 	xml.Unmarshal(b.Bytes(), &advRt)
// }

// // Run uses the BGP Summary output to collect advertised routes from all peers.
// func (advRt *AdvRouteBGP) Run(ip string) error {
// 	// establish a BGPSummary instance
// 	bgpsum := BGPSummary{}
// 	// Run the bgp summary command
// 	bgpsum.Run(ip)
// 	// iterate over the bgp peers,
// 	// grabbing the advertised routes,
// 	// and printing the table name and the destination count to std out.
// 	for _, neigh := range bgpsum.BgpInformation.BgpPeer {
// 		cmd := fmt.Sprintf("show route advertising-protocol bgp %v | display xml", neigh.PeerAddress)
// 		buf, err := RunCmd(ip, cmd)
// 		if err != nil {
// 			return fmt.Errorf("Problem getting %v output!: %w", ip, err)
// 		}
// 		advRt.mapper(buf)
// 		fmt.Println("table info for:", neigh.PeerAddress)
// 		for _, table := range advRt.RouteInformation.RouteTable {
// 			fmt.Println("table:", table.TableName)
// 			fmt.Println("table dst count:", table.DestinationCount)
// 		}
// 		fmt.Println(strings.Repeat("#", 40))
// 	}
// 	return nil
// }

// func (advRt *AdvRouteBGP) Parse(db sql.DB, v interface{}) {}

// // RecRouteBGP runs "show route advertising-protocol bgp {prefix} | display xml"
// type RecRouteBGP struct {
// 	XMLName          xml.Name `xml:"rpc-reply"`
// 	Text             string   `xml:",chardata"`
// 	Junos            string   `xml:"junos,attr"`
// 	RouteInformation struct {
// 		Text       string `xml:",chardata"`
// 		Xmlns      string `xml:"xmlns,attr"`
// 		RouteTable []struct {
// 			Text               string `xml:",chardata"`
// 			TableName          string `xml:"table-name"`
// 			DestinationCount   string `xml:"destination-count"`
// 			TotalRouteCount    string `xml:"total-route-count"`
// 			ActiveRouteCount   string `xml:"active-route-count"`
// 			HolddownRouteCount string `xml:"holddown-route-count"`
// 			HiddenRouteCount   string `xml:"hidden-route-count"`
// 			Rt                 []struct {
// 				Text          string `xml:",chardata"`
// 				Style         string `xml:"style,attr"`
// 				RtDestination string `xml:"rt-destination"`
// 				RtEntry       struct {
// 					Text            string `xml:",chardata"`
// 					ActiveTag       string `xml:"active-tag"`
// 					ProtocolName    string `xml:"protocol-name"`
// 					LocalPreference string `xml:"local-preference"`
// 					AsPath          string `xml:"as-path"`
// 					Nh              struct {
// 						Text string `xml:",chardata"`
// 						To   string `xml:"to"`
// 					} `xml:"nh"`
// 					Med string `xml:"med"`
// 				} `xml:"rt-entry"`
// 				RtPrefixLength struct {
// 					Text string `xml:",chardata"`
// 					Emit string `xml:"emit,attr"`
// 				} `xml:"rt-prefix-length"`
// 			} `xml:"rt"`
// 		} `xml:"route-table"`
// 	} `xml:"route-information"`
// 	Cli struct {
// 		Text   string `xml:",chardata"`
// 		Banner string `xml:"banner"`
// 	} `xml:"cli"`
// }

// // Mapper takes in a bytes.Buffer and returns a pointer to an RouteSum object (struct)
// func (recRt *RecRouteBGP) mapper(b bytes.Buffer) {
// 	xml.Unmarshal(b.Bytes(), &recRt)
// }

// // Run uses the BGP Summary output to collect advertised routes from all peers.
// func (recRt *RecRouteBGP) Run(ip string) error {

// 	var wg sync.WaitGroup
// 	// establish a BGPSummary instance
// 	bgpsum := BGPSummary{}
// 	// Run the bgp summary command
// 	bgpsum.Run(ip)
// 	// iterate over the bgp peers,
// 	// grabbing the received routes,
// 	// and printing the table name and the destination count to std out.
// 	for _, neigh := range bgpsum.BgpInformation.BgpPeer {

// 		wg.Add(1)
// 		// goroutine anon func
// 		go func() {
// 			defer wg.Done()
// 			cmd := fmt.Sprintf("show route receive-protocol bgp %v | display xml", neigh.PeerAddress)
// 			buf, err := RunCmd(ip, cmd)
// 			if err != nil {
// 				// return fmt.Errorf("Problem getting output!: %w", err)
// 				log.Println("Problem getting output!: %w", err)
// 			}
// 			recRt.mapper(buf)
// 			fmt.Println("table info for:", neigh.PeerAddress)
// 			for _, table := range recRt.RouteInformation.RouteTable {
// 				fmt.Println("table:", table.TableName)
// 				fmt.Println("table dst count:", table.DestinationCount)
// 			}
// 			fmt.Println(strings.Repeat("#", 40))
// 		}()

// 	}
// 	wg.Wait()
// 	return nil
// }

// func (recRt *RecRouteBGP) Parse(db sql.DB, v interface{}) {}
