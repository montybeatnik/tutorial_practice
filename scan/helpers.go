package scan

import (
	"fmt"
	"net"

	"github.com/montybeatnik/tutorial_practice/autochecks"
)

// Hosts takes in a subnet and returns the host addresses.
func Hosts(subnet string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// remove network address and broadcast address
	lenIPs := len(ips)
	switch {
	case lenIPs < 2:
		return ips, nil

	default:
		return ips[1 : len(ips)-1], nil
	}
}

// inc moves the IP address forward in the subnet.
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// func getVer(ver autochecks.SoftwareVersion) Result {
// 	res := Result{
// 		Hostname: ac.SoftwareInformation.HostName,
// 		IP:       ,
// 		Version:  ac.SoftwareInformation.JunosVersion,
// 		Error:    err,
// 	}
// 	return res
// }

func updateDB(output chan autochecks.SoftwareVersion) {
	for ac := range output {
		fmt.Println(ac.SoftwareInformation.HostName)
		fmt.Println(ac.SoftwareInformation.JunosVersion)
	}
}
