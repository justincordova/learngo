# Maps

## Why Use Maps?

Maps provide **fast key-based lookup** with O(1) average time complexity, making them significantly more efficient than searching through slices or arrays.

### Inefficient Linear Search

```go
millions := []int{/* millions of elements */}
for _, v := range millions {
    if v == userQuery {
        // do something - O(n) time complexity
    }
}
```

This approach requires checking each element sequentially, resulting in O(n) time complexity.

### Efficient Map Lookup

Maps allow you to find values by key in constant time on average (O(1)), regardless of how many elements the map contains. This makes maps ideal for scenarios where you need to quickly look up values based on unique keys.

## When NOT to Use Maps

While maps excel at key-based lookups, they are **not optimal for traversal**:

```go
for k, v := range myMap {
    // This happens in O(n) time
}
```

Looping over map keys has O(n) time complexity, making maps inefficient for operations that require iterating through all elements. If your primary use case is traversal rather than lookup, consider using a slice or array instead.

**Note:** Maps also don't provide structured data organization in the same way structs do. For grouping related fields with known names, use structs.

## Map Key Constraints

Not all types can be used as map keys. **Map keys must be comparable** using `==` and `!=`.

### Invalid Key Types

The following types cannot be map keys:

```go
map[[]string]int   // Error: slices are not comparable
map[[]int]bool     // Error: slices are not comparable
map[[]bool]string  // Error: slices are not comparable
```

**Invalid types for keys:**
- Slices (`[]T`)
- Maps (`map[K]V`)
- Function values (`func(...)`)

These types are not comparable, so they cannot serve as map keys.

### Valid Key Types

Any comparable type can be a map key, including:
- Basic types: `string`, `int`, `bool`, `float64`, etc.
- Structs (if all their fields are comparable)
- Arrays (if their element type is comparable)
- Pointers
- Interfaces

## Map Type Syntax

A map type is defined as `map[KeyType]ElementType`:

```go
map[string]map[int]bool
```

**Breaking this down:**
- **Key type:** `string`
- **Element type:** `map[int]bool`

The element type can be any type, including another map. This allows you to create nested map structures for more complex data relationships.

## Map Values Are Pointers

Behind the scenes, a map value is **a pointer to a map header**. The map header is a complex data structure that manages the actual storage and lookup mechanism.

**Practical implications:**
- When you assign a map to another variable, both variables reference the same underlying map
- Modifying the map through one variable affects all variables that reference it
- This is different from arrays, which are copied when assigned

```go
map1 := make(map[string]int)
map1["a"] = 1

map2 := map1  // map2 references the same underlying map
map2["b"] = 2

// Both map1 and map2 now contain both "a" and "b"
```

This pointer-based behavior is why maps don't need to be passed as pointers to functions—the map value itself is already a reference to the underlying data structure.
