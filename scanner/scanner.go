package main

import (
	"fmt"
	"log"
	"net"
	"sync"
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

func getVer(ip string) {
	fmt.Printf("Start processing %v\n", ip)
	var ac autochecks.SoftwareVersion
	p := autochecks.Params{
		IP: ip,
	}
	_, err := ac.Run(p)
	if err != nil {
		log.Printf("%v failed. %v", ip, err)
	}
	fmt.Println(ac.SoftwareInformation.HostName, ac.SoftwareInformation.JunosVersion)
	fmt.Printf("Finish processing %v\n", ip)
}

func updateDB(output chan autochecks.SoftwareVersion) {
	for ac := range output {
		fmt.Println(ac.SoftwareInformation.HostName)
		fmt.Println(ac.SoftwareInformation.JunosVersion)
	}
}

func main() {

	start := time.Now()
	workerPoolSize := 16
	// define channels
	ipStream := make(chan string)

	var wg sync.WaitGroup

	subnet, err := Hosts("10.1.1.0/24")
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < workerPoolSize; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for ip := range ipStream {
					getVer(ip)
				}
			}()
		}
	}()

	// Feeding the channel
	for _, ip := range subnet {
		ipStream <- ip
	}
	// close the input channel to signal we're done
	close(ipStream)
	// blocks until counter is back to zero
	wg.Wait()
	fmt.Printf("time elapsed: %v\n", time.Since(start))
}
