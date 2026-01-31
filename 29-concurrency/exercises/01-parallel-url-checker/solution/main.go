package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Result holds the check result for a URL
type Result struct {
	URL    string
	Status string
	Error  error
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://invalid-url-xyz-123.com",
		"https://www.golang.org",
		"https://www.example.com",
	}

	fmt.Printf("Checking %d URLs concurrently...\n\n", len(urls))

	start := time.Now()

	// Create channel to receive results
	results := make(chan Result, len(urls))

	// WaitGroup to track goroutines
	var wg sync.WaitGroup

	// Launch a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			results <- checkURL(u)
		}(url)
	}

	// Close results channel when all goroutines complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results as they come in
	for result := range results {
		if result.Status == "reachable" {
			fmt.Printf("%s: %s\n", result.URL, result.Status)
		} else {
			fmt.Printf("%s: %s (%v)\n", result.URL, result.Status, result.Error)
		}
	}

	fmt.Printf("\nCompleted in %.3fs\n", time.Since(start).Seconds())
}

// checkURL makes an HTTP GET request to the URL and returns the result
func checkURL(url string) Result {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Result{
			URL:    url,
			Status: "unreachable",
			Error:  err,
		}
	}

	// Make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Result{
			URL:    url,
			Status: "unreachable",
			Error:  err,
		}
	}
	defer resp.Body.Close()

	return Result{
		URL:    url,
		Status: "reachable",
		Error:  nil,
	}
}
