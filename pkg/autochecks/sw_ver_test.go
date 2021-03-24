package autochecks

import (
	"testing"

	"github.com/montybeatnik/tutorial_practice/pkg/models"
)

var wrongVersion = "19.3R3.6"
var standardVersion = "17.3R3.10"

// TestSoftwareVersion runs the autocheck and ensures the output
// of the Junos Version matches our standard version
func TestSoftwareVersion(t *testing.T) {
	ac := SoftwareVersion{}
	p := Params{
		IP: "10.63.6.13",
	}
	buf, err := ac.Run(p)
	if err != nil {
		t.Errorf("AC failed: %v", err)
	}
	var sw models.SoftwareVersion
	if err := sw.Mapper(buf); err != nil {
		t.Errorf("problem mapping struct: %v", err)
	}
}
