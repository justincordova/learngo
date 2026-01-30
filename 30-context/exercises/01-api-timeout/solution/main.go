// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// List of URLs with different delays
	urls := []string{
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/3",
		"https://httpbin.org/delay/5",
	}

	timeout := 2 * time.Second

	for _, url := range urls {
		fmt.Printf("Fetching %s with %v timeout...\n", url, timeout)

		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		// Fetch the URL
		statusCode, err := fetchURL(ctx, url)

		// Always call cancel to release resources
		cancel()

		if err != nil {
			// Check if it was a timeout
			if errors.Is(err, context.DeadlineExceeded) {
				fmt.Printf("Error: request timed out: %v\n\n", err)
			} else {
				fmt.Printf("Error: %v\n\n", err)
			}
			continue
		}

		fmt.Printf("Success: Status %d %s\n\n", statusCode, http.StatusText(statusCode))
	}
}

// fetchURL makes an HTTP GET request with the provided context
func fetchURL(ctx context.Context, url string) (int, error) {
	// Create a new request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
