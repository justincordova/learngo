# Error Handling Notes

## Why Error Handling is Needed

Things can go wrong in programs:
- Files might not exist
- Network connections might fail
- User input might be invalid
- System resources might be unavailable

Error handling allows you to:
- Control execution flow based on success/failure
- Provide meaningful feedback to users
- Prevent crashes
- Handle exceptional situations gracefully

## The nil Value

`nil` is the zero value for pointers and pointer-based types.

- Represents an uninitialized or absent value
- Used to indicate "no error" in error handling
- NOT equal to empty string: `"" == nil` is false (and a type error)

```go
var err error  // err is nil (no error)
```

## Error Values

An error value stores error details:

- Returned by functions that can fail
- Type: `error` (an interface type)
- Stored in regular variables (not global)
- Contains information about what went wrong

```go
result, err := someFunction()
```

There are no global error variables or constants in Go - errors are just values.

## How Go Handles Errors

Go uses explicit error checking with `if` statements and `nil` comparison:

```go
result, err := someFunction()
if err != nil {
    // Handle the error
}
// Use result (error was nil)
```

NOT like Java/C#:
- No `throw`/`catch` statements
- No exceptions
- Errors are returned as values
- Explicit and visible

## When to Handle Errors

Handle errors immediately after calling a function that returns an error value:

```go
d, err := time.ParseDuration("1h10s")
if err != nil {  // Check right away
    // Handle error
}
// Use d
```

Not before calling:
- Can't check before the function runs

Not after main ends:
- Too late to handle

## Identifying Functions That Return Errors

Check the function signature. Handle errors for functions that return an `error` value:

```go
func Read() error          // Returns error - handle it
func Write() error         // Returns error - handle it
func String() string       // No error - no handling needed
func Reset()               // No error - no handling needed
```

Handle errors for: `Read` and `Write`

## Error Value Meaning

### nil Error = Success

```go
d, err := time.ParseDuration("1h10s")
if err != nil {
    // This won't run - parsing succeeded
}
// err is nil, so parsing was successful
```

### non-nil Error = Failure

```go
d, err := time.ParseDuration("invalid")
if err != nil {
    // This will run - parsing failed
    // err contains error details
}
```

Note: In some cases (like `io.EOF`), a non-nil error may indicate a special condition rather than failure, but this is an advanced topic.

## Correct Error Handling Pattern

### Incorrect Example

```go
d, err := time.ParseDuration("1h10s")
if err != nil {
    fmt.Println(d)  // WRONG - using d when there's an error
}
```

Problems:
- Prints duration even when parsing failed
- Should only use `d` when `err == nil`

### Correct Example

```go
d, err := time.ParseDuration("1h10s")
if err != nil {
    // Error occurred - don't use d
    fmt.Println("Parsing error:", err)
    return  // Exit or handle error
}
// No error - safe to use d
fmt.Println(d)
```

Key points:
- Check `if err != nil` to detect errors
- If error exists, handle it (print message, return, etc.)
- Only use the result value when `err == nil`
- Use `return` to exit after handling error

## Error Handling Steps

1. Call function that returns an error
2. Immediately check if `err != nil`
3. If error exists:
   - Print or log the error
   - Return from function or handle appropriately
   - Do NOT use other returned values
4. If no error (err == nil):
   - Safe to use other returned values
   - Continue with normal execution

## Complete Pattern

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    d, err := time.ParseDuration("1h10s")
    if err != nil {
        fmt.Println("Parsing error:", err)
        return
    }
    fmt.Println(d)  // Only prints if no error
}
```

## Key Points

1. Error handling prevents crashes and handles failures
2. `nil` means uninitialized/absent value, indicates "no error"
3. Error values store error details
4. Go uses explicit `if err != nil` checking (no try/catch)
5. Check errors immediately after function calls
6. `nil` error = success, non-nil error = failure
7. Don't use result values when there's an error
8. Use `return` or appropriate action after handling error
9. Errors are values, not exceptions
10. Handle errors from any function returning `error` type
