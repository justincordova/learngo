# Error Handling in Go

Error handling is a fundamental aspect of writing robust Go programs. Unlike many languages that use exceptions, Go uses explicit error returns, making error handling visible and manageable.

## Overview

This section covers Go's error handling patterns and best practices:

- **Error Wrapping**: Adding context to errors as they propagate up the call stack
- **Error Inspection**: Checking error types and extracting information from errors
- **Custom Errors**: Creating your own error types with additional information
- **Error Chains**: Understanding and working with wrapped error chains

## Prerequisites

Before starting this section, you should be comfortable with:

- Basic Go syntax and control flow
- Functions and multiple return values
- Interfaces (particularly the `error` interface)
- Pointers (for custom error types)
- Structs (for building custom error types)

## Key Concepts

### The Error Interface

Go's built-in `error` interface is simple:

```go
type error interface {
    Error() string
}
```

Any type that implements an `Error()` method satisfies this interface.

### Error Handling Pattern

The idiomatic Go error handling pattern:

```go
result, err := someFunction()
if err != nil {
    // handle error
    return err
}
// use result
```

### Modern Error Handling (Go 1.13+)

Go 1.13 introduced powerful error handling features:

- `fmt.Errorf` with `%w` verb for wrapping errors
- `errors.Is` for checking if an error matches a target
- `errors.As` for extracting specific error types from a chain
- `errors.Unwrap` for accessing wrapped errors

## Section Contents

1. **[Error Wrapping](01-error-wrapping/)** - Learn to add context to errors using `fmt.Errorf` and `%w`

2. **[Error Inspection](02-error-inspection/)** - Use `errors.Is` and `errors.As` to examine errors

3. **[Custom Errors](03-custom-errors/)** - Create your own error types with additional data

4. **[Error Chains](04-error-chains/)** - Understand how wrapped errors form chains and how to work with them

5. **[Exercises](exercises/)** - Practice error handling patterns

## Best Practices

### Do:
- Always check errors returned from functions
- Add context when wrapping errors to help with debugging
- Use `errors.Is` and `errors.As` for error inspection
- Create custom error types when you need additional information
- Keep error messages lowercase and avoid punctuation at the end

### Don't:
- Ignore errors (don't use `_` for error returns unless absolutely necessary)
- Panic unless it's a programming error or initialization failure
- Use string matching to check error types
- Create unnecessary error types (use sentinel errors when appropriate)

## Common Patterns

### Sentinel Errors

Pre-defined errors for common cases:

```go
var ErrNotFound = errors.New("not found")
var ErrInvalidInput = errors.New("invalid input")
```

### Error Wrapping

Adding context while preserving the original error:

```go
if err != nil {
    return fmt.Errorf("failed to process user %d: %w", userID, err)
}
```

### Error Type Assertions

Checking for specific error types:

```go
var pathErr *fs.PathError
if errors.As(err, &pathErr) {
    fmt.Printf("Failed to access: %s\n", pathErr.Path)
}
```

## Resources

- [Go Blog: Error Handling and Go](https://go.dev/blog/error-handling-and-go)
- [Go Blog: Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [Effective Go: Errors](https://go.dev/doc/effective_go#errors)
- [errors package documentation](https://pkg.go.dev/errors)

## Next Steps

After completing this section, you'll be ready to:
- Handle errors effectively in real-world applications
- Create maintainable error handling code
- Debug issues using error context
- Build robust error handling into your packages
