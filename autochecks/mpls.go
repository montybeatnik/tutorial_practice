package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/web_server_for_network_devices/devcon"
)

const (
	lspDetailCMD = "show mpls lsp detail | display xml"
)

type LSPDetail struct {
	MplsLspInformation struct {
		Text            string `xml:",chardata"`
		Xmlns           string `xml:"xmlns,attr"`
		RsvpSessionData []struct {
			Text        string `xml:",chardata"`
			SessionType string `xml:"session-type"`
			Count       string `xml:"count"`
			RsvpSession []struct {
				Text    string `xml:",chardata"`
				Style   string `xml:"style,attr"`
				MplsLsp struct {
					Text                 string `xml:",chardata"`
					DestinationAddress   string `xml:"destination-address"`
					SourceAddress        string `xml:"source-address"`
					LspState             string `xml:"lsp-state"`
					RouteCount           string `xml:"route-count"`
					Name                 string `xml:"name"`
					LspDescription       string `xml:"lsp-description"`
					ActivePath           string `xml:"active-path"`
					LspType              string `xml:"lsp-type"`
					EgressLabelOperation string `xml:"egress-label-operation"`
					LoadBalance          string `xml:"load-balance"`
					MplsLspAttributes    struct {
						Text                 string `xml:",chardata"`
						EncodingType         string `xml:"encoding-type"`
						SwitchingType        string `xml:"switching-type"`
						Gpid                 string `xml:"gpid"`
						MplsLspUpstreamLabel string `xml:"mpls-lsp-upstream-label"`
					} `xml:"mpls-lsp-attributes"`
					MplsLspPath struct {
						Text               string `xml:",chardata"`
						Title              string `xml:"title"`
						Name               string `xml:"name"`
						PathActive         string `xml:"path-active"`
						PathState          string `xml:"path-state"`
						SetupPriority      string `xml:"setup-priority"`
						HoldPriority       string `xml:"hold-priority"`
						OptimizeTimer      string `xml:"optimize-timer"`
						SmartOptimizeTimer string `xml:"smart-optimize-timer"`
						ReceivedRro        string `xml:"received-rro"`
					} `xml:"mpls-lsp-path"`
				} `xml:"mpls-lsp"`
				DestinationAddress string `xml:"destination-address"`
				SourceAddress      string `xml:"source-address"`
				LspState           string `xml:"lsp-state"`
				RouteCount         string `xml:"route-count"`
				Name               string `xml:"name"`
				LspPathType        string `xml:"lsp-path-type"`
				SuggestedLabelIn   string `xml:"suggested-label-in"`
				SuggestedLabelOut  string `xml:"suggested-label-out"`
				RecoveryLabelIn    string `xml:"recovery-label-in"`
				RecoveryLabelOut   string `xml:"recovery-label-out"`
				RsbCount           string `xml:"rsb-count"`
				ResvStyle          string `xml:"resv-style"`
				LabelIn            string `xml:"label-in"`
				LabelOut           string `xml:"label-out"`
				PsbLifetime        string `xml:"psb-lifetime"`
				PsbCreationTime    string `xml:"psb-creation-time"`
				SenderTspec        string `xml:"sender-tspec"`
				LspID              string `xml:"lsp-id"`
				TunnelID           string `xml:"tunnel-id"`
				ProtoID            string `xml:"proto-id"`
				LspAttributeFlags  string `xml:"lsp-attribute-flags"`
				PacketInformation  []struct {
					Text          string `xml:",chardata"`
					Heading       string `xml:"heading,attr"`
					PreviousHop   string `xml:"previous-hop"`
					InterfaceName string `xml:"interface-name"`
					Count         string `xml:"count"`
					NextHop       string `xml:"next-hop"`
					EntropyLabel  string `xml:"entropy-label"`
				} `xml:"packet-information"`
				Adspec      string `xml:"adspec"`
				RecordRoute struct {
					Text    string   `xml:",chardata"`
					Heading string   `xml:"heading,attr"`
					Address []string `xml:"address"`
					Self    string   `xml:"self"`
				} `xml:"record-route"`
			} `xml:"rsvp-session"`
			DisplayCount string `xml:"display-count"`
			UpCount      string `xml:"up-count"`
			DownCount    string `xml:"down-count"`
		} `xml:"rsvp-session-data"`
	} `xml:"mpls-lsp-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

// RouteSumMapper takes in a bytes.Buffer and returns a pointer to an
// RouteSum object (struct)
func (lspd *LSPDetail) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &lspd)
}

// Run Issues the command against the device and returns the object
func (lspd *LSPDetail) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: lspDetailCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return lspd, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	lspd.Mapper(buf)
	return lspd, nil
}
