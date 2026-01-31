package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create list of requests
	requests := make([]string, 10)
	for i := 0; i < 10; i++ {
		requests[i] = fmt.Sprintf("Request %d", i+1)
	}

	fmt.Println("Starting rate-limited request processor...")
	fmt.Println("Rate limit: 2 requests per second")
	fmt.Println()

	start := time.Now()

	// Create a ticker that ticks every 500ms (2 per second)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Channel to send requests
	requestCh := make(chan string)

	// WaitGroup to wait for all requests to complete
	var wg sync.WaitGroup

	// Start processor goroutine
	go func() {
		for req := range requestCh {
			wg.Add(1)
			go func(r string, startTime time.Time) {
				defer wg.Done()
				processRequest(r, startTime)
			}(req, start)
		}
	}()

	// Send requests with rate limiting
	i := 0
	// Allow burst of 2 initial requests
	for j := 0; j < 2 && i < len(requests); j++ {
		requestCh <- requests[i]
		i++
	}

	// Send remaining requests at the rate limit
	for i < len(requests) {
		<-ticker.C // Wait for next tick
		// Send up to 2 requests per tick
		for j := 0; j < 2 && i < len(requests); j++ {
			requestCh <- requests[i]
			i++
		}
	}

	// Close the channel and wait for all requests to complete
	close(requestCh)
	wg.Wait()

	fmt.Printf("\nAll requests completed in %.3fs\n", time.Since(start).Seconds())
}

// processRequest simulates processing an API request
func processRequest(id string, startTime time.Time) {
	elapsed := time.Since(startTime).Seconds()
	fmt.Printf("[%.3fs] Processing %s\n", elapsed, id)

	// Simulate processing time
	time.Sleep(100 * time.Millisecond)

	elapsed = time.Since(startTime).Seconds()
	fmt.Printf("[%.3fs] Completed %s\n", elapsed, id)
}
