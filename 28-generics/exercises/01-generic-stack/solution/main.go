package main

import "fmt"

func main() {
	fmt.Println("Generic Stack Exercise - Solution")
	fmt.Println("==================================")
	fmt.Println()

	// Test 1: Integer stack
	fmt.Println("Test 1: Integer Stack")
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	intStack.Push(4)
	intStack.Push(5)

	fmt.Printf("Stack size: %d\n", intStack.Size())

	if top, ok := intStack.Peek(); ok {
		fmt.Printf("Top element: %d\n", top)
	}

	fmt.Println("Popping all elements:")
	for !intStack.IsEmpty() {
		if val, ok := intStack.Pop(); ok {
			fmt.Printf("  Popped: %d\n", val)
		}
	}

	fmt.Println("Try to pop from empty stack:")
	if val, ok := intStack.Pop(); !ok {
		fmt.Printf("  Stack is empty, got zero value: %d\n", val)
	}
	fmt.Println()

	// Test 2: String stack
	fmt.Println("Test 2: String Stack")
	stringStack := NewStack[string]()
	stringStack.Push("first")
	stringStack.Push("second")
	stringStack.Push("third")

	fmt.Println("All items (top to bottom):")
	items := stringStack.ToSlice()
	for i, item := range items {
		fmt.Printf("  %d: %s\n", i, item)
	}

	stringStack.Clear()
	fmt.Printf("After clear, is empty: %v\n", stringStack.IsEmpty())
	fmt.Println()

	// Test 3: Struct stack
	fmt.Println("Test 3: Struct Stack")
	type Task struct {
		ID       int
		Name     string
		Priority int
	}

	taskStack := NewStack[Task]()
	taskStack.Push(Task{ID: 1, Name: "Low priority task", Priority: 1})
	taskStack.Push(Task{ID: 2, Name: "Medium priority task", Priority: 2})
	taskStack.Push(Task{ID: 3, Name: "High priority task", Priority: 3})
	taskStack.Push(Task{ID: 4, Name: "Urgent task", Priority: 4})

	fmt.Println("Processing tasks (LIFO order):")
	for !taskStack.IsEmpty() {
		if task, ok := taskStack.Pop(); ok {
			fmt.Printf("  Processing: [ID:%d] %s (Priority: %d)\n",
				task.ID, task.Name, task.Priority)
		}
	}
}

// Stack is a generic LIFO (Last In, First Out) data structure
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
// Returns (zero value, false) if the stack is empty
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
// Returns (zero value, false) if the stack is empty
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

// Clear removes all items from the stack
func (s *Stack[T]) Clear() {
	s.items = make([]T, 0)
}

// ToSlice returns all items as a slice from top to bottom
func (s *Stack[T]) ToSlice() []T {
	result := make([]T, len(s.items))
	for i := 0; i < len(s.items); i++ {
		result[i] = s.items[len(s.items)-1-i]
	}
	return result
}
