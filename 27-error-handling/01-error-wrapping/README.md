# Error Wrapping

Error wrapping allows you to add context to errors as they propagate up the call stack while preserving the original error. This makes debugging much easier because you can see the full chain of what went wrong.

## Key Concepts

### The `%w` Verb

Go 1.13 introduced the `%w` verb for `fmt.Errorf`, which wraps an error while adding context:

```go
return fmt.Errorf("additional context: %w", originalErr)
```

### Why Wrap Errors?

1. **Add Context**: Each layer can add information about what was being attempted
2. **Preserve Original**: The original error is still available for inspection
3. **Better Debugging**: The full error chain helps trace issues back to their source

## Example Breakdown

### Basic Wrapping

```go
_, err := os.ReadFile(filename)
if err != nil {
    return fmt.Errorf("failed to process file %q: %w", filename, err)
}
```

This wraps the original file system error with context about which file was being processed.

### Multiple Layers

Errors can be wrapped multiple times:

```go
// Layer 3: Database operation
err := fmt.Errorf("user not found")

// Layer 2: Add database context
err = fmt.Errorf("failed to load user %d from database: %w", userID, err)

// Layer 1: Add business logic context
err = fmt.Errorf("failed to process user data: %w", err)
```

Each layer adds more context about what was happening when the error occurred.

## Running the Example

```bash
go run main.go
```

Expected output shows:
- A wrapped file system error with context
- Multiple layers of error wrapping
- How errors propagate up the call stack

## Best Practices

### Do:
- Use `%w` when you want to preserve the error for later inspection
- Add meaningful context at each layer
- Keep error messages lowercase and concise
- Include relevant values (IDs, filenames, etc.) in the context

### Don't:
- Use `%v` if you need to inspect the error later (use `%w` instead)
- Add redundant information already in the wrapped error
- Wrap errors that don't need additional context
- End error messages with punctuation

## Key Takeaways

- Error wrapping builds an error chain
- Each layer adds context without losing the original error
- The `%w` verb is essential for creating inspectable error chains
- Good error wrapping makes debugging significantly easier

## Next Steps

See [02-error-inspection](../02-error-inspection/) to learn how to inspect wrapped errors using `errors.Is` and `errors.As`.
