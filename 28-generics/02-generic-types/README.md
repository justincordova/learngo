# Generic Types

Generic types allow you to create data structures that work with any type while maintaining type safety. This is particularly useful for collections like stacks, queues, linked lists, and trees.

## Syntax

```go
type TypeName[T TypeConstraint] struct {
    field T
}
```

- The type parameter is specified in square brackets after the type name
- Methods on generic types also use the type parameter
- You can have multiple type parameters: `type Pair[T, U any] struct { ... }`

## Common Generic Data Structures

### Stack (LIFO)

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
    // ...
}
```

### Linked List

```go
type Node[T any] struct {
    Value T
    Next  *Node[T]
}

type LinkedList[T any] struct {
    head *Node[T]
    tail *Node[T]
}
```

### Pair (Tuple)

```go
type Pair[T, U any] struct {
    First  T
    Second U
}
```

## Key Points

- Generic types can have methods just like regular types
- Methods on generic types automatically inherit the type parameters
- Constructor functions should also be generic: `func NewStack[T any]() *Stack[T]`
- The zero value of a type parameter is `var zero T`
- Type parameters can be used in field types, method receivers, and return types

## Running This Example

```bash
cd 02-generic-types
go run main.go
```

## Expected Output

The program demonstrates:
- Stack operations with integers and strings
- Linked list operations with integers and custom structs
- Pair type holding different type combinations

## Use Cases

Generic types are ideal for:
- **Collections**: Lists, stacks, queues, sets, trees
- **Wrappers**: Optional values, Result types, smart pointers
- **Data structures**: Graphs, heaps, tries
- **Containers**: Caches, pools, registries

## Next Steps

See [03-type-constraints](../03-type-constraints/) to learn about limiting which types can be used with your generics.
