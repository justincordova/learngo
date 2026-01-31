package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Generic Collection Operations Exercise - Solution")
	fmt.Println("=================================================")
	fmt.Println()

	// Test 1: Working with numbers
	fmt.Println("Test 1: Numbers")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original: %v\n", numbers)

	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evens)

	squares := Map(numbers, func(n int) int { return n * n })
	fmt.Printf("Squares: %v\n", squares)

	sum := Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	fmt.Printf("Sum: %d\n", sum)

	if found, ok := Find(numbers, func(n int) bool { return n > 5 }); ok {
		fmt.Printf("First number > 5: %d\n", found)
	}

	hasLarge := Any(numbers, func(n int) bool { return n > 8 })
	fmt.Printf("Any number > 8: %v\n", hasLarge)

	allPositive := All(numbers, func(n int) bool { return n > 0 })
	fmt.Printf("All numbers positive: %v\n", allPositive)
	fmt.Println()

	// Test 2: Working with strings
	fmt.Println("Test 2: Strings")
	words := []string{"go", "rust", "python", "javascript", "c"}
	fmt.Printf("Original: %v\n", words)

	longWords := Filter(words, func(s string) bool { return len(s) > 3 })
	fmt.Printf("Words with length > 3: %v\n", longWords)

	lengths := Map(words, func(s string) int { return len(s) })
	fmt.Printf("Word lengths: %v\n", lengths)

	concatenated := Reduce(words, "", func(acc, s string) string {
		if acc == "" {
			return s
		}
		return acc + ", " + s
	})
	fmt.Printf("Concatenated: %s\n", concatenated)

	if found, ok := Find(words, func(s string) bool { return strings.HasPrefix(s, "p") }); ok {
		fmt.Printf("First word starting with 'p': %s\n", found)
	}

	hasLongWord := Any(words, func(s string) bool { return len(s) > 8 })
	fmt.Printf("Any word with length > 8: %v\n", hasLongWord)
	fmt.Println()

	// Test 3: Working with structs
	fmt.Println("Test 3: Structs")
	type Product struct {
		Name  string
		Price float64
		Stock int
	}

	products := []Product{
		{Name: "Laptop", Price: 999.99, Stock: 5},
		{Name: "Mouse", Price: 19.99, Stock: 50},
		{Name: "Keyboard", Price: 79.99, Stock: 0},
		{Name: "Monitor", Price: 299.99, Stock: 10},
		{Name: "Webcam", Price: 49.99, Stock: 0},
	}

	fmt.Println("All products:")
	for _, p := range products {
		fmt.Printf("  %s: $%.2f (Stock: %d)\n", p.Name, p.Price, p.Stock)
	}

	inStock := Filter(products, func(p Product) bool { return p.Stock > 0 })
	fmt.Println("\nIn stock products:")
	for _, p := range inStock {
		fmt.Printf("  %s: $%.2f (Stock: %d)\n", p.Name, p.Price, p.Stock)
	}

	names := Map(products, func(p Product) string { return p.Name })
	fmt.Printf("\nProduct names: %v\n", names)

	totalValue := Reduce(products, 0.0, func(acc float64, p Product) float64 {
		return acc + (p.Price * float64(p.Stock))
	})
	fmt.Printf("Total inventory value: $%.2f\n", totalValue)

	if found, ok := Find(products, func(p Product) bool { return p.Price < 20 }); ok {
		fmt.Printf("First product under $20: %s ($%.2f)\n", found.Name, found.Price)
	}

	allHaveStock := All(products, func(p Product) bool { return p.Stock > 0 })
	fmt.Printf("All products have stock: %v\n", allHaveStock)
}

// Filter returns a new slice containing only elements that satisfy the predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map transforms each element using the provided function
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Reduce aggregates all elements into a single value
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	acc := initial
	for _, v := range slice {
		acc = fn(acc, v)
	}
	return acc
}

// Find returns the first element that satisfies the predicate
// Returns (zero value, false) if no element is found
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, v := range slice {
		if predicate(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

// Any returns true if at least one element satisfies the predicate
func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements satisfy the predicate
// Returns true for empty slices
func All[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if !predicate(v) {
			return false
		}
	}
	return true
}
