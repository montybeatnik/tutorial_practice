package main

import (
	"fmt"
	"sync"
	"time"
)

type A struct {
	id int
}

func main() {
	start := time.Now()

	workerPoolSize := 100

	channel := make(chan A, 100)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < workerPoolSize; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				for a := range channel {
					process(a)
				}
			}()
		}

	}()

	// Feeding the channel
	for i := 0; i < 100000; i++ {
		channel <- A{id: i}
	}
	close(channel)

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}

func process(a A) {
	fmt.Printf("Start processing %v\n", a)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Finish processing %v\n", a)
}
