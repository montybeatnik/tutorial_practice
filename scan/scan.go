package scan

import (
	"sync"
	"time"

	"github.com/montybeatnik/tutorial_practice/autochecks"
	"github.com/montybeatnik/tutorial_practice/models"
)

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
	Data   autochecks.Check
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
// which is what it returns.
// Any type assertions and are left to the caller for futher processing.
func Devices(devices []models.Device, ac autochecks.Check) []Result {
	// define the queue size
	workerPoolSize := 64 // 128
	// define channels
	deviceStream := make(chan models.Device)
	queue := make(chan Result)
	// wait group for synchronicity
	var wg sync.WaitGroup
	// decalre Results (nil value)
	var results []Result
	// Increment the waitgroup
	wg.Add(1)
	// a closure to control the queue
	go func() {
		// ensure the wg counter is decremented
		defer wg.Done()
		// outer loop using the queue size as the condition
		for i := 0; i < workerPoolSize; i++ {
			wg.Add(1)
			// closure wrapping around the feed from the deviceStream
			go func() {
				defer wg.Done()
				for d := range deviceStream {
					// augmenting here
					p := autochecks.Params{
						IP: d.Loopback,
					}
					_, err := ac.Run(p)
					res := Result{
						Device: d,
						Error:  err,
						Data:   ac,
					}
					queue <- res
				}
			}()
		}
	}()
	// Feeding the channel
	for _, d := range devices {
		deviceStream <- d
	}
	// close the input channel to signal we're done
	close(deviceStream)
	go func() {
		wg.Wait()
		//when all gorotuines are asleep,close channel
		close(queue)
	}()
	// Catch result and append to []Result
	for r := range queue {
		results = append(results, r)
	}
	// blocks until counter is back to zero
	wg.Wait()
	return results
}

// Subnets takes in a slice of string slices and an autocheck interface.
// It returns a slice of Result.
// Closures, coupled with channels, are used to pass data from one goroutine
// to another.
func Subnets(subnets [][]string, ac autochecks.Check) []Result {

	workerPoolSize := 64 // 128
	// define channels
	ipStream := make(chan string)
	queue := make(chan Result)

	var wg sync.WaitGroup

	var results []Result
	wg.Add(1)
	// a closure to control the queue
	go func() {
		defer wg.Done()
		for i := 0; i < workerPoolSize; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for ip := range ipStream {
					// augmenting here
					p := autochecks.Params{
						IP: ip,
					}
					_, err := ac.Run(p)
					d := models.Device{Loopback: ip}
					res := Result{
						Device: d,
						Error:  err,
						Data:   ac,
					}
					queue <- res
				}
			}()
		}
	}()
	// Feeding the channel
	for _, subnet := range subnets {
		for _, ip := range subnet {
			ipStream <- ip
		}
	}
	// close the input channel to signal we're done
	close(ipStream)
	go func() {
		wg.Wait()
		//when all gorotuines are asleep,close channel
		close(queue)
	}()
	// Catch result and append to []Result
	for r := range queue {
		results = append(results, r)
	}
	return results
}
