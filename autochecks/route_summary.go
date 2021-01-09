package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/web_server_for_network_devices/devcon"
)

const (
	rtSumCMD = "show route summary | display xml"
)

// RouteSum Runs the "show route summary | display xml" command against a JunOS device.
type RouteSummary struct {
	RouteSummaryInformation struct {
		Text       string `xml:",chardata"`
		Xmlns      string `xml:"xmlns,attr"`
		AsNumber   string `xml:"as-number"`
		RouterID   string `xml:"router-id"`
		RouteTable []struct {
			Text               string `xml:",chardata"`
			TableName          string `xml:"table-name"`
			DestinationCount   string `xml:"destination-count"`
			TotalRouteCount    string `xml:"total-route-count"`
			ActiveRouteCount   string `xml:"active-route-count"`
			HolddownRouteCount string `xml:"holddown-route-count"`
			HiddenRouteCount   string `xml:"hidden-route-count"`
			Protocols          []struct {
				Text               string `xml:",chardata"`
				ProtocolName       string `xml:"protocol-name"`
				ProtocolRouteCount string `xml:"protocol-route-count"`
				ActiveRouteCount   string `xml:"active-route-count"`
			} `xml:"protocols"`
		} `xml:"route-table"`
	} `xml:"route-summary-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

// RouteSumMapper takes in a bytes.Buffer and returns a pointer to an
// RouteSum object (struct)
func (rts *RouteSummary) mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &rts)
}

// Run Issues the command against the device and returns the object
func (rts *RouteSummary) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: rtSumCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return rts, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	rts.mapper(buf)
	return rts, nil
}
