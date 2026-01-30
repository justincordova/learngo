# Custom Error Types

Custom error types allow you to attach additional context and behavior to errors beyond simple error messages. They work seamlessly with Go's error inspection functions.

## Key Concepts

### Creating Custom Error Types

Any type that implements the `error` interface can be an error:

```go
type error interface {
    Error() string
}
```

### Basic Custom Error

```go
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %q: %s", e.Field, e.Message)
}
```

### Custom Error with Unwrap

To support error wrapping, implement `Unwrap()`:

```go
type DatabaseError struct {
    Operation string
    Table     string
    Err       error
}

func (e *DatabaseError) Error() string {
    return fmt.Sprintf("database error during %s on %q: %v",
        e.Operation, e.Table, e.Err)
}

func (e *DatabaseError) Unwrap() error {
    return e.Err
}
```

### Custom Error with Is

Override how `errors.Is` compares your error:

```go
type NetworkError struct {
    Host string
    Port int
    Err  error
}

func (e *NetworkError) Is(target error) bool {
    t, ok := target.(*NetworkError)
    if !ok {
        return false
    }
    return e.Host == t.Host && e.Port == t.Port
}
```

## When to Use Custom Errors

### Use Custom Error Types When:
- You need to attach additional context (fields, metadata)
- Callers need to programmatically inspect error details
- You want type-safe error handling
- Different error types require different handling logic

### Use Sentinel Errors When:
- The error has no additional context
- You only need to check if a specific error occurred
- Example: `var ErrNotFound = errors.New("not found")`

## Example Breakdown

### ValidationError

```go
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}
```

Stores which field failed validation, its value, and why. Perfect for API validation where you need to tell users exactly what's wrong.

### DatabaseError with Unwrap

```go
type DatabaseError struct {
    Operation string
    Table     string
    Err       error
    Timestamp time.Time
}

func (e *DatabaseError) Unwrap() error {
    return e.Err
}
```

Wraps database errors with operation context. The `Unwrap()` method allows `errors.Is` and `errors.As` to inspect the wrapped error.

### NetworkError with Is

```go
func (e *NetworkError) Is(target error) bool {
    t, ok := target.(*NetworkError)
    if !ok {
        return false
    }
    return e.Host == t.Host && e.Port == t.Port
}
```

Custom `Is` method allows checking if errors match based on host and port, ignoring other fields like the underlying error or timeout.

## Using Custom Errors

### Creating and Returning

```go
if age < 18 {
    return &ValidationError{
        Field:   "age",
        Value:   age,
        Message: "must be at least 18 years old",
    }
}
```

### Inspecting with errors.As

```go
var validationErr *ValidationError
if errors.As(err, &validationErr) {
    fmt.Printf("Field: %s\n", validationErr.Field)
    fmt.Printf("Value: %v\n", validationErr.Value)
}
```

### Wrapping Custom Errors

```go
if err := validateUser(name, age); err != nil {
    return fmt.Errorf("registration failed: %w", err)
}
```

Even when wrapped, `errors.As` can still find your custom error type.

## Running the Example

```bash
go run main.go
```

Expected output demonstrates:
- Creating custom error types with additional fields
- Wrapping errors in custom types
- Using `errors.As` to extract custom error details
- Custom `Is` methods for specialized comparison

## Best Practices

### Do:
- Use pointer receivers for error methods (`*ValidationError`)
- Implement `Unwrap()` if your error wraps another error
- Export error types that callers need to inspect
- Keep error types focused on a single concern
- Store timestamp or request ID for debugging

### Don't:
- Create custom errors when sentinel errors suffice
- Forget to implement `Error()` method (required)
- Use value receivers for error types
- Export private implementation details in error fields

## Method Summary

| Method | Purpose | Required |
|--------|---------|----------|
| `Error() string` | Provides error message | Yes |
| `Unwrap() error` | Enables error chain traversal | Optional |
| `Is(error) bool` | Customizes equality checking | Optional |
| `As(interface{}) bool` | Customizes type matching | Rare |

## Key Takeaways

- Custom error types store additional context beyond error messages
- Implement `Unwrap()` to support error wrapping
- Use `errors.As` to extract custom error types from chains
- Custom `Is()` methods enable specialized error comparison
- Custom errors work seamlessly with Go's error handling functions

## Next Steps

See [04-error-chains](../04-error-chains/) to learn more about how wrapped errors form chains and advanced error handling patterns.
