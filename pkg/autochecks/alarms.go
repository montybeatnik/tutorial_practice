package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/tutorial_practice/pkg/devcon"
)

const (
	chassisAlarmsCMD = "show chassis alarms | display xml"
	systemAlarmsCMD  = "show system alarms| display xml"
)

// ChassisAlarms  holds the data of the "show chassis alarms" commmand
type ChassisAlarms struct {
	AlarmInformation struct {
		Text         string `xml:",chardata"`
		Xmlns        string `xml:"xmlns,attr"`
		AlarmSummary struct {
			Text             string `xml:",chardata"`
			ActiveAlarmCount string `xml:"active-alarm-count"`
		} `xml:"alarm-summary"`
		AlarmDetail []struct {
			Text      string `xml:",chardata"`
			AlarmTime struct {
				Text    string `xml:",chardata"`
				Seconds string `xml:"seconds,attr"`
			} `xml:"alarm-time"`
			AlarmClass            string `xml:"alarm-class"`
			AlarmDescription      string `xml:"alarm-description"`
			AlarmShortDescription string `xml:"alarm-short-description"`
			AlarmType             string `xml:"alarm-type"`
		} `xml:"alarm-detail"`
	} `xml:"alarm-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

// IsalmacencyMapper takes in a bytes.Buffer and returns a pointer to an
// Isalmacency object (struct)
func (alm *ChassisAlarms) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &alm)
}

// Run Issues the command against the device and returns the object
func (alm *ChassisAlarms) Run(p Params) (bytes.Buffer, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: chassisAlarmsCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return buf, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	return buf, nil
}

// SystemAlarms  holds the data of the "show system alarms" commmand
type SystemAlarms struct {
	AlarmInformation struct {
		Text         string `xml:",chardata"`
		Xmlns        string `xml:"xmlns,attr"`
		AlarmSummary struct {
			Text             string `xml:",chardata"`
			ActiveAlarmCount string `xml:"active-alarm-count"`
		} `xml:"alarm-summary"`
		AlarmDetail []struct {
			Text      string `xml:",chardata"`
			AlarmTime struct {
				Text    string `xml:",chardata"`
				Seconds string `xml:"seconds,attr"`
			} `xml:"alarm-time"`
			AlarmClass            string `xml:"alarm-class"`
			AlarmDescription      string `xml:"alarm-description"`
			AlarmShortDescription string `xml:"alarm-short-description"`
			AlarmType             string `xml:"alarm-type"`
		} `xml:"alarm-detail"`
	} `xml:"alarm-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

// almMapper takes in a bytes.Buffer and returns a pointer to an
// almMapper object (struct)
func (alm *SystemAlarms) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &alm)
}

// Run Issues the command against the device and returns the object
func (alm *SystemAlarms) Run(p Params) (interface{}, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: systemAlarmsCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return alm, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	alm.Mapper(buf)
	return alm, nil
}
