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
	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println("Generic Functions in Go")
	fmt.Println("=======================")
	fmt.Println()

	// Example 1: Generic Min/Max functions
	fmt.Println("1. Min/Max with different types:")
	fmt.Printf("Min(5, 10) = %d\n", Min(5, 10))
	fmt.Printf("Max(5, 10) = %d\n", Max(5, 10))
	fmt.Printf("Min(3.14, 2.71) = %.2f\n", Min(3.14, 2.71))
	fmt.Printf("Max(3.14, 2.71) = %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Min(\"apple\", \"banana\") = %s\n", Min("apple", "banana"))
	fmt.Println()

	// Example 2: Generic Map function
	fmt.Println("2. Map function - transform slices:")
	numbers := []int{1, 2, 3, 4, 5}
	doubled := Map(numbers, func(n int) int { return n * 2 })
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Doubled: %v\n", doubled)

	words := []string{"hello", "world"}
	lengths := Map(words, func(s string) int { return len(s) })
	fmt.Printf("Words: %v\n", words)
	fmt.Printf("Lengths: %v\n", lengths)
	fmt.Println()

	// Example 3: Generic Filter function
	fmt.Println("3. Filter function - filter slices:")
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := Filter(allNumbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("All numbers: %v\n", allNumbers)
	fmt.Printf("Even numbers: %v\n", evens)

	names := []string{"Alice", "Bob", "Charlie", "David"}
	shortNames := Filter(names, func(s string) bool { return len(s) <= 4 })
	fmt.Printf("All names: %v\n", names)
	fmt.Printf("Short names: %v\n", shortNames)
	fmt.Println()

	// Example 4: Generic Reduce function
	fmt.Println("4. Reduce function - aggregate values:")
	nums := []int{1, 2, 3, 4, 5}
	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	product := Reduce(nums, 1, func(acc, n int) int { return acc * n })
	fmt.Printf("Numbers: %v\n", nums)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Product: %d\n", product)
	fmt.Println()

	// Example 5: Generic Contains function
	fmt.Println("5. Contains function - check membership:")
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v\n", intSlice)
	fmt.Printf("Contains 3? %v\n", Contains(intSlice, 3))
	fmt.Printf("Contains 10? %v\n", Contains(intSlice, 10))

	stringSlice := []string{"apple", "banana", "cherry"}
	fmt.Printf("Slice: %v\n", stringSlice)
	fmt.Printf("Contains \"banana\"? %v\n", Contains(stringSlice, "banana"))
	fmt.Printf("Contains \"grape\"? %v\n", Contains(stringSlice, "grape"))
}

// Min returns the smaller of two values
// Uses constraints.Ordered which includes all types that support < operator
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of two values
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Map transforms a slice by applying a function to each element
// Takes a slice of type T and a function that converts T to U
// Returns a new slice of type U
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter returns a new slice containing only elements that satisfy the predicate
// The predicate is a function that takes an element and returns true to keep it
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce aggregates slice elements into a single value
// Takes an initial accumulator value and a function that combines the accumulator with each element
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	acc := initial
	for _, v := range slice {
		acc = fn(acc, v)
	}
	return acc
}

// Contains checks if a value exists in a slice
// Uses comparable constraint since we need to use == operator
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}
