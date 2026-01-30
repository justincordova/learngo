// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Pipeline Processor
//
//  Create a data processing pipeline using chained channels.
//  This demonstrates the pipeline pattern in Go concurrency.
//
//  1- Create a pipeline with three stages:
//     Stage 1: Generate numbers from 1 to 10
//     Stage 2: Square each number
//     Stage 3: Filter only even results
//
//  2- Implement these functions:
//     - generate(nums ...int) <-chan int
//       Returns a channel that sends the input numbers
//
//     - square(in <-chan int) <-chan int
//       Receives numbers, squares them, sends results
//
//     - filterEven(in <-chan int) <-chan int
//       Receives numbers, filters even ones, sends results
//
//  3- In main:
//     - Chain the pipeline stages together
//     - Print the final results
//     - Each stage should run in its own goroutine
//
//  4- Bonus: Add a fourth stage that sums all the results
//     and prints the total
//
//
// EXPECTED OUTPUT:
//
//  Pipeline Processing Results:
//
//  Stage 1 (Generate): 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
//  Stage 2 (Square): 1, 4, 9, 16, 25, 36, 49, 64, 81, 100
//  Stage 3 (Filter Even): 4, 16, 36, 64, 100
//
//  Final Results: [4 16 36 64 100]
//  Sum: 220
//
// ---------------------------------------------------------

package main

func main() {
	// TODO: Implement the pipeline processor
}
