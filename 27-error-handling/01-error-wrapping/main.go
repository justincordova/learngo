package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Error Wrapping Example")
	fmt.Println("======================")
	fmt.Println()

	// Example 1: Reading a non-existent file
	fmt.Println("1. Attempting to read a non-existent file:")
	if err := processFile("nonexistent.txt"); err != nil {
		fmt.Printf("Error: %v\n\n", err)
	}

	// Example 2: Multiple layers of wrapping
	fmt.Println("2. Multiple layers of error wrapping:")
	if err := processUserData(999); err != nil {
		fmt.Printf("Error: %v\n\n", err)
	}

	// Example 3: Successful operation
	fmt.Println("3. Successful operation:")
	if err := processUserData(42); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("User data processed successfully")
	}
}

// processFile attempts to read a file and wraps any errors
func processFile(filename string) error {
	_, err := os.ReadFile(filename)
	if err != nil {
		// Wrap the error with context using %w
		// This preserves the original error in the error chain
		return fmt.Errorf("failed to process file %q: %w", filename, err)
	}
	return nil
}

// processUserData demonstrates multiple layers of error wrapping
func processUserData(userID int) error {
	if err := loadUser(userID); err != nil {
		// Add another layer of context
		return fmt.Errorf("failed to process user data: %w", err)
	}
	return nil
}

// loadUser simulates loading user data
func loadUser(userID int) error {
	// Simulate a database error for user ID 999
	if userID == 999 {
		err := fmt.Errorf("user not found")
		// Wrap with additional context
		return fmt.Errorf("failed to load user %d from database: %w", userID, err)
	}
	return nil
}
