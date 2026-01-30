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
	fmt.Println("Context Basics")
	fmt.Println("==============")
	fmt.Println()

	// Example 1: context.Background() - The root context
	fmt.Println("1. context.Background() - Root context for main/init:")
	ctx := context.Background()
	fmt.Printf("   Type: %T\n", ctx)
	fmt.Printf("   Done channel: %v\n", ctx.Done())
	fmt.Printf("   Err: %v\n", ctx.Err())
	fmt.Println()

	// Example 2: context.TODO() - When you're unsure which context to use
	fmt.Println("2. context.TODO() - Placeholder context during refactoring:")
	todoCtx := context.TODO()
	fmt.Printf("   Type: %T\n", todoCtx)
	fmt.Printf("   Use TODO when refactoring or planning to add proper context\n")
	fmt.Println()

	// Example 3: Why context exists - Passing context through function calls
	fmt.Println("3. Passing context through function calls:")
	processRequest(ctx, "user-123")
	fmt.Println()

	// Example 4: Context in goroutines - Function signature pattern
	fmt.Println("4. Context as first parameter (Go convention):")
	showFunctionSignatures()
	fmt.Println()

	// Example 5: Checking context state
	fmt.Println("5. Context state inspection:")
	inspectContext(ctx)
	fmt.Println()

	// Example 6: Context is immutable
	fmt.Println("6. Context immutability:")
	parent := context.Background()
	child := context.WithValue(parent, "key", "value")
	fmt.Printf("   Parent has value: %v\n", parent.Value("key"))
	fmt.Printf("   Child has value: %v\n", child.Value("key"))
	fmt.Println("   Creating child context doesn't modify parent")
	fmt.Println()

	// Example 7: Real-world pattern - Simulating HTTP handler
	fmt.Println("7. Real-world pattern - HTTP handler simulation:")
	simulateHTTPHandler()
}

// processRequest demonstrates passing context through call chain
func processRequest(ctx context.Context, userID string) {
	fmt.Printf("   Processing request for user: %s\n", userID)

	// Context is passed to nested function calls
	fetchUserData(ctx, userID)
}

// fetchUserData shows context being passed deeper in the call chain
func fetchUserData(ctx context.Context, userID string) {
	fmt.Printf("   Fetching data for user: %s\n", userID)

	// In real code, this would check ctx.Done() for cancellation
	// or pass ctx to database/http calls
	validateData(ctx, userID)
}

// validateData shows the final level of the call chain
func validateData(ctx context.Context, userID string) {
	fmt.Printf("   Validating data for user: %s\n", userID)
	fmt.Printf("   Context successfully passed through entire chain\n")
}

// showFunctionSignatures demonstrates proper context function signatures
func showFunctionSignatures() {
	fmt.Println("   Correct function signatures:")
	fmt.Println("   func doWork(ctx context.Context, data string) error")
	fmt.Println("   func fetchUser(ctx context.Context, id int) (*User, error)")
	fmt.Println("   func query(ctx context.Context, sql string, args ...any) error")
	fmt.Println()
	fmt.Println("   Context is always:")
	fmt.Println("   - First parameter")
	fmt.Println("   - Named 'ctx' by convention")
	fmt.Println("   - Not stored in structs (with rare exceptions)")
}

// inspectContext shows how to check context properties
func inspectContext(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Printf("   Context has deadline: %v\n", deadline)
	} else {
		fmt.Printf("   Context has no deadline\n")
	}

	select {
	case <-ctx.Done():
		fmt.Printf("   Context is done: %v\n", ctx.Err())
	default:
		fmt.Printf("   Context is not done\n")
	}

	value := ctx.Value("example")
	fmt.Printf("   Context value for 'example': %v\n", value)
}

// simulateHTTPHandler shows a realistic use of context
func simulateHTTPHandler() {
	// In a real HTTP handler, you'd get context from request.Context()
	ctx := context.Background()

	// Simulate request processing
	requestID := "req-12345"
	fmt.Printf("   Handling request: %s\n", requestID)

	// Pass context to business logic
	err := processHTTPRequest(ctx, requestID)
	if err != nil {
		fmt.Printf("   Request failed: %v\n", err)
	} else {
		fmt.Printf("   Request completed successfully\n")
	}
}

// processHTTPRequest simulates business logic that accepts context
func processHTTPRequest(ctx context.Context, requestID string) error {
	// Check if context is already cancelled (client disconnected)
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// Simulate some work
	fmt.Printf("   Processing business logic for %s\n", requestID)
	time.Sleep(10 * time.Millisecond)

	// In real code, you'd pass ctx to database queries, HTTP calls, etc.
	return nil
}
