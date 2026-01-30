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
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	fmt.Println("Worker Pool Pattern")
	fmt.Println("===================")
	fmt.Println()

	// Example 1: Basic Worker Pool
	fmt.Println("Example 1: Basic worker pool with number processing")
	basicWorkerPool()
	fmt.Println()

	// Example 2: Worker Pool with Error Handling
	fmt.Println("Example 2: Worker pool with error handling")
	workerPoolWithErrors()
	fmt.Println()

	// Example 3: URL Fetcher Worker Pool (practical example)
	fmt.Println("Example 3: URL fetcher worker pool")
	urlFetcherPool()
	fmt.Println()

	// Example 4: Image Processor Worker Pool
	fmt.Println("Example 4: Image processor worker pool (simulated)")
	imageProcessorPool()
	fmt.Println()

	// Example 5: Worker Pool with Timeout and Cancellation
	fmt.Println("Example 5: Worker pool with timeout and cancellation")
	workerPoolWithTimeout()
}

// Example 1: Basic Worker Pool
func basicWorkerPool() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("  Worker %d processing job %d\n", id, job)
				time.Sleep(100 * time.Millisecond) // Simulate work
				results <- job * 2                 // Process job
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect results
	fmt.Println("  Results:")
	for result := range results {
		fmt.Printf("  %d ", result)
	}
	fmt.Println()
}

// Example 2: Worker Pool with Error Handling
type Job struct {
	ID    int
	Value int
}

type Result struct {
	Job   Job
	Value int
	Error error
}

func workerPoolWithErrors() {
	const numWorkers = 3
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				result := Result{Job: job}

				// Simulate error for negative values
				if job.Value < 0 {
					result.Error = fmt.Errorf("worker %d: cannot process negative value %d", id, job.Value)
				} else {
					result.Value = job.Value * 2
					fmt.Printf("  Worker %d processed job %d successfully\n", id, job.ID)
				}

				results <- result
			}
		}(w)
	}

	// Send jobs (including some that will error)
	jobData := []int{1, 2, -3, 4, -5, 6, 7, 8}
	for i, val := range jobData {
		jobs <- Job{ID: i + 1, Value: val}
	}
	close(jobs)

	// Collect results in a separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	successCount := 0
	errorCount := 0
	for result := range results {
		if result.Error != nil {
			fmt.Printf("  Error: %v\n", result.Error)
			errorCount++
		} else {
			successCount++
		}
	}
	fmt.Printf("  Completed: %d successful, %d errors\n", successCount, errorCount)
}

// Example 3: URL Fetcher Worker Pool
type URLJob struct {
	URL string
}

type URLResult struct {
	URL    string
	Status int
	Size   int
	Error  error
}

func urlFetcherPool() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://github.com",
		"https://www.google.com",
		"https://invalid-url-that-doesnt-exist.xyz",
	}

	const numWorkers = 3
	jobs := make(chan URLJob, len(urls))
	results := make(chan URLResult, len(urls))

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			client := &http.Client{
				Timeout: 5 * time.Second,
			}

			for job := range jobs {
				result := URLResult{URL: job.URL}
				fmt.Printf("  Worker %d fetching %s\n", id, job.URL)

				resp, err := client.Get(job.URL)
				if err != nil {
					result.Error = err
					results <- result
					continue
				}

				result.Status = resp.StatusCode
				body, err := io.ReadAll(resp.Body)
				resp.Body.Close()

				if err != nil {
					result.Error = err
				} else {
					result.Size = len(body)
				}

				results <- result
			}
		}(w)
	}

	// Send jobs
	for _, url := range urls {
		jobs <- URLJob{URL: url}
	}
	close(jobs)

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	for result := range results {
		if result.Error != nil {
			fmt.Printf("  %s: ERROR - %v\n", result.URL, result.Error)
		} else {
			fmt.Printf("  %s: %d bytes (status: %d)\n", result.URL, result.Size, result.Status)
		}
	}
}

// Example 4: Image Processor Worker Pool
type ImageJob struct {
	ID       int
	Filename string
}

type ImageResult struct {
	Job       ImageJob
	Processed bool
	Duration  time.Duration
	Error     error
}

func imageProcessorPool() {
	images := []string{
		"image1.jpg", "image2.jpg", "image3.jpg",
		"image4.jpg", "image5.jpg", "image6.jpg",
		"image7.jpg", "image8.jpg",
	}

	const numWorkers = 4
	jobs := make(chan ImageJob, len(images))
	results := make(chan ImageResult, len(images))

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				start := time.Now()
				result := ImageResult{Job: job}

				fmt.Printf("  Worker %d processing %s\n", id, job.Filename)

				// Simulate image processing (resize, compress, etc.)
				processingTime := time.Duration(50+job.ID*10) * time.Millisecond
				time.Sleep(processingTime)

				result.Processed = true
				result.Duration = time.Since(start)
				results <- result
			}
		}(w)
	}

	// Send jobs
	for i, img := range images {
		jobs <- ImageJob{ID: i + 1, Filename: img}
	}
	close(jobs)

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	totalDuration := time.Duration(0)
	processedCount := 0

	for result := range results {
		if result.Processed {
			processedCount++
			totalDuration += result.Duration
			fmt.Printf("  Processed %s in %v\n", result.Job.Filename, result.Duration)
		}
	}

	fmt.Printf("  Processed %d images in total time: %v\n", processedCount, totalDuration)
	fmt.Printf("  Average processing time: %v\n", totalDuration/time.Duration(processedCount))
}

// Example 5: Worker Pool with Timeout and Cancellation
func workerPoolWithTimeout() {
	const numWorkers = 3
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("  Worker %d shutting down due to timeout\n", id)
					return
				case job, ok := <-jobs:
					if !ok {
						return
					}
					// Simulate work
					select {
					case <-ctx.Done():
						fmt.Printf("  Worker %d cancelled job %d\n", id, job)
						return
					case <-time.After(100 * time.Millisecond):
						results <- job * 2
					}
				}
			}
		}(w)
	}

	// Send jobs (will be interrupted by timeout)
	go func() {
		for j := 1; j <= 20; j++ {
			select {
			case <-ctx.Done():
				fmt.Println("  Job sender stopped due to timeout")
				close(jobs)
				return
			case jobs <- j:
				fmt.Printf("  Sent job %d\n", j)
			}
		}
		close(jobs)
	}()

	// Collect results with timeout
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results until channel closes or context cancelled
	processedCount := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("  Result collector stopped: %v\n", ctx.Err())
			return
		case result, ok := <-results:
			if !ok {
				fmt.Printf("  Processed %d jobs before shutdown\n", processedCount)
				return
			}
			processedCount++
			fmt.Printf("  Result: %d\n", result)
		}
	}
}
