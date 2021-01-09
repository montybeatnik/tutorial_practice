package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/tutorial_practice/devcon"
)

const (
	isisAdjCMD      = "show isis adjacency | display xml"
	isisHostnameCMD = "show isis hostname | display xml"
)

// IsisAdjacency  holds the data of the "show isis adjacency" commmand
type ISISAdjacency struct {
	IsisAdjacencyInformation struct {
		Text          string `xml:",chardata"`
		Xmlns         string `xml:"xmlns,attr"`
		Style         string `xml:"style,attr"`
		IsisAdjacency []struct {
			Text           string `xml:",chardata"`
			InterfaceName  string `xml:"interface-name"`
			SystemName     string `xml:"system-name"`
			Level          string `xml:"level"`
			AdjacencyState string `xml:"adjacency-state"`
			Holdtime       string `xml:"holdtime"`
		} `xml:"isis-adjacency"`
	} `xml:"isis-adjacency-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

// IsisAdjacencyMapper takes in a bytes.Buffer and returns a pointer to an
// IsisAdjacency object (struct)
func (isAdj *ISISAdjacency) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &isAdj)
}

// Run Issues the command against the device and returns the object
func (isAdj *ISISAdjacency) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: isisAdjCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return isAdj, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	isAdj.Mapper(buf)
	return isAdj, nil
}

// ISISAdjacencyBB
// IsisAdjacency  holds the data of the "show isis adjacency" commmand
type ISISAdjacencyBB struct {
	IsisAdjacencyInformation struct {
		Text          string `xml:",chardata"`
		Xmlns         string `xml:"xmlns,attr"`
		Style         string `xml:"style,attr"`
		IsisAdjacency []struct {
			Text           string `xml:",chardata"`
			InterfaceName  string `xml:"interface-name"`
			SystemName     string `xml:"system-name"`
			Level          string `xml:"level"`
			AdjacencyState string `xml:"adjacency-state"`
			Holdtime       string `xml:"holdtime"`
		} `xml:"isis-adjacency"`
	} `xml:"isis-adjacency-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

// IsisAdjacencyMapper takes in a bytes.Buffer and returns a pointer to an
// IsisAdjacency object (struct)
func (isAdj *ISISAdjacencyBB) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &isAdj)
}

// Run Issues the command against the device and returns the object
func (isAdj *ISISAdjacencyBB) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: isisAdjCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return isAdj, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	isAdj.Mapper(buf)
	return isAdj, nil
}

type ISISHostname struct {
	NumberOfDevices         int
	IsisHostnameInformation struct {
		Text         string `xml:",chardata"`
		Xmlns        string `xml:"xmlns,attr"`
		IsisHostname []struct {
			Text             string `xml:",chardata"`
			SystemID         string `xml:"system-id"`
			SystemName       string `xml:"system-name"`
			IsisHostnameType string `xml:"isis-hostname-type"`
		} `xml:"isis-hostname"`
	} `xml:"isis-hostname-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

func (isisHostnamej *ISISHostname) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &isisHostnamej)
}

// Run Issues the command against the device and returns the object
func (isisHostnamej *ISISHostname) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: isisHostnameCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return isisHostnamej, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	isisHostnamej.Mapper(buf)
	return isisHostnamej, nil
}

// func (isAdj *ISISAdjacencyBB) Parse(db sql.DB, v interface{}) {

// 	for _, adj := range isAdj.IsisAdjacencyInformation.IsisAdjacency {
// 		fmt.Println(adj.InterfaceName)
// 		fmt.Println(adj.SystemName)
// 		fmt.Println(adj.Level)
// 	}
// }

// func (isAdj *ISISAdjacencyBB) Compare() {}
