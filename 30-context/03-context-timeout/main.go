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
	"time"
)

func main() {
	fmt.Println("Context Timeout and Deadline")
	fmt.Println("=============================")
	fmt.Println()

	// Example 1: Basic timeout with WithTimeout
	fmt.Println("1. Basic timeout with WithTimeout:")
	example1BasicTimeout()
	fmt.Println()

	// Example 2: Basic deadline with WithDeadline
	fmt.Println("2. Basic deadline with WithDeadline:")
	example2BasicDeadline()
	fmt.Println()

	// Example 3: Timeout vs Deadline - Understanding the difference
	fmt.Println("3. Timeout vs Deadline - The difference:")
	example3TimeoutVsDeadline()
	fmt.Println()

	// Example 4: Simulating API call with timeout
	fmt.Println("4. API call with timeout:")
	example4APICallTimeout()
	fmt.Println()

	// Example 5: Database query with timeout
	fmt.Println("5. Database query with timeout:")
	example5DatabaseTimeout()
	fmt.Println()

	// Example 6: Successful operation within timeout
	fmt.Println("6. Successful operation within timeout:")
	example6SuccessWithinTimeout()
	fmt.Println()

	// Example 7: Checking remaining time before deadline
	fmt.Println("7. Checking remaining time before deadline:")
	example7CheckDeadline()
	fmt.Println()

	// Example 8: Nested timeouts (child inherits parent deadline)
	fmt.Println("8. Nested timeouts:")
	example8NestedTimeouts()
	fmt.Println()
}

// example1BasicTimeout demonstrates context.WithTimeout
func example1BasicTimeout() {
	// Create a context that times out after 200ms
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel() // Always defer cancel to release resources

	// Start a long-running operation
	go func() {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("   Operation completed")
		case <-ctx.Done():
			fmt.Println("   Operation cancelled:", ctx.Err())
		}
	}()

	// Wait for context to be done
	<-ctx.Done()
	fmt.Println("   Context timed out after 200ms")
}

// example2BasicDeadline demonstrates context.WithDeadline
func example2BasicDeadline() {
	// Create a context that expires at a specific time
	deadline := time.Now().Add(200 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// Check the deadline
	dl, ok := ctx.Deadline()
	if ok {
		fmt.Printf("   Deadline set for: %v\n", dl.Format("15:04:05.000"))
		fmt.Printf("   Time until deadline: %v\n", time.Until(dl))
	}

	// Wait for deadline
	<-ctx.Done()
	fmt.Println("   Context reached deadline:", ctx.Err())
}

// example3TimeoutVsDeadline explains the difference
func example3TimeoutVsDeadline() {
	now := time.Now()

	// WithTimeout: relative duration from now
	fmt.Println("   WithTimeout (relative):")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel1()
	fmt.Printf("   - Expires in: 100ms from now\n")
	fmt.Printf("   - Created at: %v\n", now.Format("15:04:05.000"))

	// WithDeadline: absolute point in time
	fmt.Println()
	fmt.Println("   WithDeadline (absolute):")
	deadline := time.Now().Add(100 * time.Millisecond)
	ctx2, cancel2 := context.WithDeadline(context.Background(), deadline)
	defer cancel2()
	fmt.Printf("   - Expires at: %v\n", deadline.Format("15:04:05.000"))

	// Both achieve the same result, just different ways to specify when
	<-ctx1.Done()
	<-ctx2.Done()
	fmt.Println()
	fmt.Println("   Note: WithTimeout is implemented using WithDeadline internally")
	fmt.Println("   Use WithTimeout for relative durations (e.g., '5 seconds from now')")
	fmt.Println("   Use WithDeadline for absolute times (e.g., 'at 3:00 PM')")
}

// example4APICallTimeout simulates an API call with timeout
func example4APICallTimeout() {
	// Simulate a slow API call
	slowAPI := func(ctx context.Context) error {
		select {
		case <-time.After(500 * time.Millisecond):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	// Call API with 200ms timeout
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	fmt.Println("   Calling slow API (500ms) with 200ms timeout...")
	err := slowAPI(ctx)
	if err != nil {
		fmt.Printf("   API call failed: %v\n", err)
		if err == context.DeadlineExceeded {
			fmt.Println("   Reason: Request took too long")
		}
	} else {
		fmt.Println("   API call succeeded")
	}
}

// example5DatabaseTimeout simulates a database query with timeout
func example5DatabaseTimeout() {
	// Simulate database query
	queryDatabase := func(ctx context.Context, query string) (string, error) {
		// Check if context already expired before starting
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		// Simulate query execution
		resultChan := make(chan string, 1)
		go func() {
			time.Sleep(150 * time.Millisecond)
			resultChan <- "query result"
		}()

		select {
		case result := <-resultChan:
			return result, nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}

	// Query with 100ms timeout (will fail)
	fmt.Println("   Query 1: 150ms query with 100ms timeout")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel1()

	result, err := queryDatabase(ctx1, "SELECT * FROM users")
	if err != nil {
		fmt.Printf("   Failed: %v\n", err)
	} else {
		fmt.Printf("   Success: %s\n", result)
	}

	fmt.Println()

	// Query with 200ms timeout (will succeed)
	fmt.Println("   Query 2: 150ms query with 200ms timeout")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel2()

	result, err = queryDatabase(ctx2, "SELECT * FROM users")
	if err != nil {
		fmt.Printf("   Failed: %v\n", err)
	} else {
		fmt.Printf("   Success: %s\n", result)
	}
}

// example6SuccessWithinTimeout shows successful completion before timeout
func example6SuccessWithinTimeout() {
	// Fast operation that completes before timeout
	fastOperation := func(ctx context.Context) error {
		select {
		case <-time.After(50 * time.Millisecond):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	fmt.Println("   Running fast operation (50ms) with 200ms timeout...")
	err := fastOperation(ctx)
	if err != nil {
		fmt.Printf("   Failed: %v\n", err)
	} else {
		fmt.Println("   Completed successfully within timeout")
		// Calling cancel() early is good practice - releases resources immediately
		cancel()
		fmt.Println("   Called cancel() to release resources early")
	}
}

// example7CheckDeadline shows how to check remaining time
func example7CheckDeadline() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	deadline, ok := ctx.Deadline()
	if !ok {
		fmt.Println("   No deadline set")
		return
	}

	fmt.Printf("   Total timeout: 300ms\n")

	for i := 0; i < 3; i++ {
		remaining := time.Until(deadline)
		fmt.Printf("   Check %d - Time remaining: %v\n", i+1, remaining.Round(time.Millisecond))

		if remaining <= 0 {
			fmt.Println("   Deadline reached!")
			break
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// example8NestedTimeouts shows how child contexts inherit parent deadlines
func example8NestedTimeouts() {
	// Parent context with 500ms timeout
	parentCtx, parentCancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer parentCancel()

	fmt.Println("   Parent timeout: 500ms")

	// Child context with 1s timeout (but will inherit parent's shorter deadline)
	childCtx, childCancel := context.WithTimeout(parentCtx, 1*time.Second)
	defer childCancel()

	// Check child's deadline
	deadline, ok := childCtx.Deadline()
	if ok {
		fmt.Printf("   Child timeout requested: 1000ms\n")
		fmt.Printf("   Child actual deadline: %v (inherited from parent)\n",
			time.Until(deadline).Round(time.Millisecond))
	}

	// Wait for child context to finish
	<-childCtx.Done()
	fmt.Println("   Child context finished:", childCtx.Err())
	fmt.Println()
	fmt.Println("   Note: Child inherits parent's deadline if parent's is sooner")
}
