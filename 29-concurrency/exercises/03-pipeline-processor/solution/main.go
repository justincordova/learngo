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
	"strings"
)

func main() {
	fmt.Println("Pipeline Processing Results:")
	fmt.Println()

	// Input numbers
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Stage 1: Generate numbers
	stage1 := generate(nums...)
	generated := collectAndPrint(stage1, "Stage 1 (Generate)")

	// Stage 2: Square the numbers
	stage2 := square(generate(nums...))
	squared := collectAndPrint(stage2, "Stage 2 (Square)")

	// Stage 3: Filter even numbers
	stage3 := filterEven(square(generate(nums...)))
	filtered := collectAndPrint(stage3, "Stage 3 (Filter Even)")

	// Create final pipeline
	fmt.Println()
	final := filterEven(square(generate(nums...)))

	// Collect final results
	var results []int
	for n := range final {
		results = append(results, n)
	}

	// Print final results
	fmt.Printf("Final Results: %v\n", results)

	// Bonus: Sum all results
	sum := 0
	for _, n := range results {
		sum += n
	}
	fmt.Printf("Sum: %d\n", sum)

	// Avoid unused variable warnings
	_, _, _ = generated, squared, filtered
}

// generate sends numbers to a channel
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

// square receives numbers from a channel, squares them, and sends to output channel
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

// filterEven receives numbers and sends only even ones to output channel
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// collectAndPrint collects values from a channel and prints them
func collectAndPrint(ch <-chan int, label string) []int {
	var values []int
	for n := range ch {
		values = append(values, n)
	}

	// Convert to string slice for joining
	strValues := make([]string, len(values))
	for i, v := range values {
		strValues[i] = fmt.Sprintf("%d", v)
	}

	fmt.Printf("%s: %s\n", label, strings.Join(strValues, ", "))
	return values
}
