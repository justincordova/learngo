// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// Sentinel errors for demonstration
var (
	ErrNotFound      = errors.New("resource not found")
	ErrUnauthorized  = errors.New("unauthorized access")
	ErrInvalidInput  = errors.New("invalid input")
)

func main() {
	fmt.Println("Error Inspection Example")
	fmt.Println("========================\n")

	// Example 1: Using errors.Is to check for specific errors
	fmt.Println("1. Using errors.Is to check sentinel errors:")
	err1 := processResource(404)
	inspectWithIs(err1)

	// Example 2: Using errors.Is with wrapped errors
	fmt.Println("\n2. errors.Is works through wrapped errors:")
	err2 := fmt.Errorf("failed to complete operation: %w", ErrUnauthorized)
	inspectWithIs(err2)

	// Example 3: Using errors.As to extract specific error types
	fmt.Println("\n3. Using errors.As to extract error details:")
	err3 := readConfig("nonexistent.json")
	inspectWithAs(err3)

	// Example 4: Combining errors.Is and errors.As
	fmt.Println("\n4. Combining inspection methods:")
	err4 := loadData("data.txt")
	advancedInspection(err4)
}

// processResource simulates processing a resource and returns different errors
func processResource(id int) error {
	switch id {
	case 404:
		return fmt.Errorf("resource with id %d: %w", id, ErrNotFound)
	case 401:
		return fmt.Errorf("access denied for resource %d: %w", id, ErrUnauthorized)
	default:
		return nil
	}
}

// readConfig simulates reading a config file
func readConfig(filename string) error {
	_, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config from %q: %w", filename, err)
	}
	return nil
}

// loadData simulates loading data that might fail with file system errors
func loadData(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to load data: %w", err)
	}
	return nil
}

// inspectWithIs demonstrates using errors.Is to check for specific errors
func inspectWithIs(err error) {
	if err == nil {
		fmt.Println("No error occurred")
		return
	}

	fmt.Printf("Error: %v\n", err)

	// errors.Is checks if err or any error in its chain matches the target
	if errors.Is(err, ErrNotFound) {
		fmt.Println("  → This is a 'not found' error")
	} else if errors.Is(err, ErrUnauthorized) {
		fmt.Println("  → This is an 'unauthorized' error")
	} else if errors.Is(err, ErrInvalidInput) {
		fmt.Println("  → This is an 'invalid input' error")
	} else {
		fmt.Println("  → This is an unknown error type")
	}
}

// inspectWithAs demonstrates using errors.As to extract specific error types
func inspectWithAs(err error) {
	if err == nil {
		fmt.Println("No error occurred")
		return
	}

	fmt.Printf("Error: %v\n", err)

	// errors.As finds the first error in the chain that matches the target type
	// and assigns it to the target
	var pathErr *fs.PathError
	if errors.As(err, &pathErr) {
		fmt.Printf("  → PathError details:\n")
		fmt.Printf("    - Operation: %s\n", pathErr.Op)
		fmt.Printf("    - Path: %s\n", pathErr.Path)
		fmt.Printf("    - Underlying error: %v\n", pathErr.Err)
	}

	// You can check for os.PathError specifically (deprecated, use fs.PathError)
	var osPathErr *os.PathError
	if errors.As(err, &osPathErr) {
		fmt.Printf("  → This is also an os.PathError (legacy type)\n")
	}
}

// advancedInspection combines both inspection methods
func advancedInspection(err error) {
	if err == nil {
		fmt.Println("No error occurred")
		return
	}

	fmt.Printf("Error: %v\n", err)

	// First check if it's a known sentinel error
	if errors.Is(err, ErrNotFound) {
		fmt.Println("  → Resource not found")
		return
	}

	// Then try to extract more detailed error types
	var pathErr *fs.PathError
	if errors.As(err, &pathErr) {
		fmt.Printf("  → File system error:\n")
		fmt.Printf("    - Operation: %s\n", pathErr.Op)
		fmt.Printf("    - Path: %s\n", pathErr.Path)

		// Check if the underlying error is a specific type
		if errors.Is(pathErr.Err, fs.ErrNotExist) {
			fmt.Println("    - Reason: File does not exist")
		} else if errors.Is(pathErr.Err, fs.ErrPermission) {
			fmt.Println("    - Reason: Permission denied")
		}
	}
}
