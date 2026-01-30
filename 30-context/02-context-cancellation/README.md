# Context Cancellation

This example demonstrates how to use `context.WithCancel()` to gracefully stop goroutines, clean up resources, and handle cancellation signals in concurrent Go programs.

## Concepts Covered

### 1. context.WithCancel()

Creates a cancellable context:

```go
ctx, cancel := context.WithCancel(parent)
defer cancel() // ALWAYS defer cancel to prevent leaks
```

**Returns:**
- `ctx`: A new context that can be cancelled
- `cancel`: A function to cancel the context

### 2. Listening for Cancellation

Check for cancellation in goroutines:

```go
select {
case <-ctx.Done():
    // Context was cancelled
    return
case data := <-dataChan:
    // Process data
}
```

### 3. Cancel Function

The cancel function:
- Closes the `ctx.Done()` channel
- Sets `ctx.Err()` to `context.Canceled`
- Can be called multiple times safely (only first call has effect)
- Must be called to free resources (use `defer`)

### 4. Cancellation Propagation

When a parent context is cancelled, all child contexts are automatically cancelled:

```
Parent Context (cancelled)
    |
    +-- Child 1 (automatically cancelled)
    |
    +-- Child 2 (automatically cancelled)
```

## Key Patterns

### Pattern 1: Basic Worker Cancellation

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            // Clean up and exit
            return
        default:
            // Do work
        }
    }
}
```

### Pattern 2: Worker Pool Shutdown

```go
func workerPool(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            return // Shutdown
        case job := <-jobs:
            process(job)
        }
    }
}
```

### Pattern 3: Always Defer Cancel

```go
func doWork() error {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // Prevents resource leak

    // Use ctx...
    return nil
}
```

## Running the Example

```bash
go run main.go
```

## Expected Output

The program demonstrates:
1. Basic cancellation of a single goroutine
2. Cancelling multiple goroutines simultaneously
3. Graceful worker pool shutdown
4. Cancellation propagation from parent to children
5. Proper cleanup with defer
6. Checking cancellation reason with ctx.Err()

## Use Cases

### 1. Stopping Background Workers

```go
ctx, cancel := context.WithCancel(context.Background())

go backgroundWorker(ctx)

// Later, when shutting down
cancel()
```

### 2. HTTP Request Cancellation

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context() // Cancelled if client disconnects

    result, err := longRunningOperation(ctx)
    if ctx.Err() == context.Canceled {
        // Client disconnected, stop processing
        return
    }
}
```

### 3. Graceful Server Shutdown

```go
func server() {
    ctx, cancel := context.WithCancel(context.Background())

    // Start workers
    for i := 0; i < 10; i++ {
        go worker(ctx, i)
    }

    // Wait for shutdown signal
    <-shutdownSignal

    // Cancel all workers
    cancel()
}
```

### 4. Coordinating Multiple Operations

```go
func processFiles(files []string) error {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    errChan := make(chan error, len(files))

    for _, file := range files {
        go func(f string) {
            errChan <- processFile(ctx, f)
        }(file)
    }

    for range files {
        if err := <-errChan; err != nil {
            cancel() // Cancel remaining operations
            return err
        }
    }

    return nil
}
```

## Common Mistakes to Avoid

### Mistake 1: Not Deferring Cancel

```go
// ❌ WRONG: Resource leak if function returns early
ctx, cancel := context.WithCancel(parent)
// ... use ctx ...
cancel() // Might not be called if function returns early

// ✅ CORRECT: Always defer
ctx, cancel := context.WithCancel(parent)
defer cancel() // Always called, even on early return or panic
```

### Mistake 2: Ignoring ctx.Done()

```go
// ❌ WRONG: Goroutine never stops
func worker(ctx context.Context) {
    for {
        // No cancellation check!
        doWork()
    }
}

// ✅ CORRECT: Check ctx.Done()
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            doWork()
        }
    }
}
```

### Mistake 3: Blocking Without Cancellation Check

```go
// ❌ WRONG: Can't be cancelled while blocked
func worker(ctx context.Context, ch <-chan Data) {
    for data := range ch { // Blocks here
        process(data)
    }
}

// ✅ CORRECT: Use select with ctx.Done()
func worker(ctx context.Context, ch <-chan Data) {
    for {
        select {
        case <-ctx.Done():
            return
        case data := <-ch:
            process(data)
        }
    }
}
```

### Mistake 4: Not Checking ctx.Err()

```go
// ❌ WRONG: Don't know why context ended
select {
case <-ctx.Done():
    return errors.New("operation failed")
}

// ✅ CORRECT: Check ctx.Err() for the reason
select {
case <-ctx.Done():
    return ctx.Err() // Returns context.Canceled or context.DeadlineExceeded
}
```

## Best Practices

1. **Always defer cancel()**: Prevents resource leaks
2. **Check ctx.Done() regularly**: In loops and before expensive operations
3. **Propagate cancellation**: Pass context to called functions
4. **Use select for responsiveness**: Check cancellation while waiting
5. **Check ctx.Err()**: To distinguish between cancellation reasons
6. **Don't ignore Done()**: Always handle cancellation signals

## Context Errors

After cancellation, `ctx.Err()` returns:

- `context.Canceled`: Context was explicitly cancelled via cancel()
- `context.DeadlineExceeded`: Context deadline/timeout was exceeded
- `nil`: Context is not cancelled

## Next Steps

Now that you understand cancellation, move on to:
- **Context Timeout**: Automatic cancellation with time limits
- **Context Values**: Passing request-scoped data
- **Context in Production**: Real-world patterns and anti-patterns

## Resources

- [Go Blog: Context](https://go.dev/blog/context)
- [context package documentation](https://pkg.go.dev/context)
- [Go Concurrency Patterns: Context](https://go.dev/blog/pipelines)
