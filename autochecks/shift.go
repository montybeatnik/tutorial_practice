package autochecks

// import (
// 	"bytes"
// 	"encoding/xml"
// 	"fmt"
// 	"log"
// )

// // LDPNeighborBB  holds the data of the "show ldp neighbor logical-system BACKBONE | display xml" commmand
// type Shift struct {
// 	LdpNeighborInformation struct {
// 		LdpNeighbor []struct {
// 			Text               string `xml:",chardata"`
// 			LdpNeighborAddress string `xml:"ldp-neighbor-address"`
// 			InterfaceName      string `xml:"interface-name"`
// 			LdpLabelSpaceID    string `xml:"ldp-label-space-id"`
// 			// LdpRemainingTime   string `xml:"ldp-remaining-time"`
// 		} `xml:"ldp-neighbor"`
// 	} `xml:"ldp-neighbor-information"`
// 	Cli struct {
// 		Text   string `xml:",chardata"`
// 		Banner string `xml:"banner"`
// 	} `xml:"cli"`
// }

// // Mapper takes in a bytes.Buffer and returns a pointer to an
// // IsisLDPNeighborBB object (struct)
// func (ldpNeigh *LDPNeighborBB) Mapper(b bytes.Buffer) {
// 	xml.Unmarshal(b.Bytes(), &ldpNeigh)
// }

// // Run Issues the command against the device and returns the object
// func (ldpNeigh *LDPNeighborBB) Run(ip string) error {
// 	buf, err := RunCmd(ip, "show ldp neighbor logical-system BACKBONE | display xml")
// 	if err != nil {
// 		return fmt.Errorf("Problem getting %v output!: %w", ip, err)
// 	}
// 	ldpNeigh.Mapper(buf)
// }

// func (ldpNeigh *LDPNeighborBB) Parse(db sql.DB, v interface{}) {

// 	for _, adj := range ldpNeigh.LdpNeighborInformation.LdpNeighbor {
// 		fmt.Println(adj.InterfaceName)
// 		fmt.Println(adj.LdpNeighborAddress)
// 		fmt.Println(adj.LdpLabelSpaceID)
// 	}
// }

// type Result struct {
// 	MisMatched     bool
// 	Got            string
// 	Expected       string
// 	Recommendation string
// }

// // CompareLDPNeighborsSameInterface checks if the interface name is the same in the
// // pre|post output.
// // If the LDP neighbor addresses are different, a mismatch has occurred.
// func CompareLDPNeighborsSameInterface(preCheck, postCheck *LDPNeighborBB) Result {
// 	var res Result
// 	for _, pre := range preCheck.LdpNeighborInformation.LdpNeighbor {
// 		for _, post := range postCheck.LdpNeighborInformation.LdpNeighbor {
// 			if pre.InterfaceName == post.InterfaceName {
// 				if pre.LdpNeighborAddress != post.LdpNeighborAddress {
// 					res.MisMatched = true
// 					res.Got = post.LdpNeighborAddress
// 					res.Expected = pre.LdpNeighborAddress
// 					res.Recommendation = "check port configuration!"
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// // LDPRouteBB  holds the data of the "show ldp route logical-system BACKBONE | display xml" commmand
// type LDPRouteBB struct {
// 	XMLName             xml.Name `xml:"rpc-reply"`
// 	Text                string   `xml:",chardata"`
// 	Junos               string   `xml:"junos,attr"`
// 	LdpRouteInformation struct {
// 		Text     string `xml:",chardata"`
// 		Xmlns    string `xml:"xmlns,attr"`
// 		LdpRoute []struct {
// 			Text       string `xml:",chardata"`
// 			LdpPrefix  string `xml:"ldp-prefix"`
// 			LdpNexthop struct {
// 				Text             string `xml:",chardata"`
// 				InterfaceName    string `xml:"interface-name"`
// 				InterfaceAddress string `xml:"interface-address"`
// 			} `xml:"ldp-nexthop"`
// 		} `xml:"ldp-route"`
// 	} `xml:"ldp-route-information"`
// 	Cli struct {
// 		Text   string `xml:",chardata"`
// 		Banner string `xml:"banner"`
// 	} `xml:"cli"`
// }

// // Mapper takes in a bytes.Buffer and returns a pointer to an
// // IsisLDPRouteBB object (struct)
// func (ldpRoute *LDPRouteBB) Mapper(b bytes.Buffer) {
// 	xml.Unmarshal(b.Bytes(), &ldpRoute)
// }

// // Run Issues the command against the device and returns the object
// func (ldpRoute *LDPRouteBB) Run(ip string) error {
// 	buf, err := RunCmd(ip, "show ldp route logical-system BACKBONE | display xml")
// 	if err != nil {
// 		return fmt.Errorf("Problem getting %v output!: %w", ip, err)
// 	}
// 	ldpRoute.Mapper(buf)
// }

// func (ldpRoute *LDPRouteBB) Parse(db sql.DB, v interface{}) {

// 	for _, adj := range ldpRoute.LdpRouteInformation.LdpRoute {
// 		fmt.Println(adj.LdpPrefix)
// 		fmt.Println(adj.LdpNexthop)
// 		fmt.Println(adj.LdpNexthop.InterfaceName)
// 		fmt.Println(adj.LdpNexthop.InterfaceAddress)
// 	}
// }

// // COMPARE FUNCTIONS
// // CompareLDPSameNeighborsDifferentInterface
// // TODO: finish the logic here!!!
// func CompareLDPSameNeighborsDifferentInterface(preCheck, postCheck *LDPNeighborBB) Result {
// 	var res Result
// 	for _, post := range postCheck.LdpNeighborInformation.LdpNeighbor {
// 		for _, pre := range preCheck.LdpNeighborInformation.LdpNeighbor {
// 			if pre.LdpNeighborAddress == post.LdpNeighborAddress {
// 				if pre.LdpNeighborAddress != post.LdpNeighborAddress {
// 					res.MisMatched = true
// 					res.Got = post.LdpNeighborAddress
// 					res.Expected = pre.LdpNeighborAddress
// 					res.Recommendation = "check port configuration!"
// 				}
// 			}
// 		}
// 	}
// 	return res
// }
