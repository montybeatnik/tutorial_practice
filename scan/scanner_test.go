package scan

import (
	"testing"

	"github.com/montybeatnik/tutorial_practice/autochecks"
)

func TestSubnet(t *testing.T) {
	netOne := "192.168.1.0/30"
	netTwo := "192.168.1.0/30"
	hostsOne, err := Hosts(netOne)
	if err != nil {
		t.Errorf("Hosts failed: %v", err)
	}
	hostsTwo, err := Hosts(netTwo)
	if err != nil {
		t.Errorf("Hosts failed: %v", err)
	}
	nets := [][]string{hostsOne, hostsTwo}
	var swVer autochecks.SoftwareVersion
	results := Subnets(nets, &swVer)
	for _, r := range results {
		t.Log(r.Error.Error())
	}
}

func BenchScanner(b *testing.B) {

}
