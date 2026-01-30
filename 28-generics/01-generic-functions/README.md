# Generic Functions

Generic functions use type parameters to work with multiple types while maintaining type safety. This is useful for writing reusable algorithms that don't depend on specific types.

## Syntax

```go
func FunctionName[T TypeConstraint](param T) T {
    // function body
}
```

- `[T TypeConstraint]` defines the type parameter(s)
- `T` is the type parameter name (can be any identifier)
- `TypeConstraint` limits what types can be used
- Multiple type parameters: `[T any, U any]`

## Common Use Cases

### 1. Min/Max Functions

Instead of writing separate functions for each type:

```go
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

### 2. Slice Operations

Functions that transform or filter slices:

```go
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}
```

### 3. Contains/Search

Check if an element exists:

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

## Key Points

- Type parameters are specified in square brackets `[T TypeConstraint]`
- The constraint determines what operations you can perform on the type
- `any` means no constraints (equivalent to `interface{}`)
- `comparable` allows use of `==` and `!=`
- `constraints.Ordered` allows comparison operators (`<`, `>`, etc.)
- Multiple type parameters are separated by commas

## Running This Example

```bash
cd 01-generic-functions
go mod init example  # if needed
go get golang.org/x/exp/constraints
go run main.go
```

## Expected Output

The program demonstrates:
- Min/Max with different numeric types and strings
- Map transforming slices from one type to another
- Filter selecting elements based on predicates
- Reduce aggregating values
- Contains checking membership

## Next Steps

See [02-generic-types](../02-generic-types/) to learn about creating generic data structures.
