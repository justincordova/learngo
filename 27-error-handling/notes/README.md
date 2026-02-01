# Error Handling Notes

## The Error Interface

Go's built-in `error` interface is simple:

```go
type error interface {
    Error() string
}
```

Any type that implements an `Error()` method satisfies this interface.

## Basic Error Handling Pattern

The idiomatic Go error handling pattern:

```go
result, err := someFunction()
if err != nil {
    // handle error
    return err
}
// use result
```

## Modern Error Handling (Go 1.13+)

### Error Wrapping with fmt.Errorf

Add context to errors using `%w` verb:

```go
if err != nil {
    return fmt.Errorf("failed to process user %d: %w", userID, err)
}
```

This preserves the original error while adding context.

### errors.Is - Check Error Type

Check if an error matches a target:

```go
if errors.Is(err, fs.ErrNotExist) {
    fmt.Println("File does not exist")
}
```

This works even if the error has been wrapped multiple times.

### errors.As - Extract Specific Error Types

Extract specific error types from a chain:

```go
var pathErr *fs.PathError
if errors.As(err, &pathErr) {
    fmt.Printf("Failed to access: %s\n", pathErr.Path)
}
```

### errors.Unwrap - Access Wrapped Errors

Unwrap errors to access the underlying error:

```go
unwrapped := errors.Unwrap(err)
```

## Sentinel Errors

Pre-defined errors for common cases:

```go
var ErrNotFound = errors.New("not found")
var ErrInvalidInput = errors.New("invalid input")

// Usage
if errors.Is(err, ErrNotFound) {
    // handle not found case
}
```

## Custom Error Types

Create custom error types with additional information:

```go
type ValidationError struct {
    Field string
    Value interface{}
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on field %s: %s", e.Field, e.Message)
}

// Usage
var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Printf("Invalid value for %s: %v\n", valErr.Field, valErr.Value)
}
```

## Error Chains

Wrapped errors form chains:

```
Original Error
    ↓
Wrapped with context "database query failed"
    ↓
Wrapped with context "failed to get user"
```

Use `errors.Is` and `errors.As` to inspect any error in the chain.

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

## Error Handling vs Exceptions

Unlike many languages that use exceptions, Go uses explicit error returns. This makes error handling:
- Visible in the code
- Part of the function signature
- Easier to track and manage
- More explicit about what can go wrong

## Resources

- [Go Blog: Error Handling and Go](https://go.dev/blog/error-handling-and-go)
- [Go Blog: Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [Effective Go: Errors](https://go.dev/doc/effective_go#errors)
- [errors package documentation](https://pkg.go.dev/errors)
