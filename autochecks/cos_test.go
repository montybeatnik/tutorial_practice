package autochecks

import "testing"

func TestCOSInterface(t *testing.T) {
	ac := COSInterfaceACX{}
	a := []string{"ge-0/1/0"}
	p := Params{
		IP: "10.63.247.190",
		Args: a,
	}
	_, err := ac.Run(p)
	if err != nil {
		t.Errorf("AC failed: %v", err)
	}
}
