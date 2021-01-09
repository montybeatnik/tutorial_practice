package autochecks

import "testing"

var (
	MacTableACs = []AutoCheck{&MacTable{}}
)

func TestMacTable(t *testing.T) {	
	for _, ac := range MacTableACs {
		a := []string{"BASE1_C"}
		p := Params{
			IP: "10.63.244.76",
			Args: a,
		}
		_, err := ac.Run(p)
		if err != nil {
			t.Errorf("AC failed: %v", err)
		}
	}
}