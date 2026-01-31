// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

//go:build goexperiment.jsonv2

package main

import (
	"bytes"
	"encoding/json"
	jsonv2 "encoding/json/v2"
	"fmt"
	"strings"
	"time"
)

// User demonstrates various JSON v2 features
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age,omitzero"`        // Omit if zero
	Balance   float64   `json:"balance,string"`      // Stringify numbers
	CreatedAt time.Time `json:"created_at"`
	Tags      []string  `json:"tags,omitempty"`      // Omit if empty
	Active    bool      `json:"active"`
}

// Product demonstrates inline and unknown fields
type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Metadata struct {
		Category string `json:"category"`
		Brand    string `json:"brand"`
	} `json:",inline"` // Inline (embed) in parent object
}

func main() {
	fmt.Println("JSON v2 (encoding/json/v2) Examples")
	fmt.Println("====================================")
	fmt.Println()
	fmt.Println("Note: Requires GOEXPERIMENT=jsonv2 build flag")
	fmt.Println()

	// Example 1: Basic Marshal/Unmarshal comparison
	example1()

	// Example 2: MarshalWrite for streaming
	example2()

	// Example 3: UnmarshalRead for streaming
	example3()

	// Example 4: Advanced struct tags
	example4()

	// Example 5: Performance comparison
	example5()

	// Example 6: Security features
	example6()
}

func example1() {
	fmt.Println("1. Basic Marshal/Unmarshal with v2")
	fmt.Println("-----------------------------------")

	user := User{
		ID:        1,
		Name:      "Alice Johnson",
		Email:     "alice@example.com",
		Age:       0, // Will be omitted due to omitzero
		Balance:   1234.56,
		CreatedAt: time.Now(),
		Active:    true,
	}

	// Using json/v2 Marshal
	data, err := jsonv2.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}
	fmt.Printf("JSON v2 output:\n%s\n\n", data)

	// Unmarshal back
	var decoded User
	if err := jsonv2.Unmarshal(data, &decoded); err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}
	fmt.Printf("Decoded user: %+v\n\n", decoded)
}

func example2() {
	fmt.Println("2. MarshalWrite for Streaming")
	fmt.Println("------------------------------")

	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Balance: 100.50, CreatedAt: time.Now(), Active: true},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Balance: 250.75, CreatedAt: time.Now(), Active: true},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", Balance: 75.25, CreatedAt: time.Now(), Active: false},
	}

	// Write directly to a buffer using MarshalWrite
	// This is more efficient than creating an Encoder
	var buf bytes.Buffer
	for i, user := range users {
		if err := jsonv2.MarshalWrite(&buf, user); err != nil {
			fmt.Printf("Error writing user %d: %v\n", i, err)
			continue
		}
		buf.WriteByte('\n') // Add newline between records
	}

	fmt.Printf("Streamed output:\n%s\n", buf.String())
}

func example3() {
	fmt.Println("3. UnmarshalRead for Streaming")
	fmt.Println("-------------------------------")

	// Simulate reading from a stream
	jsonData := `{"id":42,"name":"Dave","email":"dave@example.com","balance":"999.99","created_at":"2025-01-15T10:30:00Z","active":true}`
	reader := strings.NewReader(jsonData)

	// Read directly from the reader using UnmarshalRead
	var user User
	if err := jsonv2.UnmarshalRead(reader, &user); err != nil {
		fmt.Printf("Error reading: %v\n", err)
		return
	}

	fmt.Printf("Read from stream: %+v\n\n", user)
}

func example4() {
	fmt.Println("4. Advanced Struct Tags (Inline)")
	fmt.Println("---------------------------------")

	product := Product{
		Name:  "Laptop",
		Price: 1299.99,
	}
	product.Metadata.Category = "Electronics"
	product.Metadata.Brand = "TechCorp"

	data, err := jsonv2.Marshal(product)
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}

	// Note: inline flattens the Metadata fields into the parent object
	fmt.Printf("Product JSON (inline metadata):\n%s\n\n", data)
}

func example5() {
	fmt.Println("5. Performance Comparison")
	fmt.Println("-------------------------")

	largeData := make([]User, 100)
	for i := range largeData {
		largeData[i] = User{
			ID:        i + 1,
			Name:      fmt.Sprintf("User %d", i+1),
			Email:     fmt.Sprintf("user%d@example.com", i+1),
			Balance:   float64(i * 100),
			CreatedAt: time.Now(),
			Active:    i%2 == 0,
		}
	}

	// Test v1 (standard encoding/json)
	startV1 := time.Now()
	_, err := json.Marshal(largeData)
	durationV1 := time.Since(startV1)
	if err != nil {
		fmt.Printf("Error with v1: %v\n", err)
	}
	fmt.Printf("encoding/json (v1):    %v\n", durationV1)

	// Test v2 (encoding/json/v2)
	startV2 := time.Now()
	_, err = jsonv2.Marshal(largeData)
	durationV2 := time.Since(startV2)
	if err != nil {
		fmt.Printf("Error with v2: %v\n", err)
	}
	fmt.Printf("encoding/json/v2:      %v\n", durationV2)

	if durationV2 < durationV1 {
		improvement := float64(durationV1-durationV2) / float64(durationV1) * 100
		fmt.Printf("v2 is %.1f%% faster\n\n", improvement)
	} else {
		fmt.Println("(Results may vary; v2 typically excels with large datasets)")
		fmt.Println()
	}
}

func example6() {
	fmt.Println("6. Security Features")
	fmt.Println("--------------------")

	// JSON with duplicate keys (security concern)
	duplicateJSON := `{"id":1,"name":"Alice","name":"Bob"}`

	// v2 rejects duplicates by default (more secure)
	var user User
	err := jsonv2.Unmarshal([]byte(duplicateJSON), &user)
	if err != nil {
		fmt.Printf("v2 correctly rejects duplicate keys: %v\n", err)
	} else {
		fmt.Printf("Unexpected: v2 accepted duplicate keys\n")
	}

	// v1 silently accepts duplicates (uses last value)
	err = json.Unmarshal([]byte(duplicateJSON), &user)
	if err != nil {
		fmt.Printf("v1 error: %v\n", err)
	} else {
		fmt.Printf("v1 silently accepted duplicate keys (name=%s)\n", user.Name)
	}
	fmt.Println()

	// JSON with unknown fields
	unknownJSON := `{"id":2,"name":"Charlie","unknown_field":"value"}`

	// Can configure v2 to reject unknown fields for strict validation
	var strictUser User
	err = jsonv2.Unmarshal([]byte(unknownJSON), &strictUser,
		jsonv2.RejectUnknownMembers(true))
	if err != nil {
		fmt.Printf("v2 with RejectUnknownMembers rejects unknown fields: %v\n", err)
	} else {
		fmt.Printf("Unexpected: accepted unknown fields\n")
	}

	fmt.Println()
}
