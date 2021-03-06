package autochecks

import "testing"

func TestInterfaceTerse(t *testing.T) {
	ac := InterfaceTerse{}
	p := Params{
		IP: "10.63.6.13",
	}
	_, err := ac.Run(p)
	if err != nil {
		t.Errorf("AC failed: %v", err)
	}
}

func TestInterfaceDescription(t *testing.T) {
	ac := InterfaceDescription{}
	p := Params{
		IP: "10.63.6.13",
	}
	_, err := ac.Run(p)
	if err != nil {
		t.Errorf("AC failed: %v", err)
	}
}