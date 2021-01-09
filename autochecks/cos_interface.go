package autochecks

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/montybeatnik/tutorial_practice/devcon"
)

type COSInterfaceACX struct {
	CosInterfaceInformation struct {
		Text         string `xml:",chardata"`
		Xmlns        string `xml:"xmlns,attr"`
		InterfaceMap struct {
			Text                               string `xml:",chardata"`
			InterfaceName                      string `xml:"interface-name"`
			InterfaceIndex                     string `xml:"interface-index"`
			InterfaceQueuesSupported           string `xml:"interface-queues-supported"`
			InterfaceQueuesInUse               string `xml:"interface-queues-in-use"`
			InterfaceShapingRate               string `xml:"interface-shaping-rate"`
			SchedulerMapName                   string `xml:"scheduler-map-name"`
			SchedulerMapIndex                  string `xml:"scheduler-map-index"`
			InterfaceCongestionNotificationMap string `xml:"interface-congestion-notification-map"`
			CosObjects                         struct {
				Text             string   `xml:",chardata"`
				CosObjectType    []string `xml:"cos-object-type"`
				CosObjectName    []string `xml:"cos-object-name"`
				CosObjectSubtype []string `xml:"cos-object-subtype"`
				CosObjectIndex   []string `xml:"cos-object-index"`
			} `xml:"cos-objects"`
			ForwardingClassSetAttachment string `xml:"forwarding-class-set-attachment"`
			ILogicalMap                  []struct {
				Text          string `xml:",chardata"`
				ILogicalName  string `xml:"i-logical-name"`
				ILogicalIndex string `xml:"i-logical-index"`
				CosObjects    struct {
					Text             string `xml:",chardata"`
					CosObjectType    string `xml:"cos-object-type"`
					CosObjectName    string `xml:"cos-object-name"`
					CosObjectSubtype string `xml:"cos-object-subtype"`
					CosObjectIndex   string `xml:"cos-object-index"`
				} `xml:"cos-objects"`
			} `xml:"i-logical-map"`
		} `xml:"interface-map"`
	} `xml:"cos-interface-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

func (it *COSInterfaceACX) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &it)
}

func (it *COSInterfaceACX) Run(p Params) (interface{}, error) {
	cmd := fmt.Sprintf("show class-of-service interface %v | display xml", p.Args[0])
	conInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: cmd,
	}
	c := devcon.NewConfig(conInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return it, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	it.Mapper(buf)
	return it, nil
}

type COSInterfaceMX struct {
	CosInterfaceInformation struct {
		Text         string `xml:",chardata"`
		Xmlns        string `xml:"xmlns,attr"`
		InterfaceMap struct {
			Text        string `xml:",chardata"`
			ILogicalMap struct {
				Text          string `xml:",chardata"`
				ILogicalName  string `xml:"i-logical-name"`
				ILogicalIndex string `xml:"i-logical-index"`
				ShapingRate   string `xml:"shaping-rate"`
				CosObjects    struct {
					Text             string   `xml:",chardata"`
					CosObjectType    []string `xml:"cos-object-type"`
					CosObjectName    []string `xml:"cos-object-name"`
					CosObjectSubtype []string `xml:"cos-object-subtype"`
					CosObjectIndex   []string `xml:"cos-object-index"`
				} `xml:"cos-objects"`
			} `xml:"i-logical-map"`
		} `xml:"interface-map"`
	} `xml:"cos-interface-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

func (it *COSInterfaceMX) Mapper(b bytes.Buffer) {
	xml.Unmarshal(b.Bytes(), &it)
}

func (it *COSInterfaceMX) Run(p Params) (interface{}, error) {
	cmd := fmt.Sprintf("show class-of-service interface %s %s| display xml", p.Args[0], p.Args[1])
	conInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: cmd,
	}
	c := devcon.NewConfig(conInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return it, fmt.Errorf("Problem getting %v output!: %w", p.IP, err)
	}
	it.Mapper(buf)
	return it, nil
}
