# Exercise: Generic Collection Operations

## Goal

Implement common functional programming operations (filter, map, reduce, etc.) as generic functions that work with any type.

## Requirements

Implement the following generic functions:

1. **Filter[T any](slice []T, predicate func(T) bool) []T**
   - Returns a new slice containing only elements where predicate returns true

2. **Map[T, U any](slice []T, fn func(T) U) []U**
   - Transforms each element from type T to type U using the provided function

3. **Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U**
   - Aggregates all elements into a single value starting from initial

4. **Find[T any](slice []T, predicate func(T) bool) (T, bool)**
   - Returns the first element matching the predicate (or zero value and false)

5. **Any[T any](slice []T, predicate func(T) bool) bool**
   - Returns true if at least one element satisfies the predicate

6. **All[T any](slice []T, predicate func(T) bool) bool**
   - Returns true if all elements satisfy the predicate (true for empty slices)

## Implementation Notes

- Use `any` constraint since these operations work with any type
- Filter and Map should allocate new slices
- Find should return `(zero value, false)` when no element is found
- All should return `true` for empty slices (vacuous truth)
- Any should return `false` for empty slices

## Test Cases

Your implementation should handle:

1. **Numbers**: Filter evens, map to squares, reduce to sum
2. **Strings**: Filter by length, map to lengths, reduce to concatenation
3. **Structs**: Filter by field values, map to specific fields, reduce to aggregate values

## Example Output

```
Generic Collection Operations Exercise - Solution
=================================================

Test 1: Numbers
Original: [1 2 3 4 5 6 7 8 9 10]
Even numbers: [2 4 6 8 10]
Squares: [1 4 9 16 25 36 49 64 81 100]
Sum: 55
First number > 5: 6
Any number > 8: true
All numbers positive: true

Test 2: Strings
Original: [go rust python javascript c]
Words with length > 3: [rust python javascript]
Word lengths: [2 4 6 10 1]
Concatenated: go, rust, python, javascript, c
First word starting with 'p': python
Any word with length > 8: true

Test 3: Structs
All products:
  Laptop: $999.99 (Stock: 5)
  Mouse: $19.99 (Stock: 50)
  Keyboard: $79.99 (Stock: 0)
  Monitor: $299.99 (Stock: 10)
  Webcam: $49.99 (Stock: 0)

In stock products:
  Laptop: $999.99 (Stock: 5)
  Mouse: $19.99 (Stock: 50)
  Monitor: $299.99 (Stock: 10)

Product names: [Laptop Mouse Keyboard Monitor Webcam]
Total inventory value: $9999.25
First product under $20: Mouse ($19.99)
All products have stock: false
```

## Running

```bash
# Run your solution
go run main.go

# Or check the reference solution
go run solution/main.go
```

## Learning Objectives

- Write generic functions with multiple type parameters
- Use function types as parameters (higher-order functions)
- Transform data between different types using Map
- Handle optional return values with the (T, bool) pattern
- Understand functional programming patterns in Go
- Work with predicates and transformation functions

## Hints

- **Filter**: Start with an empty slice and append matching elements
- **Map**: Pre-allocate the result slice with the same length as input
- **Reduce**: Use a loop to apply the function to accumulator and each element
- **Find**: Return immediately when you find a match
- **Any**: Return true as soon as you find one match
- **All**: Return false as soon as you find one non-match
