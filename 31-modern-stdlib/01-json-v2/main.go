package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents a user in our system
type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Age      int      `json:"age,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Active   bool     `json:"active"`
}

func main() {
	fmt.Println("JSON v1 vs v2 Comparison")
	fmt.Println("========================")
	fmt.Println()
	
	// Note: This example uses encoding/json (v1)
	// To use encoding/json/v2, you need Go 1.25+ with GOEXPERIMENT=jsonv2
	
	demonstrateJSONv1()
	fmt.Println()
	demonstrateJSONv2Note()
}

func demonstrateJSONv1() {
	fmt.Println("Using encoding/json (v1):")
	fmt.Println("--------------------------")
	
	// Sample data
	user := User{
		ID:     1,
		Name:   "Alice",
		Email:  "alice@example.com",
		Age:    30,
		Tags:   []string{"developer", "golang"},
		Active: true,
	}
	
	// Marshal to JSON
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Marshaled: %s\n", data)
	
	// Unmarshal from JSON
	var decoded User
	if err := json.Unmarshal(data, &decoded); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Unmarshaled: %+v\n", decoded)
	
	// Pretty print
	pretty, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("\nPretty printed:\n%s\n", pretty)
}

func demonstrateJSONv2Note() {
	fmt.Println("About encoding/json/v2 (Go 1.25+):")
	fmt.Println("-----------------------------------")
	fmt.Println("The new json/v2 package offers:")
	fmt.Println("  • Substantially faster decoding")
	fmt.Println("  • Encoding performance at parity with v1")
	fmt.Println("  • Better error messages")
	fmt.Println("  • More control over encoding/decoding")
	fmt.Println("  • Lower-level jsontext package for streaming")
	fmt.Println()
	fmt.Println("To use json/v2:")
	fmt.Println("  1. Install Go 1.25+")
	fmt.Println("  2. Build with: GOEXPERIMENT=jsonv2 go build")
	fmt.Println("  3. Import: encoding/json/v2")
	fmt.Println()
	fmt.Println("Note: json/v2 is experimental and API may change")
}
