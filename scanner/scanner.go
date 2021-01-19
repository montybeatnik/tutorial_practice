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

func getVer(ip string) Result {
	fmt.Printf("Start processing %v\n", ip)
	var ac autochecks.SoftwareVersion
	p := autochecks.Params{
		IP: ip,
	}
	_, err := ac.Run(p)
	res := Result{
		Hostname: ac.SoftwareInformation.HostName,
		IP:       ip,
		Version:  ac.SoftwareInformation.JunosVersion,
		Error:    err,
	}
	fmt.Printf("Finished processing %v\n", ip)
	return res
}

func updateDB(output chan autochecks.SoftwareVersion) {
	for ac := range output {
		fmt.Println(ac.SoftwareInformation.HostName)
		fmt.Println(ac.SoftwareInformation.JunosVersion)
	}
}

type Results struct {
	Container []Result
}

type Result struct {
	Hostname string
	IP       string
	Version  string
	Error    error
}

func Scanner() Results {

	workerPoolSize := 4
	// define channels
	ipStream := make(chan string)

	var wg sync.WaitGroup

	subnetOne, err := Hosts("10.63.244.0/28")
	if err != nil {
		log.Fatal(err)
	}
	subnetTwo, err := Hosts("10.63.244.16/28")
	if err != nil {
		log.Fatal(err)
	}

	var results Results
	wg.Add(1)
	// a closure to control the queue
	go func() {
		defer wg.Done()
		for i := 0; i < workerPoolSize; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for ip := range ipStream {
					res := getVer(ip)
					results.Container = append(results.Container, res)
				}
			}()
		}
	}()

	subnets := [][]string{subnetOne, subnetTwo}
	// Feeding the channel
	for _, subnet := range subnets {
		for _, ip := range subnet {
			ipStream <- ip
		}
	}
	// close the input channel to signal we're done
	close(ipStream)
	// blocks until counter is back to zero
	wg.Wait()
	return results
}

func main() {
	start := time.Now()
	results := Scanner()
	for _, res := range results.Container {
		if res.Error != nil {
			fmt.Println(res.IP, res.Error)
			continue
		}
		fmt.Println(res.Hostname, res.Version)
	}
	fmt.Printf("time elapsed: %v\n", time.Since(start))
}
