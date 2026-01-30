# Type Constraints

Type constraints specify which types can be used as type arguments for generic functions and types. They define the operations that are allowed on generic type parameters.

## Built-in Constraints

### `any`

No constraints - accepts any type (equivalent to `interface{}`):

```go
func Print[T any](value T) {
    fmt.Println(value)
}
```

### `comparable`

Types that support `==` and `!=` operators:

```go
func Equal[T comparable](a, b T) bool {
    return a == b
}
```

Includes: integers, floats, strings, booleans, pointers, channels, interfaces, arrays/structs of comparable types.

## Standard Library Constraints

The `golang.org/x/exp/constraints` package provides useful constraints:

### `constraints.Ordered`

Types that support `<`, `>`, `<=`, `>=`:

```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

Includes: integers, floats, strings.

### `constraints.Integer`

All integer types (signed and unsigned).

### `constraints.Float`

All floating-point types.

### `constraints.Signed`

All signed integer types.

### `constraints.Unsigned`

All unsigned integer types.

## Custom Constraints

### Union Type Constraints

Use `|` to specify multiple types:

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

### Method-Based Constraints

Require specific methods:

```go
type Shape interface {
    Area() float64
}

func GetArea[T Shape](s T) float64 {
    return s.Area()
}
```

### The `~` Operator

The `~` means "underlying type" - includes user-defined types:

```go
type Integer interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Works with custom types like:
type MyInt int
```

Without `~`, only the exact type `int` would be allowed, not `MyInt`.

## Constraint Composition

Combine multiple constraints:

```go
type Stringer interface {
    String() string
}

type PrintableOrdered interface {
    constraints.Ordered
    Stringer
}
```

## Key Points

- Constraints define what operations are allowed on type parameters
- Use the most specific constraint that makes sense
- `any` is the least restrictive, `comparable` allows equality checks
- Custom constraints can combine type unions and method requirements
- The `~` operator allows user-defined types based on built-in types
- Method-based constraints work like regular interfaces

## Running This Example

```bash
cd 03-type-constraints
go mod init example  # if needed
go get golang.org/x/exp/constraints
go run main.go
```

## Expected Output

The program demonstrates:
- Built-in constraints (`any`, `comparable`)
- Standard library constraints (`constraints.Ordered`)
- Custom numeric constraints
- Method-based constraints
- Union type constraints
- Constraint composition

## Common Patterns

### Numeric Operations

```go
type Number interface {
    int | float64
}

func Add[T Number](a, b T) T {
    return a + b
}
```

### Comparison Operations

```go
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

### Behavior-Based

```go
type Serializer interface {
    Serialize() ([]byte, error)
}

func Save[T Serializer](item T) error {
    data, err := item.Serialize()
    // ...
}
```

## Next Steps

Complete the exercises in the [exercises](../exercises/) directory to practice using type constraints.
