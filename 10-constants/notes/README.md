# Constants Notes

## Magic Values vs Named Constants

### Magic Value

An unnamed constant value appearing directly in source code:

```go
area := 3.14 * radius * radius  // 3.14 is a magic value
```

Magic values reduce readability because their meaning isn't clear.

### Named Constant

A constant value declared with a name:

```go
const Pi = 3.14
area := Pi * radius * radius  // Clear what 3.14 represents
```

Named constants improve readability and maintainability.

## Declaring Constants

### Syntax

```go
const version int = 3       // With explicit type
const version = 3           // Type inferred
```

Incorrect:
- `Const version int = 3` - Wrong case (must be lowercase `const`)
- `const version int := 3` - Can't use `:=` with constants
- `const version int` - Constants must be initialized

## Constant Initialization Rules

### Must Use Constant Expressions

Constants can only be initialized with:
- Literal values
- Other constants
- Certain built-in functions (like `len`) with constant arguments

```go
// Valid:
const message = "pick me!"
const length = len(message)  // OK - len of constant string

// Invalid:
s := "pick me"
const length = len(s)  // ERROR - s is not constant

const length = utf8.RuneCountInString("pick me")  // ERROR - can't call functions
```

Functions allowed in constant initialization:
- `len()` - if argument is constant
- `cap()` - if argument is constant
- Built-in functions that compile-time evaluate

## Typeless (Untyped) Constants

Constants without explicit type are "typeless" and can be used flexibly:

```go
const speed = 100      // Typeless int constant
porsche := speed * 3   // porsche is int (inferred)
```

Key points:
- `speed` has no specific type (it's typeless)
- `porsche` is a variable with type `int`
- Variables cannot be typeless

### Flexibility of Typeless Constants

```go
const speed = 100       // Typeless
var mph int64 = speed   // OK - speed adapts to int64
var kph float64 = speed // OK - speed adapts to float64
```

Typed constants are less flexible:

```go
const speed int8 = 100
var mph int64 = speed   // ERROR - type mismatch
```

## Constants Must Be Initialized

```go
// Invalid:
const spell string      // ERROR - not initialized
spell = "Abracadabra"

// Valid:
const spell = "Abracadabra"  // OK - type inferred from value
```

Constants are compile-time values and must have their value at declaration.

## Type Compatibility with Constants

Typeless constants avoid type mismatch errors:

```go
// Problem with typed constant:
const total int8 = 10
x := 5                   // x is int
fmt.Print(total * x)     // ERROR - can't multiply int8 and int

// Solution 1: Make constant typeless
const total = 10         // Typeless
x := 5
fmt.Print(total * x)     // OK - total adapts to int

// Solution 2: Type conversion (not ideal)
const total int64 = 10
x := 5
fmt.Print(int64(total) * x)  // Still ERROR - x is int, not int64
```

## Constants with iota

`iota` is a special identifier that generates sequential values:

```go
const (
    Yes = (iota * 5) + 2  // 0*5 + 2 = 2
    No                     // 1*5 + 2 = 7 (expression repeated)
    Both                   // 2*5 + 2 = 12
)
// Result: Yes=2, No=7, Both=12
```

How iota works:
- Starts at 0 for the first constant in a `const` block
- Increments by 1 for each subsequent constant
- The expression from the first line is repeated for following lines
- Resets to 0 in each new `const` block

Common iota patterns:

```go
const (
    A = iota  // 0
    B         // 1
    C         // 2
)

const (
    KB = 1 << (10 * iota)  // 1 << 0 = 1
    MB                      // 1 << 10 = 1024
    GB                      // 1 << 20 = 1048576
)
```

## Key Points

1. Named constants improve readability over magic values
2. Constants must be initialized at declaration
3. Typeless constants are more flexible than typed ones
4. Constants can use `len()` with constant arguments
5. Most function calls aren't allowed in constant initialization
6. Variables cannot be typeless, only constants can
7. Use `const`, not `Const` or `CONST`
8. iota generates sequential values starting from 0
9. Prefer typeless constants unless type safety is needed

## Best Practices

```go
// Good - typeless, flexible:
const Pi = 3.14159
const MaxRetries = 3

// Less flexible - typed when needed for type safety:
type Status int
const (
    Pending Status = iota
    Approved
    Rejected
)
```
