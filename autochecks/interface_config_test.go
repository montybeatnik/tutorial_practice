package autochecks

import "testing"

func TestInterfaceConfig(t *testing.T) {
	ac := InterfaceConfig{}
	a := []string{"ge-0/0/4"}
	p := Params{
		IP: "10.63.244.76",
		Args: a,
	}
	_, err := ac.Run(p)
	if err != nil {
		t.Errorf("AC failed: %v", err)
	}
}