// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

/*
EXERCISE: Generic Collection Operations

Implement the following generic functions for working with slices:

1. Filter[T any](slice []T, predicate func(T) bool) []T
   - Returns a new slice containing only elements that satisfy the predicate

2. Map[T, U any](slice []T, fn func(T) U) []U
   - Transforms each element using the provided function

3. Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U
   - Aggregates all elements into a single value

4. Find[T any](slice []T, predicate func(T) bool) (T, bool)
   - Returns the first element that satisfies the predicate

5. Any[T any](slice []T, predicate func(T) bool) bool
   - Returns true if at least one element satisfies the predicate

6. All[T any](slice []T, predicate func(T) bool) bool
   - Returns true if all elements satisfy the predicate

Requirements:
- All functions should work with any type
- Filter and Map should allocate new slices
- Find should return (zero value, false) if not found
- Handle empty slices appropriately

Test your implementation with:
- Numbers (filtering evens, mapping to squares, sum)
- Strings (filtering by length, mapping to uppercase, concatenation)
- Custom structs (filtering by field values)
*/

func main() {
	fmt.Println("Generic Collection Operations Exercise")
	fmt.Println("======================================")
	fmt.Println()

	// Test 1: Working with numbers
	fmt.Println("Test 1: Numbers")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// TODO: Filter even numbers
	// TODO: Map numbers to their squares
	// TODO: Calculate sum using Reduce
	// TODO: Find first number > 5
	// TODO: Check if any number is > 8
	// TODO: Check if all numbers are positive
	fmt.Println()

	// Test 2: Working with strings
	fmt.Println("Test 2: Strings")
	words := []string{"go", "rust", "python", "javascript", "c"}
	// TODO: Filter words with length > 3
	// TODO: Map words to their lengths
	// TODO: Concatenate all words using Reduce
	// TODO: Find first word starting with 'p'
	// TODO: Check if any word has length > 8
	fmt.Println()

	// Test 3: Working with structs
	fmt.Println("Test 3: Structs")
	type Product struct {
		Name  string
		Price float64
		Stock int
	}
	// TODO: Create a slice of products
	// TODO: Filter products with stock > 0
	// TODO: Map products to their names
	// TODO: Calculate total inventory value using Reduce
	// TODO: Find first product under $20
	// TODO: Check if all products have stock
}

// TODO: Implement the following functions:

// Filter returns a new slice containing only elements that satisfy the predicate
// func Filter[T any](slice []T, predicate func(T) bool) []T {
//     ...
// }

// Map transforms each element using the provided function
// func Map[T, U any](slice []T, fn func(T) U) []U {
//     ...
// }

// Reduce aggregates all elements into a single value
// func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
//     ...
// }

// Find returns the first element that satisfies the predicate
// func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
//     ...
// }

// Any returns true if at least one element satisfies the predicate
// func Any[T any](slice []T, predicate func(T) bool) bool {
//     ...
// }

// All returns true if all elements satisfy the predicate
// func All[T any](slice []T, predicate func(T) bool) bool {
//     ...
// }
