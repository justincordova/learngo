# Error Inspection

Error inspection allows you to examine wrapped errors and make decisions based on their type or value. Go provides two primary functions for this: `errors.Is` and `errors.As`.

## Key Concepts

### `errors.Is`

Checks if an error (or any error in its chain) matches a target error:

```go
if errors.Is(err, ErrNotFound) {
    // Handle not found case
}
```

### `errors.As`

Extracts a specific error type from the error chain:

```go
var pathErr *fs.PathError
if errors.As(err, &pathErr) {
    fmt.Printf("Failed to access: %s\n", pathErr.Path)
}
```

## Why Use These Functions?

### Before Go 1.13 (Don't Do This)

```go
// Bad: String comparison breaks with wrapped errors
if err.Error() == "not found" { ... }

// Bad: Type assertion only checks the outermost error
if _, ok := err.(*fs.PathError); ok { ... }
```

### With Go 1.13+ (Do This)

```go
// Good: Works with wrapped errors
if errors.Is(err, ErrNotFound) { ... }

// Good: Searches the entire error chain
var pathErr *fs.PathError
if errors.As(err, &pathErr) { ... }
```

## Example Breakdown

### Using `errors.Is`

```go
err := fmt.Errorf("operation failed: %w", ErrNotFound)

// errors.Is unwraps the error chain to find ErrNotFound
if errors.Is(err, ErrNotFound) {
    // This will match even though err is wrapped
    fmt.Println("Resource not found")
}
```

`errors.Is` is perfect for checking against sentinel errors (predefined error values).

### Using `errors.As`

```go
_, err := os.ReadFile("config.json")
err = fmt.Errorf("failed to load config: %w", err)

// errors.As finds *fs.PathError anywhere in the chain
var pathErr *fs.PathError
if errors.As(err, &pathErr) {
    fmt.Printf("File: %s\n", pathErr.Path)      // config.json
    fmt.Printf("Operation: %s\n", pathErr.Op)   // open
}
```

`errors.As` extracts concrete error types, giving you access to their fields and methods.

### Combining Both Methods

```go
// First check for sentinel errors
if errors.Is(err, ErrNotFound) {
    return handleNotFound()
}

// Then extract detailed error information
var pathErr *fs.PathError
if errors.As(err, &pathErr) {
    return handlePathError(pathErr)
}
```

## Running the Example

```bash
go run main.go
```

Expected output demonstrates:
- Checking sentinel errors with `errors.Is`
- Extracting error details with `errors.As`
- How both functions work through wrapped error chains

## Best Practices

### Do:
- Use `errors.Is` to check for specific sentinel errors
- Use `errors.As` when you need to access error fields or methods
- Check the most specific errors first, then more general ones
- Pass a pointer to `errors.As` (e.g., `&pathErr`)

### Don't:
- Use string comparison to check error types
- Use type assertions directly on errors (they won't unwrap)
- Assume the error format won't change
- Forget that `errors.As` modifies its second argument

## How It Works

Both functions traverse the error chain by repeatedly calling `Unwrap()`:

```
Your error -> Wrapped error -> Wrapped error -> Original error
    ↓              ↓                 ↓              ↓
errors.Is and errors.As check each level until they find a match
```

## Key Takeaways

- `errors.Is` checks if any error in the chain matches a value
- `errors.As` finds the first error of a specific type in the chain
- Both functions handle wrapped errors automatically
- They're essential for robust error handling in modern Go

## Next Steps

See [03-custom-errors](../03-custom-errors/) to learn how to create your own error types that work seamlessly with `errors.Is` and `errors.As`.
