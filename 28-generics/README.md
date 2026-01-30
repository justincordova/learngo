# Generics in Go

Generics (type parameters) were introduced in Go 1.18, bringing type-safe reusable code without sacrificing performance. Generics allow you to write functions and data structures that work with multiple types while maintaining compile-time type safety.

## Overview

This section covers Go's generics feature and best practices:

- **Generic Functions**: Writing functions that work with multiple types
- **Generic Types**: Creating type-parameterized data structures
- **Type Constraints**: Defining and using constraints to limit type parameters
- **When to Use Generics**: Understanding when generics add value vs interfaces

## Prerequisites

Before starting this section, you should be comfortable with:

- Basic Go syntax and control flow
- Functions and method declarations
- Interfaces and type assertions
- Structs and composite types
- The concept of type safety

## Key Concepts

### Type Parameters

Type parameters allow functions and types to work with multiple types:

```go
func Print[T any](value T) {
    fmt.Println(value)
}
```

### Type Constraints

Constraints specify what operations are allowed on type parameters:

```go
func Min[T comparable](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

### Generic Types

Types can also be parameterized:

```go
type Stack[T any] struct {
    items []T
}
```

## Section Contents

1. **[Generic Functions](01-generic-functions/)** - Learn to write functions with type parameters

2. **[Generic Types](02-generic-types/)** - Create data structures that work with any type

3. **[Type Constraints](03-type-constraints/)** - Define and use constraints to limit type parameters

4. **[Exercises](exercises/)** - Practice working with generics

## When to Use Generics

### Use Generics When:
- Writing data structures (lists, stacks, queues, trees)
- Implementing algorithms that work on multiple types (sorting, searching)
- Creating utility functions that don't depend on type-specific behavior
- Avoiding interface{} and type assertions for better type safety

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

## Common Patterns

### Generic Slice Operations

```go
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
    result := make([]T, 0)
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}
```

### Generic Data Structures

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

### Custom Constraints

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

## Resources

- [Go Blog: An Introduction To Generics](https://go.dev/blog/intro-generics)
- [Go Blog: When To Use Generics](https://go.dev/blog/when-generics)
- [Go Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
- [constraints package documentation](https://pkg.go.dev/golang.org/x/exp/constraints)

## Next Steps

After completing this section, you'll be ready to:
- Write type-safe generic functions and data structures
- Choose between generics and interfaces appropriately
- Create custom type constraints for your use cases
- Build reusable generic utilities for your applications
