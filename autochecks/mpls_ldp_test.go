package autochecks

import "testing"

var (
	LDPACs = []AutoCheck{&LDPNeighborBB{}}
)

func TestLDPNeighbor(t *testing.T) {
	p := Params{
		IP: "10.1.1.56",
	}
	for _, ac := range LDPACs {
		_, err := ac.Run(p)
		if err != nil {
			t.Errorf("AC failed: %v", err)
		}
	}
}