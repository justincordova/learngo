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
	fmt.Println("Channel Select Statement")
	fmt.Println("========================")
	fmt.Println()

	// Example 1: Basic select with multiple channels
	fmt.Println("1. Basic select with two channels:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	// Select waits on multiple channel operations
	select {
	case msg1 := <-ch1:
		fmt.Printf("   Received: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("   Received: %s\n", msg2)
	}
	fmt.Println()

	// Example 2: Select with timeout
	fmt.Println("2. Select with timeout pattern:")
	ch3 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch3 <- "slow response"
	}()

	select {
	case msg := <-ch3:
		fmt.Printf("   Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("   Timeout: operation took too long")
	}
	fmt.Println()

	// Example 3: Select with default case (non-blocking)
	fmt.Println("3. Non-blocking select with default case:")
	ch4 := make(chan string)

	select {
	case msg := <-ch4:
		fmt.Printf("   Received: %s\n", msg)
	default:
		fmt.Println("   No message available, continuing without blocking")
	}
	fmt.Println()

	// Example 4: Receiving from multiple channels
	fmt.Println("4. Receiving from multiple channels until done:")
	done := make(chan bool)
	messages := make(chan string)

	go func() {
		for i := 1; i <= 3; i++ {
			messages <- fmt.Sprintf("message %d", i)
			time.Sleep(50 * time.Millisecond)
		}
		done <- true
	}()

	for {
		select {
		case msg := <-messages:
			fmt.Printf("   Received: %s\n", msg)
		case <-done:
			fmt.Println("   All messages received")
			goto next
		}
	}
next:
	fmt.Println()

	// Example 5: Select with send operations
	fmt.Println("5. Select with send operations:")
	ch5 := make(chan string, 1)
	ch6 := make(chan string, 1)

	select {
	case ch5 <- "sent to ch5":
		fmt.Println("   Sent to channel 5")
	case ch6 <- "sent to ch6":
		fmt.Println("   Sent to channel 6")
	}
	fmt.Println()

	// Example 6: Multiple select in a loop
	fmt.Println("6. Ticker and timeout with select:")
	ticker := time.NewTicker(50 * time.Millisecond)
	timeout := time.After(250 * time.Millisecond)

	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("   Tick at %v\n", t.Format("15:04:05.000"))
		case <-timeout:
			fmt.Println("   Timeout reached, stopping ticker")
			ticker.Stop()
			goto done6
		}
	}
done6:
	fmt.Println()

	// Example 7: Worker with quit channel
	fmt.Println("7. Worker with quit channel:")
	jobs := make(chan int)
	quit := make(chan bool)

	go worker(jobs, quit)

	// Send some jobs
	for i := 1; i <= 3; i++ {
		jobs <- i
	}

	// Signal worker to quit
	quit <- true
	time.Sleep(50 * time.Millisecond)
	fmt.Println()

	// Example 8: Multiplexing (combining channels)
	fmt.Println("8. Multiplexing multiple input channels:")
	input1 := generator("source1", 3)
	input2 := generator("source2", 3)

	combined := fanIn(input1, input2)

	for i := 0; i < 6; i++ {
		fmt.Printf("   %s\n", <-combined)
	}
	fmt.Println()

	// Example 9: Priority select pattern
	fmt.Println("9. Priority select (checking high priority first):")
	highPriority := make(chan string, 1)
	lowPriority := make(chan string, 1)

	highPriority <- "urgent message"
	lowPriority <- "normal message"

	// Check high priority channel first
	select {
	case msg := <-highPriority:
		fmt.Printf("   High priority: %s\n", msg)
	default:
		select {
		case msg := <-lowPriority:
			fmt.Printf("   Low priority: %s\n", msg)
		default:
			fmt.Println("   No messages")
		}
	}
	fmt.Println()

	// Example 10: Graceful shutdown pattern
	fmt.Println("10. Graceful shutdown with context-like pattern:")
	data := make(chan int)
	stop := make(chan bool)

	go processor(data, stop)

	// Send some data
	for i := 1; i <= 5; i++ {
		data <- i
		time.Sleep(30 * time.Millisecond)
	}

	// Initiate shutdown
	stop <- true
	time.Sleep(50 * time.Millisecond)
}

// worker processes jobs until told to quit
func worker(jobs <-chan int, quit <-chan bool) {
	for {
		select {
		case job := <-jobs:
			fmt.Printf("   Processing job %d\n", job)
			time.Sleep(50 * time.Millisecond)
		case <-quit:
			fmt.Println("   Worker received quit signal")
			return
		}
	}
}

// generator creates a channel that produces messages
func generator(prefix string, count int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= count; i++ {
			ch <- fmt.Sprintf("%s: message %d", prefix, i)
			time.Sleep(50 * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}

// fanIn multiplexes multiple channels into one
func fanIn(inputs ...<-chan string) <-chan string {
	out := make(chan string)

	for _, input := range inputs {
		go func(ch <-chan string) {
			for msg := range ch {
				out <- msg
			}
		}(input)
	}

	return out
}

// processor handles data until stop signal is received
func processor(data <-chan int, stop <-chan bool) {
	for {
		select {
		case d := <-data:
			fmt.Printf("   Processing data: %d\n", d)
		case <-stop:
			fmt.Println("   Processor shutting down gracefully")
			return
		}
	}
}
