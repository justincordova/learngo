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
)

func main() {
	fmt.Println("Generic Types in Go")
	fmt.Println("===================")
	fmt.Println()

	// Example 1: Generic Stack
	fmt.Println("1. Generic Stack:")
	intStack := NewStack[int]()
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Printf("Stack size: %d\n", intStack.Size())
	if val, ok := intStack.Peek(); ok {
		fmt.Printf("Top element: %d\n", val)
	}

	for !intStack.IsEmpty() {
		val, _ := intStack.Pop()
		fmt.Printf("Popped: %d\n", val)
	}
	fmt.Println()

	// Example 2: Stack with strings
	fmt.Println("2. String Stack:")
	stringStack := NewStack[string]()
	stringStack.Push("first")
	stringStack.Push("second")
	stringStack.Push("third")

	for !stringStack.IsEmpty() {
		val, _ := stringStack.Pop()
		fmt.Printf("Popped: %s\n", val)
	}
	fmt.Println()

	// Example 3: Generic Linked List
	fmt.Println("3. Generic Linked List:")
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Prepend(0)

	fmt.Println("List contents:")
	list.Print()
	fmt.Printf("Length: %d\n", list.Length())
	fmt.Println()

	// Example 4: Linked List with structs
	fmt.Println("4. Linked List with User structs:")
	type User struct {
		ID   int
		Name string
	}

	userList := NewLinkedList[User]()
	userList.Append(User{ID: 1, Name: "Alice"})
	userList.Append(User{ID: 2, Name: "Bob"})
	userList.Append(User{ID: 3, Name: "Charlie"})

	fmt.Println("User list:")
	userList.Print()
	fmt.Println()

	// Example 5: Generic Pair (tuple-like structure)
	fmt.Println("5. Generic Pair:")
	p1 := Pair[string, int]{First: "age", Second: 25}
	p2 := Pair[string, string]{First: "language", Second: "Go"}
	p3 := Pair[int, bool]{First: 42, Second: true}

	fmt.Printf("Pair 1: %v = %v\n", p1.First, p1.Second)
	fmt.Printf("Pair 2: %v = %v\n", p2.First, p2.Second)
	fmt.Printf("Pair 3: %v = %v\n", p3.First, p3.Second)
}

// Stack is a generic LIFO data structure
type Stack[T any] struct {
	items []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
// Returns false if the stack is empty
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

// Peek returns the top item without removing it
// Returns false if the stack is empty
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack has no items
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Node is a node in a linked list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList is a generic singly-linked list
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// NewLinkedList creates a new empty linked list
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Append adds an item to the end of the list
func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
	l.size++
}

// Prepend adds an item to the beginning of the list
func (l *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value, Next: l.head}
	l.head = newNode

	if l.tail == nil {
		l.tail = newNode
	}
	l.size++
}

// Length returns the number of items in the list
func (l *LinkedList[T]) Length() int {
	return l.size
}

// Print displays all items in the list
func (l *LinkedList[T]) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// Pair holds two values of potentially different types
type Pair[T, U any] struct {
	First  T
	Second U
}

// Swap returns a new Pair with the values swapped
func (p Pair[T, U]) Swap() Pair[U, T] {
	return Pair[U, T]{First: p.Second, Second: p.First}
}
