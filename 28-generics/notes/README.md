# Generics Notes

## Introduction

Generics (type parameters) were introduced in Go 1.18, bringing type-safe reusable code without sacrificing performance. Generics allow you to write functions and data structures that work with multiple types while maintaining compile-time type safety.

## Type Parameters

Type parameters allow functions to work with multiple types:

```go
func Print[T any](value T) {
    fmt.Println(value)
}

// Usage
Print[int](42)
Print[string]("hello")
Print(42)        // Type inference - compiler determines T is int
Print("hello")   // Type inference - compiler determines T is string
```

## Type Constraints

### The any Constraint

`any` is an alias for `interface{}` and allows any type:

```go
func First[T any](slice []T) (T, bool) {
    if len(slice) == 0 {
        var zero T
        return zero, false
    }
    return slice[0], true
}
```

### The comparable Constraint

`comparable` allows types that can be compared with `==` and `!=`:

```go
func Contains[T comparable](slice []T, target T) bool {
    for _, v := range slice {
        if v == target {
            return true
        }
    }
    return false
}
```

### Custom Constraints

Define your own constraints using interfaces:

```go
type Number interface {
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 |
    float32 | float64
}

func Sum[T Number](numbers []T) T {
    var total T
    for _, n := range numbers {
        total += n
    }
    return total
}
```

### Using constraints Package

The `golang.org/x/exp/constraints` package provides common constraints:

```go
import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

## Generic Types

Types can also be parameterized:

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Usage
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)

stringStack := Stack[string]{}
stringStack.Push("hello")
```

## Common Generic Patterns

### Map Function

```go
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Usage
numbers := []int{1, 2, 3}
doubled := Map(numbers, func(n int) int { return n * 2 })
// doubled is []int{2, 4, 6}

strings := Map(numbers, func(n int) string { return fmt.Sprint(n) })
// strings is []string{"1", "2", "3"}
```

### Filter Function

```go
func Filter[T any](slice []T, predicate func(T) bool) []T {
    result := make([]T, 0)
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4, 5, 6}
evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
// evens is []int{2, 4, 6}
```

### Generic Queue

```go
type Queue[T any] struct {
    items []T
}

func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}
```

## When to Use Generics vs Interfaces

### Use Generics When:
- Writing data structures (lists, stacks, queues, trees)
- Implementing algorithms that work on multiple types (sorting, searching)
- Creating utility functions that don't depend on type-specific behavior
- Avoiding `interface{}` and type assertions for better type safety

### Use Interfaces When:
- You need runtime polymorphism
- Different types should have different behaviors
- You're working with existing code that uses interfaces
- The abstraction is about behavior, not type

### Example: Generics vs Interfaces

**Generics** - When the operation is the same for all types:
```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

**Interfaces** - When different types have different implementations:
```go
type Serializer interface {
    Serialize() ([]byte, error)
}

func SaveToFile(s Serializer, filename string) error {
    data, err := s.Serialize()
    // ...
}
```

## Best Practices

### Do:
- Use meaningful type parameter names (T for a single type, K/V for key/value)
- Keep generic functions simple and focused
- Use the most specific constraint that makes sense
- Document type parameter constraints clearly
- Consider if generics add value over interfaces

### Don't:
- Use generics just to avoid writing a few extra functions
- Create overly complex generic types
- Use `any` when a more specific constraint would work
- Forget that generics are a compile-time feature (no runtime type information)
- Over-engineer with unnecessary type parameters

## Type Inference

Go can often infer type parameters:

```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Explicit type parameter
result := Max[int](5, 10)

// Type inference (preferred)
result := Max(5, 10)  // Compiler infers T is int
```

## Zero Values in Generics

When you need a zero value for a generic type:

```go
func GetOrDefault[T any](m map[string]T, key string) T {
    if val, ok := m[key]; ok {
        return val
    }
    var zero T  // Zero value for type T
    return zero
}
```

## Resources

- [Go Blog: An Introduction To Generics](https://go.dev/blog/intro-generics)
- [Go Blog: When To Use Generics](https://go.dev/blog/when-generics)
- [Go Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
- [constraints package documentation](https://pkg.go.dev/golang.org/x/exp/constraints)
