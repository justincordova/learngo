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
EXERCISE: Generic Stack with Additional Operations

Create a generic Stack type with the following operations:
1. Push(item T) - add an item to the stack
2. Pop() (T, bool) - remove and return the top item
3. Peek() (T, bool) - return the top item without removing it
4. IsEmpty() bool - check if the stack is empty
5. Size() int - return the number of items
6. Clear() - remove all items
7. ToSlice() []T - return all items as a slice (top to bottom)

Requirements:
- The Stack should work with any type
- Pop and Peek should return (zero value, false) when empty
- ToSlice should return items in order from top to bottom

Test your implementation with:
- A stack of integers
- A stack of strings
- A stack of custom struct types
*/

func main() {
	fmt.Println("Generic Stack Exercise")
	fmt.Println("======================")
	fmt.Println()

	// Test 1: Integer stack
	fmt.Println("Test 1: Integer Stack")
	// TODO: Create a stack of integers
	// TODO: Push 1, 2, 3, 4, 5
	// TODO: Print the size
	// TODO: Peek at the top element
	// TODO: Pop all elements and print them
	// TODO: Try to pop from an empty stack
	fmt.Println()

	// Test 2: String stack
	fmt.Println("Test 2: String Stack")
	// TODO: Create a stack of strings
	// TODO: Push "first", "second", "third"
	// TODO: Print all items using ToSlice()
	// TODO: Clear the stack
	// TODO: Check if it's empty
	fmt.Println()

	// Test 3: Struct stack
	fmt.Println("Test 3: Struct Stack")
	type Task struct {
		ID       int
		Name     string
		Priority int
	}
	// TODO: Create a stack of Task structs
	// TODO: Push several tasks with different priorities
	// TODO: Pop and process tasks (LIFO order)
}

// TODO: Implement the Stack type and its methods below
// type Stack[T any] struct {
//     ...
// }
