package scan

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/montybeatnik/tutorial_practice/pkg/autochecks"
	"github.com/montybeatnik/tutorial_practice/pkg/models"
)

func checkError(e error) {
	if e != nil {
		log.Println(e)
	}
}

// Results holds the interesting bits of a scan.
type Results struct {
	Scanner   string
	Attempted int
	Failures  int
	Container []Result
}

// Result holds the interesting bits of an individual device query
type Result struct {
	Device models.Device
	Data   bytes.Buffer
	Error  error
}

// Telemetry keeps track of scan statistics.
type Telemetry struct {
	Attempted     int
	Failures      int
	ExecutionTime time.Duration
}

// Crawler will be an awesome interface in the future.
type Crawler interface {
	// Methods here
	SubnetScanner(subnet []string) Result
	ComplianceScanner()
	Rebaser()
}

// Devices takes in a slice of devices and an autocheck. It calls the run method
// on the autocheck and stores the results and any errors in the Result struct,
// which is passed into a channel for output to Stdout.
// Using the fan out concurrency pattern.
func Devices(devices []models.Device, ac autochecks.Check) {

	devCount := len(devices)
	ch := make(chan Result, len(devices))
	for d := 0; d < devCount; d++ {
		go func(dev models.Device) {
			// augmenting here
			p := autochecks.Params{
				IP: dev.Loopback,
			}
			buf, err := ac.Run(p)
			res := Result{
				Device: dev,
				Error:  err,
				Data:   buf,
			}
			ch <- res
		}(devices[d])
	}

	for devCount > 0 {
		result := <-ch
		devCount--

		var sw models.SoftwareVersion
		err := sw.Mapper(result.Data)
		checkError(err)
		if result.Error != nil {
			fmt.Println(sw.SoftwareInformation.HostName, sw.SoftwareInformation.JunosVersion)
		}
	}
	fmt.Println("-------------------------------------------------------------")
}

// Subnets takes in a slice of string slices and an autocheck interface.
// It returns a slice of Result.
// Closures, coupled with channels, are used to pass data from one goroutine
// to another.
func Subnets(subnet []string, ac autochecks.Check) {
	// Feeding the channel
	var allIPs int
	ch := make(chan Result, allIPs)
	allIPs += len(subnet)
	for _, ip := range subnet {
		go func(ip string) {
			d := models.Device{Loopback: ip}
			// augmenting here
			p := autochecks.Params{
				IP: ip,
			}
			buf, err := ac.Run(p)
			res := Result{
				Device: d,
				Error:  err,
				Data:   buf,
			}
			ch <- res
		}(ip)
	}
	for allIPs > 0 {
		result := <-ch
		allIPs--
		var sw models.SoftwareVersion
		sw.Mapper(result.Data)
		fmt.Println(result.Device.Loopback, result.Error, sw.SoftwareInformation.HostName, sw.SoftwareInformation.JunosVersion)
	}
}
