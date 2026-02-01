# Printf Formatting Notes

## Basic Syntax

`fmt.Printf` requires a format string with verbs that match the number and types of arguments:

```go
fmt.Printf("Hi %s", "hello")  // Correct: 1 verb, 1 argument
```

Incorrect examples:
- `fmt.Printf("Hi %s")` - verb without argument
- `fmt.Printf("Hi %s", "how", "are you")` - 1 verb, 2 arguments (mismatch)
- `fmt.Printf("Hi %s", true)` - wrong type (bool for string verb)

## Format Verbs by Type

### String: %s
```go
fmt.Printf("Hi %s %s", "there", ".")  // Prints: Hi there .
```

Both verbs and arguments must be strings. Mixing types causes errors:
- `fmt.Printf("Hi %s %s", "5", true)` - ERROR: second argument is bool
- `fmt.Printf("Hi %s %s", "true", false)` - ERROR: second argument is bool

### Integer: %d
```go
fmt.Printf("Number: %d", 42)  // Prints: Number: 42
```

### Float: %f
```go
fmt.Printf("Pi: %f", 3.14)  // Prints: Pi: 3.140000
```

### Boolean: %t
```go
fmt.Printf("Valid: %t", true)  // Prints: Valid: true
```

### Universal: %v
```go
fmt.Printf("Value: %v", anything)  // Works with any type
```

%v prints the value in a default format, regardless of type.

### Type Information: %T
```go
fmt.Printf("%T", 3.14)   // Prints: float64
fmt.Printf("%T", true)   // Prints: bool
fmt.Printf("%T", 42)     // Prints: int
fmt.Printf("%T", "hi")   // Prints: string
```

%T prints the type of the value, not the value itself.

## Escape Sequences

### Newline: \n
```go
fmt.Printf("Line 1\nLine 2")
// Prints:
// Line 1
// Line 2
```

### Escaped Backslash: \\
```go
fmt.Printf("\\n")  // Prints: \n (literal characters, not newline)
```

### Backslash in Paths
```go
fmt.Printf("c:\\secret\\directory")  // Prints: c:\secret\directory
```

Each `\\` becomes a single `\` in the output.

### Escaped Quotes: \"
```go
fmt.Printf("\"heisenberg\"")  // Prints: "heisenberg"
```

Use `\"` to include literal quote marks in the string.

## Summary Table

| Verb | Type    | Example                           |
|------|---------|-----------------------------------|
| %s   | string  | `fmt.Printf("%s", "hello")`      |
| %d   | int     | `fmt.Printf("%d", 42)`           |
| %f   | float   | `fmt.Printf("%f", 3.14)`         |
| %t   | bool    | `fmt.Printf("%t", true)`         |
| %v   | any     | `fmt.Printf("%v", anything)`     |
| %T   | type    | `fmt.Printf("%T", value)`        |

## Key Points

1. Number of verbs must match number of arguments
2. Verb types must match argument types
3. %v works with any type
4. %T prints the type, not the value
5. Use backslash to escape special characters
