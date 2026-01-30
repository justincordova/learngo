// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Channels")
	fmt.Println("========")
	fmt.Println()

	// Example 1: Basic unbuffered channel
	fmt.Println("1. Basic unbuffered channel:")
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine" // Send to channel
	}()

	msg := <-ch // Receive from channel
	fmt.Printf("Received: %s\n\n", msg)

	// Example 2: Unbuffered channel blocks
	fmt.Println("2. Unbuffered channels block until both sender and receiver are ready:")
	ch2 := make(chan int)

	go func() {
		fmt.Println("   Goroutine: Preparing to send...")
		time.Sleep(100 * time.Millisecond)
		ch2 <- 42
		fmt.Println("   Goroutine: Sent!")
	}()

	fmt.Println("   Main: Waiting to receive...")
	value := <-ch2
	fmt.Printf("   Main: Received %d\n\n", value)

	// Example 3: Buffered channel
	fmt.Println("3. Buffered channel (capacity 3):")
	buffered := make(chan int, 3)

	// Can send up to 3 values without blocking
	buffered <- 1
	buffered <- 2
	buffered <- 3
	fmt.Println("   Sent 3 values without blocking")

	// Receive the values
	fmt.Printf("   Received: %d, %d, %d\n\n", <-buffered, <-buffered, <-buffered)

	// Example 4: Closing channels
	fmt.Println("4. Closing channels:")
	ch3 := make(chan int, 3)

	// Send some values
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	close(ch3) // Close the channel

	// Receive until channel is closed
	for value := range ch3 {
		fmt.Printf("   Received: %d\n", value)
	}
	fmt.Println("   Channel closed and drained")
	fmt.Println()

	// Example 5: Checking if channel is closed
	fmt.Println("5. Checking if channel is closed:")
	ch4 := make(chan string, 1)
	ch4 <- "last message"
	close(ch4)

	msg1, ok := <-ch4
	fmt.Printf("   First receive: %q, ok=%v\n", msg1, ok)

	msg2, ok := <-ch4
	fmt.Printf("   Second receive: %q, ok=%v (channel closed)\n\n", msg2, ok)

	// Example 6: Pipeline pattern
	fmt.Println("6. Pipeline pattern (generator -> processor -> consumer):")
	numbers := generate(1, 2, 3, 4, 5)
	squares := square(numbers)

	fmt.Print("   Squares: ")
	for result := range squares {
		fmt.Printf("%d ", result)
	}
	fmt.Println()
	fmt.Println()

	// Example 7: Fan-out/Fan-in pattern basics
	fmt.Println("7. Fan-out/Fan-in pattern:")
	input := generate(1, 2, 3, 4, 5, 6)

	// Fan-out: distribute work to multiple workers
	worker1 := square(input)
	worker2 := square(input)

	// Fan-in: merge results from multiple workers
	results := merge(worker1, worker2)

	fmt.Print("   Results: ")
	for result := range results {
		fmt.Printf("%d ", result)
	}
	fmt.Println()
	fmt.Println()

	// Example 8: Channel direction (send-only, receive-only)
	fmt.Println("8. Channel direction:")
	ch5 := make(chan int, 1)

	go sendOnly(ch5)  // Pass as send-only
	receiveOnly(ch5)  // Pass as receive-only
	fmt.Println()

	// Example 9: Worker pool pattern
	fmt.Println("9. Worker pool pattern:")
	jobs := make(chan int, 5)
	results2 := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results2)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		<-results2
	}
}

// generate creates a channel and sends values to it
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square receives numbers and sends their squares
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// merge combines multiple channels into one
func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)

	for _, ch := range channels {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
		}(ch)
	}

	// Close output channel when all inputs are done
	go func() {
		time.Sleep(100 * time.Millisecond) // Simple synchronization for demo
		close(out)
	}()

	return out
}

// sendOnly demonstrates send-only channel parameter
func sendOnly(ch chan<- int) {
	ch <- 99
	fmt.Println("   Sent to send-only channel")
}

// receiveOnly demonstrates receive-only channel parameter
func receiveOnly(ch <-chan int) {
	value := <-ch
	fmt.Printf("   Received from receive-only channel: %d\n", value)
}

// worker processes jobs and sends results
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("   Worker %d processing job %d\n", id, j)
		time.Sleep(50 * time.Millisecond)
		results <- j * 2
	}
}
