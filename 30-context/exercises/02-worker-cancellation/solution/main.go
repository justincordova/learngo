package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Channel for jobs
	jobs := make(chan int, 10)

	// WaitGroup to track workers
	var wg sync.WaitGroup

	// Counter for completed jobs
	var completed atomic.Int32

	// Start workers
	numWorkers := 3
	fmt.Printf("Starting %d workers...\n", numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg, &completed)
	}

	// Send jobs
	numJobs := 10
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Let workers process for 1 second
	time.Sleep(1 * time.Second)

	// Cancel the context
	fmt.Println("\nCancelling context...")
	cancel()

	// Wait for all workers to finish
	wg.Wait()

	fmt.Printf("\nAll workers stopped. Jobs processed: %d/%d\n", completed.Load(), numJobs)
}

// worker processes jobs until context is cancelled
func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup, completed *atomic.Int32) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Shutting down (context cancelled)\n", id)
			return

		case job, ok := <-jobs:
			if !ok {
				// Jobs channel closed and empty
				fmt.Printf("Worker %d: No more jobs\n", id)
				return
			}

			// Process the job
			if processJob(ctx, id, job) {
				completed.Add(1)
			}
		}
	}
}

// processJob simulates processing a job and checks for cancellation
func processJob(ctx context.Context, workerID, jobID int) bool {
	fmt.Printf("Worker %d: Processing job %d\n", workerID, jobID)

	// Simulate work with random duration
	workDuration := time.Duration(100+rand.Intn(400)) * time.Millisecond
	halfDuration := workDuration / 2

	// Do half the work
	time.Sleep(halfDuration)

	// Check if context was cancelled mid-job
	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d: Job %d cancelled mid-processing\n", workerID, jobID)
		return false
	default:
		// Continue working
	}

	// Do the remaining work
	time.Sleep(halfDuration)

	// Final check before marking complete
	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d: Job %d cancelled before completion\n", workerID, jobID)
		return false
	default:
		fmt.Printf("Worker %d: Completed job %d\n", workerID, jobID)
		return true
	}
}
