package autochecks

import "testing"

var (
	FWPolCfgACs = []AutoCheck{&SecFWPol{}}
)

func TestFWSecPol(t *testing.T) {
	for _, ac := range FWPolCfgACs {
		p := Params{
			IP: "10.1.1.59",
		}
		_, err := ac.Run(p)
		if err != nil {
			t.Errorf("AC failed: %v", err)
		}
	}
}