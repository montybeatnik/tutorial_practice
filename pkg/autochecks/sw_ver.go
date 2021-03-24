package autochecks

import (
	"bytes"

	"github.com/montybeatnik/tutorial_practice/pkg/devcon"
	"github.com/pkg/errors"
)

var (
	softwareVerCMD = "show version | display xml"
)

type SoftwareVersion struct{}

// Run Issues the command against the device and returns the object
func (sv *SoftwareVersion) Run(p Params) (bytes.Buffer, error) {
	connInfo := devcon.ConnInfo{
		IP:      p.IP,
		Command: softwareVerCMD,
	}
	c := devcon.NewConfig(connInfo)
	buf, err := devcon.RunCmd(c)
	if err != nil {
		return buf, errors.Wrap(err, "reason")
	}
	return buf, nil
}
