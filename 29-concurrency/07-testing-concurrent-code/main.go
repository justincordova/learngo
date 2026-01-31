package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Testing Concurrent Code (Go 1.25+)")
	fmt.Println("===================================")
	fmt.Println()
	fmt.Println("This example demonstrates patterns for testing concurrent code.")
	fmt.Println("Run 'go test -v' to see the synctest package in action!")
	fmt.Println()

	// Example 1: RetryWithBackoff - retries operation with exponential backoff
	fmt.Println("Example 1: Retry with backoff")
	attempts := 0
	err := RetryWithBackoff(context.Background(), 3, 50*time.Millisecond, func() error {
		attempts++
		fmt.Printf("  Attempt %d\n", attempts)
		if attempts < 3 {
			return fmt.Errorf("temporary error")
		}
		return nil
	})
	if err != nil {
		fmt.Printf("  Failed after retries: %v\n", err)
	} else {
		fmt.Printf("  Succeeded after %d attempts\n", attempts)
	}
	fmt.Println()

	// Example 2: Debouncer - debounces rapid events
	fmt.Println("Example 2: Debouncer")
	var result string
	var mu sync.Mutex

	debouncer := NewDebouncer(100 * time.Millisecond)
	defer debouncer.Stop()

	// Simulate rapid events
	for i := 1; i <= 5; i++ {
		i := i
		debouncer.Debounce(func() {
			mu.Lock()
			result = fmt.Sprintf("Event %d", i)
			fmt.Printf("  Executed: %s\n", result)
			mu.Unlock()
		})
		time.Sleep(30 * time.Millisecond)
	}

	// Wait for last event to execute
	time.Sleep(150 * time.Millisecond)
	mu.Lock()
	fmt.Printf("  Final result: %s\n", result)
	mu.Unlock()
	fmt.Println()

	// Example 3: Timeout - operation with timeout
	fmt.Println("Example 3: Operation with timeout")

	// Fast operation - should succeed
	err = WithTimeout(100*time.Millisecond, func() error {
		time.Sleep(50 * time.Millisecond)
		return nil
	})
	if err != nil {
		fmt.Printf("  Fast operation failed: %v\n", err)
	} else {
		fmt.Println("  Fast operation succeeded")
	}

	// Slow operation - should timeout
	err = WithTimeout(100*time.Millisecond, func() error {
		time.Sleep(200 * time.Millisecond)
		return nil
	})
	if err != nil {
		fmt.Printf("  Slow operation failed: %v\n", err)
	} else {
		fmt.Println("  Slow operation succeeded")
	}
}

// RetryWithBackoff retries an operation with exponential backoff
func RetryWithBackoff(ctx context.Context, maxAttempts int, initialDelay time.Duration, operation func() error) error {
	delay := initialDelay

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err := operation()
		if err == nil {
			return nil
		}

		if attempt == maxAttempts {
			return fmt.Errorf("failed after %d attempts: %w", maxAttempts, err)
		}

		// Wait before retry with exponential backoff
		select {
		case <-time.After(delay):
			delay *= 2
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}

// Debouncer delays execution until events stop arriving
type Debouncer struct {
	delay time.Duration
	timer *time.Timer
	mu    sync.Mutex
	stop  chan struct{}
}

func NewDebouncer(delay time.Duration) *Debouncer {
	return &Debouncer{
		delay: delay,
		stop:  make(chan struct{}),
	}
}

func (d *Debouncer) Debounce(f func()) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}

	d.timer = time.AfterFunc(d.delay, f)
}

func (d *Debouncer) Stop() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}
	close(d.stop)
}

// WithTimeout executes an operation with a timeout
func WithTimeout(timeout time.Duration, operation func() error) error {
	done := make(chan error, 1)

	go func() {
		done <- operation()
	}()

	select {
	case err := <-done:
		return err
	case <-time.After(timeout):
		return fmt.Errorf("operation timed out after %v", timeout)
	}
}

// BatchProcessor processes items in batches with delays
func BatchProcessor(items []string, batchSize int, delay time.Duration) []string {
	var results []string
	var mu sync.Mutex

	for i := 0; i < len(items); i += batchSize {
		end := i + batchSize
		if end > len(items) {
			end = len(items)
		}

		batch := items[i:end]

		// Process batch
		for _, item := range batch {
			processed := fmt.Sprintf("processed-%s", item)
			mu.Lock()
			results = append(results, processed)
			mu.Unlock()
		}

		// Delay between batches (except after last batch)
		if end < len(items) {
			time.Sleep(delay)
		}
	}

	return results
}
