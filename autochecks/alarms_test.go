package autochecks

import "testing"

func TestChassisAlarms(t *testing.T) {
	ac := ChassisAlarms{}
	p := Params{
		IP: "10.63.6.13",
	}
	_, err := ac.Run(p)
	if err != nil {
		t.Errorf("AC failed: %v", err)
	}
}

func TestSystemAlarms(t *testing.T) {
	ac := SystemAlarms{}
	p := Params{
		IP: "10.63.6.13",
	}
	_, err := ac.Run(p)
	if err != nil {
		t.Errorf("AC failed: %v", err)
	}
}
