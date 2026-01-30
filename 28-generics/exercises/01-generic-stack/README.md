# Exercise: Generic Stack with Additional Operations

## Goal

Implement a generic Stack data structure with comprehensive operations.

## Requirements

Create a `Stack[T any]` type with the following methods:

1. **Push(item T)** - Add an item to the top of the stack
2. **Pop() (T, bool)** - Remove and return the top item (returns false if empty)
3. **Peek() (T, bool)** - View the top item without removing it (returns false if empty)
4. **IsEmpty() bool** - Check if the stack has no items
5. **Size() int** - Return the number of items in the stack
6. **Clear()** - Remove all items from the stack
7. **ToSlice() []T** - Return all items as a slice from top to bottom

## Implementation Notes

- The stack should work with **any type** (use `any` constraint)
- When the stack is empty, `Pop()` and `Peek()` should return the zero value and `false`
- `ToSlice()` should return items in LIFO order (most recently pushed first)
- Use a slice as the underlying storage

## Test Cases

Your implementation should handle:

1. **Integer stack**: Push numbers, pop them, verify LIFO order
2. **String stack**: Push strings, use ToSlice(), clear the stack
3. **Struct stack**: Create a Task struct with ID, Name, and Priority fields, push multiple tasks, pop and process them

## Example Output

```
Generic Stack Exercise - Solution
==================================

Test 1: Integer Stack
Stack size: 5
Top element: 5
Popping all elements:
  Popped: 5
  Popped: 4
  Popped: 3
  Popped: 2
  Popped: 1
Try to pop from empty stack:
  Stack is empty, got zero value: 0

Test 2: String Stack
All items (top to bottom):
  0: third
  1: second
  2: first
After clear, is empty: true

Test 3: Struct Stack
Processing tasks (LIFO order):
  Processing: [ID:4] Urgent task (Priority: 4)
  Processing: [ID:3] High priority task (Priority: 3)
  Processing: [ID:2] Medium priority task (Priority: 2)
  Processing: [ID:1] Low priority task (Priority: 1)
```

## Running

```bash
# Run your solution
go run main.go

# Or check the reference solution
go run solution/main.go
```

## Learning Objectives

- Create generic types that work with any type
- Implement multiple methods on a generic type
- Handle edge cases (empty stack)
- Return meaningful values when operations fail
- Work with the zero value of generic types
