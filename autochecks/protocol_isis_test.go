package autochecks

import "testing"

var (
	ACs = []AutoCheck{&ISISAdjacency{},&ISISAdjacencyBB{}}
)

func TestISISAdj(t *testing.T) {
	for _, ac := range ACs {
		p := Params{
			IP: "10.63.6.13",
		}
		_, err := ac.Run(p)
		if err != nil {
			t.Errorf("AC failed: %v", err)
		}
	}
}