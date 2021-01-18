package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/montybeatnik/tutorial_practice/autochecks"
	"github.com/pkg/errors"
)

func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
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

func getVer(ip string) (autochecks.SoftwareVersion, error) {
	var ac autochecks.SoftwareVersion
	p := autochecks.Params{
		IP: ip,
	}
	_, err := ac.Run(p)
	if err != nil {
		return ac, errors.Wrap(err, "autocheck failed")
	}
	return ac, nil
}

func updateDB(v autochecks.SoftwareVersion) error {
	fmt.Println("Some table updated here...")
	output := v.SoftwareInformation
	fmt.Println(output.HostName)
	fmt.Println(output.JunosVersion)
	fmt.Println("done")
	return nil
}

func main() {
	start := time.Now()
	tmp, err := Hosts("10.1.1.48/29")
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range tmp {
		scanner := bufio.NewScanner(strings.NewReader(ip))
		fmt.Println(ip)
		for scanner.Scan() {
			ver, err := getVer(ip)
			if err != nil {
				log.Println(err)
				continue
			}
			updateDB(ver)
		}
	}
	fmt.Printf("time elapsed: %v\n", time.Since(start))
}
