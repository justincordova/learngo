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
	"time"
)

// ValidationError represents an error that occurs during validation
// Custom error types allow you to attach additional context
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field %q (value: %v): %s",
		e.Field, e.Value, e.Message)
}

// DatabaseError represents a database operation error
type DatabaseError struct {
	Operation string
	Table     string
	Err       error // wrapped error
	Timestamp time.Time
}

// Error implements the error interface
func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error during %s on table %q at %s: %v",
		e.Operation, e.Table, e.Timestamp.Format(time.RFC3339), e.Err)
}

// Unwrap allows errors.Is and errors.As to work with wrapped errors
func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// NetworkError represents a network-related error
type NetworkError struct {
	Host    string
	Port    int
	Timeout time.Duration
	Err     error
}

// Error implements the error interface
func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error connecting to %s:%d (timeout: %v): %v",
		e.Host, e.Port, e.Timeout, e.Err)
}

// Unwrap allows unwrapping the underlying error
func (e *NetworkError) Unwrap() error {
	return e.Err
}

// Is allows errors.Is to match this error type
func (e *NetworkError) Is(target error) bool {
	t, ok := target.(*NetworkError)
	if !ok {
		return false
	}
	return e.Host == t.Host && e.Port == t.Port
}

func main() {
	fmt.Println("Custom Error Types Example")
	fmt.Println("===========================")
	fmt.Println()

	// Example 1: ValidationError
	fmt.Println("1. Validation error with custom fields:")
	err1 := validateUser("", 15)
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
		inspectValidationError(err1)
	}

	// Example 2: DatabaseError with wrapped error
	fmt.Println("\n2. Database error with wrapped context:")
	err2 := queryUser(999)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
		inspectDatabaseError(err2)
	}

	// Example 3: NetworkError with custom Is method
	fmt.Println("\n3. Network error with custom comparison:")
	err3 := connectToService("api.example.com", 443)
	if err3 != nil {
		fmt.Printf("Error: %v\n", err3)
		inspectNetworkError(err3)
	}

	// Example 4: Wrapping custom errors
	fmt.Println("\n4. Wrapped custom errors:")
	err4 := processUserRegistration("", 15)
	if err4 != nil {
		fmt.Printf("Error: %v\n", err4)
		inspectWrappedCustomError(err4)
	}
}

// validateUser validates user input and returns custom errors
func validateUser(name string, age int) error {
	if name == "" {
		return &ValidationError{
			Field:   "name",
			Value:   name,
			Message: "name cannot be empty",
		}
	}
	if age < 18 {
		return &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "must be at least 18 years old",
		}
	}
	return nil
}

// queryUser simulates a database query with custom error
func queryUser(id int) error {
	// Simulate a database error
	if id == 999 {
		baseErr := errors.New("record not found")
		return &DatabaseError{
			Operation: "SELECT",
			Table:     "users",
			Err:       baseErr,
			Timestamp: time.Now(),
		}
	}
	return nil
}

// connectToService simulates connecting to a remote service
func connectToService(host string, port int) error {
	// Simulate a connection error
	baseErr := errors.New("connection refused")
	return &NetworkError{
		Host:    host,
		Port:    port,
		Timeout: 30 * time.Second,
		Err:     baseErr,
	}
}

// processUserRegistration wraps custom errors
func processUserRegistration(name string, age int) error {
	if err := validateUser(name, age); err != nil {
		return fmt.Errorf("user registration failed: %w", err)
	}
	return nil
}

// inspectValidationError extracts ValidationError details
func inspectValidationError(err error) {
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("  → Validation Error Details:\n")
		fmt.Printf("    - Field: %s\n", validationErr.Field)
		fmt.Printf("    - Value: %v\n", validationErr.Value)
		fmt.Printf("    - Message: %s\n", validationErr.Message)
	}
}

// inspectDatabaseError extracts DatabaseError details
func inspectDatabaseError(err error) {
	var dbErr *DatabaseError
	if errors.As(err, &dbErr) {
		fmt.Printf("  → Database Error Details:\n")
		fmt.Printf("    - Operation: %s\n", dbErr.Operation)
		fmt.Printf("    - Table: %s\n", dbErr.Table)
		fmt.Printf("    - Timestamp: %s\n", dbErr.Timestamp.Format(time.RFC3339))
		fmt.Printf("    - Underlying error: %v\n", dbErr.Err)
	}
}

// inspectNetworkError extracts NetworkError details
func inspectNetworkError(err error) {
	var netErr *NetworkError
	if errors.As(err, &netErr) {
		fmt.Printf("  → Network Error Details:\n")
		fmt.Printf("    - Host: %s\n", netErr.Host)
		fmt.Printf("    - Port: %d\n", netErr.Port)
		fmt.Printf("    - Timeout: %v\n", netErr.Timeout)
		fmt.Printf("    - Underlying error: %v\n", netErr.Err)
	}

	// Demonstrate custom Is method
	targetErr := &NetworkError{Host: "api.example.com", Port: 443}
	if errors.Is(err, targetErr) {
		fmt.Println("  → This error matches the target host and port")
	}
}

// inspectWrappedCustomError shows how errors.As works with wrapped custom errors
func inspectWrappedCustomError(err error) {
	fmt.Printf("  → Inspecting wrapped error:\n")

	// errors.As can find the ValidationError even when it's wrapped
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("    - Found ValidationError in chain\n")
		fmt.Printf("    - Field: %s\n", validationErr.Field)
		fmt.Printf("    - Message: %s\n", validationErr.Message)
	}
}
