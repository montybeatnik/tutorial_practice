package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/montybeatnik/tutorial_practice/autochecks"
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

func getVer(ipStream chan string, output chan autochecks.SoftwareVersion) {

	for ip := range ipStream {
		var ac autochecks.SoftwareVersion
		p := autochecks.Params{
			IP: ip,
		}
		_, err := ac.Run(p)
		if err != nil {
			log.Printf("%v failed. %v", ip, err)
		}
		output <- ac
	}
}

func updateDB(output chan autochecks.SoftwareVersion) {
	for ac := range output {
		fmt.Println(ac.SoftwareInformation.HostName)
		fmt.Println(ac.SoftwareInformation.JunosVersion)
	}
}

func main() {
	// define channels
	ipStream := make(chan string)
	swVerStream := make(chan autochecks.SoftwareVersion)

	start := time.Now()
	subnet, err := Hosts("10.1.1.48/29")
	if err != nil {
		log.Fatal(err)
	}

	go getVer(ipStream, swVerStream)

	go updateDB(swVerStream)

	for _, ip := range subnet {
		// send ip onto the ipStream channel
		ipStream <- ip
	}
	fmt.Printf("time elapsed: %v\n", time.Since(start))
}
