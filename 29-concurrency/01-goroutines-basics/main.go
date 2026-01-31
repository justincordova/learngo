package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Goroutines Basics")
	fmt.Println("=================")
	fmt.Println()

	// Example 1: Basic goroutine with named function
	fmt.Println("1. Basic goroutine with named function:")
	go printMessage("Hello from goroutine")
	time.Sleep(100 * time.Millisecond) // Give goroutine time to execute
	fmt.Println()

	// Example 2: Anonymous function goroutine
	fmt.Println("2. Anonymous function goroutine:")
	go func() {
		fmt.Println("Hello from anonymous goroutine")
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// Example 3: Multiple concurrent goroutines
	fmt.Println("3. Multiple concurrent goroutines:")
	for i := 1; i <= 5; i++ {
		go printNumber(i)
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println()

	// Example 4: Goroutine with closure
	fmt.Println("4. Goroutine with closure (capturing variables):")
	message := "Captured message"
	go func() {
		fmt.Printf("Closure says: %s\n", message)
	}()
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// Example 5: Using WaitGroup for synchronization
	fmt.Println("5. Using WaitGroup for proper synchronization:")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All workers completed")
	fmt.Println()

	// Example 6: Demonstrating race condition
	fmt.Println("6. Race condition demonstration:")
	fmt.Println("   (Run with 'go run -race main.go' to detect races)")
	counter := 0
	var wg2 sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			// RACE CONDITION: Multiple goroutines accessing counter
			temp := counter
			time.Sleep(time.Millisecond) // Simulate work
			counter = temp + 1
		}()
	}

	wg2.Wait()
	fmt.Printf("Counter value (with race): %d (expected 5, but may vary)\n", counter)
	fmt.Println()

	// Example 7: Fixed version with mutex
	fmt.Println("7. Fixed version with mutex:")
	safeCounter := 0
	var mu sync.Mutex
	var wg3 sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			mu.Lock()
			temp := safeCounter
			time.Sleep(time.Millisecond)
			safeCounter = temp + 1
			mu.Unlock()
		}()
	}

	wg3.Wait()
	fmt.Printf("Counter value (thread-safe): %d\n", safeCounter)
}

// printMessage prints a message from a goroutine
func printMessage(msg string) {
	fmt.Println(msg)
}

// printNumber prints a number with a small delay
func printNumber(n int) {
	time.Sleep(time.Duration(n*10) * time.Millisecond)
	fmt.Printf("Number: %d\n", n)
}
