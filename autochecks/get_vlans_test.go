package autochecks

import "testing"

var (
	VlansACs = []AutoCheck{&Vlans{}}
)

func TestVlans(t *testing.T) {	
	for _, ac := range VlansACs {
		p := Params{
			IP: "10.63.244.76",
		}
		_, err := ac.Run(p)
		if err != nil {
			t.Errorf("AC failed: %v", err)
		}
	}
}