package autochecks

import "testing"

var (
	LSPACs = []AutoCheck{&LSPDetail{}}
)

func TestLSPDetail(t *testing.T) {
	for _, ac := range LSPACs {
		p := Params{
			IP: "10.63.6.13",
		}
		_, err := ac.Run(p)
		if err != nil {
			t.Errorf("AC failed: %v", err)
		}
	}
}