package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Context Cancellation")
	fmt.Println("====================")
	fmt.Println()

	// Example 1: Basic cancellation with WithCancel
	fmt.Println("1. Basic cancellation with WithCancel:")
	example1BasicCancellation()
	fmt.Println()

	// Example 2: Cancelling multiple goroutines
	fmt.Println("2. Cancelling multiple goroutines:")
	example2MultipleGoroutines()
	fmt.Println()

	// Example 3: Graceful worker shutdown
	fmt.Println("3. Graceful worker shutdown:")
	example3WorkerShutdown()
	fmt.Println()

	// Example 4: Cancellation propagation (parent cancels children)
	fmt.Println("4. Cancellation propagation:")
	example4Propagation()
	fmt.Println()

	// Example 5: Cleanup with defer and cancel
	fmt.Println("5. Cleanup with defer and cancel:")
	example5Cleanup()
	fmt.Println()

	// Example 6: Checking cancellation reason
	fmt.Println("6. Checking cancellation reason:")
	example6CancellationReason()
	fmt.Println()
}

// example1BasicCancellation shows basic context cancellation
func example1BasicCancellation() {
	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Start a goroutine that respects cancellation
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("   Goroutine received cancellation signal")
				fmt.Println("   Reason:", ctx.Err())
				return
			default:
				fmt.Println("   Working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Let it work for a bit
	time.Sleep(350 * time.Millisecond)

	// Cancel the context
	fmt.Println("   Calling cancel()...")
	cancel()

	// Give goroutine time to clean up
	time.Sleep(100 * time.Millisecond)
}

// example2MultipleGoroutines shows cancelling multiple goroutines at once
func example2MultipleGoroutines() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Start multiple workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// Let workers run for a bit
	time.Sleep(300 * time.Millisecond)

	// Cancel all workers at once
	fmt.Println("   Cancelling all workers...")
	cancel()

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("   All workers stopped")
}

// worker is a simple worker that respects context cancellation
func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("   Worker %d shutting down\n", id)
			return
		default:
			fmt.Printf("   Worker %d processing...\n", id)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// example3WorkerShutdown demonstrates graceful worker pool shutdown
func example3WorkerShutdown() {
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	// Start worker pool
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go jobWorker(ctx, i, jobs, &wg)
	}

	// Send some jobs
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
			time.Sleep(50 * time.Millisecond)
		}
		close(jobs)
	}()

	// Let workers process some jobs
	time.Sleep(200 * time.Millisecond)

	// Initiate shutdown
	fmt.Println("   Initiating graceful shutdown...")
	cancel()

	// Wait for workers to finish
	wg.Wait()
	fmt.Println("   Worker pool shutdown complete")
}

// jobWorker processes jobs until context is cancelled
func jobWorker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("   Worker %d: received shutdown signal\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("   Worker %d: job channel closed\n", id)
				return
			}
			fmt.Printf("   Worker %d: processing job %d\n", id, job)
			time.Sleep(75 * time.Millisecond)
		}
	}
}

// example4Propagation shows how cancellation propagates to children
func example4Propagation() {
	// Create parent context
	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	// Create child contexts
	childCtx1, cancel1 := context.WithCancel(parentCtx)
	defer cancel1()

	childCtx2, cancel2 := context.WithCancel(parentCtx)
	defer cancel2()

	// Start goroutines with child contexts
	go monitorContext(childCtx1, "Child 1")
	go monitorContext(childCtx2, "Child 2")

	time.Sleep(100 * time.Millisecond)

	// Cancel parent - this cancels all children
	fmt.Println("   Cancelling parent context...")
	parentCancel()

	time.Sleep(100 * time.Millisecond)
}

// monitorContext monitors a context and reports when it's cancelled
func monitorContext(ctx context.Context, name string) {
	<-ctx.Done()
	fmt.Printf("   %s cancelled: %v\n", name, ctx.Err())
}

// example5Cleanup demonstrates proper cleanup with defer
func example5Cleanup() {
	ctx, cancel := context.WithCancel(context.Background())
	// IMPORTANT: Always defer cancel() to prevent resource leaks
	defer cancel()

	done := make(chan bool)

	go func() {
		<-ctx.Done()
		fmt.Println("   Cleanup: goroutine received cancellation")
		// Perform cleanup operations here
		fmt.Println("   Cleanup: releasing resources...")
		time.Sleep(50 * time.Millisecond)
		fmt.Println("   Cleanup: done")
		done <- true
	}()

	// Do some work
	time.Sleep(100 * time.Millisecond)

	// Trigger cancellation (defer will also call it, which is safe)
	fmt.Println("   Cleanup: triggering shutdown...")
	cancel()

	// Wait for cleanup to complete
	<-done
}

// example6CancellationReason shows how to check why context was cancelled
func example6CancellationReason() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-ctx.Done()

		// Check the reason for cancellation
		err := ctx.Err()
		switch err {
		case context.Canceled:
			fmt.Println("   Context was explicitly cancelled")
		case context.DeadlineExceeded:
			fmt.Println("   Context deadline exceeded")
		default:
			fmt.Printf("   Unknown error: %v\n", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	// Explicitly cancel
	cancel()

	time.Sleep(100 * time.Millisecond)
}
