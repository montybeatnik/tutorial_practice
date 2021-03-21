package scan

import (
	"testing"

	"github.com/montybeatnik/tutorial_practice/autochecks"
)

func TestHosts(t *testing.T) {
	expected := 254
	hosts, err := Hosts("192.168.1.0/24")
	if err != nil {
		t.Errorf("Hosts call failed: %v", err)
	}
	if len(hosts) != 254 {
		t.Errorf("Expected %v, got: %v", expected, len(hosts))
	}
}

func BenchmarkHosts(b *testing.B) {
	// run the Hosts func b.N times
	for n := 0; n < b.N; n++ {
		Hosts("192.168.1.0/24")
	}
}

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
	Subnets(nets, &swVer)
}

// func BenchmarkSubnets(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		Subnets([][]string{"192.168.1.0/30"})
// 	}
// }
