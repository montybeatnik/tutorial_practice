package autochecks

import "testing"

var wrongVersion = "19.3R3.6"
var standardVersion = "17.3R3.10"

// TestSoftwareVersion runs the autocheck and ensures the output
// of the Junos Version matches our standard version
func TestSoftwareVersion(t *testing.T) {
	ac := SoftwareVersion{}
	p := Params{
		IP: "10.63.6.13",
	}
	sw, err := ac.Run(p)
	if err != nil {
		t.Errorf("AC failed: %v", err)
	}
	v := sw.(*SoftwareVersion)
	if v.SoftwareInformation.JunosVersion != standardVersion {
		t.Errorf("expected %q; got %q", standardVersion, v.SoftwareInformation.JunosVersion)
	}
}
