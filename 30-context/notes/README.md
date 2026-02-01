# Context Package Notes

## Introduction

The `context` package is essential for managing request-scoped values, cancellation signals, and deadlines across API boundaries and goroutines.

## The Context Interface

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

## Creating Contexts

### context.Background()

The top-level context for main, init, and tests:

```go
ctx := context.Background()
```

Use this as the root of your context tree.

### context.TODO()

Use when you're unsure which context to use or during refactoring:

```go
ctx := context.TODO()
```

## Context Cancellation

### WithCancel

Manually cancel a context:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()  // Always call cancel to avoid resource leaks

go func() {
    <-ctx.Done()
    fmt.Println("Cancelled:", ctx.Err())
}()

// Later...
cancel()  // Triggers cancellation
```

### Checking for Cancellation

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Cancelled:", ctx.Err())
            return
        default:
            // Do work
        }
    }
}
```

## Context Timeouts

### WithTimeout

Automatically cancel after a duration:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

result, err := doWork(ctx)
if err == context.DeadlineExceeded {
    fmt.Println("Operation timed out")
}
```

### WithDeadline

Cancel at a specific time:

```go
deadline := time.Now().Add(10 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

## Context Values

### Storing Values

Pass request-scoped data through the call chain:

```go
type key string

const requestIDKey key = "requestID"

ctx := context.WithValue(context.Background(), requestIDKey, "abc123")
```

### Retrieving Values

```go
func processRequest(ctx context.Context) {
    if reqID := ctx.Value(requestIDKey); reqID != nil {
        fmt.Println("Request ID:", reqID)
    }
}
```

## Common Patterns

### HTTP Server Pattern

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context() // Get request context

    result, err := doWork(ctx)
    if err != nil {
        if ctx.Err() == context.Canceled {
            // Client disconnected
            return
        }
        // Handle other errors
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(result)
}
```

### Database Query with Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

rows, err := db.QueryContext(ctx, "SELECT * FROM users WHERE active = ?", true)
if err == context.DeadlineExceeded {
    log.Println("Query timed out")
    return
}
defer rows.Close()
```

### Worker Pool with Graceful Shutdown

```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            log.Println("Worker shutting down:", ctx.Err())
            return
        case job := <-jobs:
            if err := processJob(ctx, job); err != nil {
                log.Printf("Job failed: %v", err)
            }
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    jobs := make(chan Job)

    // Start workers
    for i := 0; i < 5; i++ {
        go worker(ctx, jobs)
    }

    // On shutdown signal
    cancel() // All workers will gracefully shut down
}
```

### Propagating Context Through API Calls

```go
func fetchUser(ctx context.Context, userID int) (*User, error) {
    req, err := http.NewRequestWithContext(ctx, "GET",
        fmt.Sprintf("https://api.example.com/users/%d", userID), nil)
    if err != nil {
        return nil, err
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // If ctx is cancelled, the HTTP request is automatically cancelled
    // ...
}
```

## Context Tree

Contexts form a tree structure:

```
context.Background()
    |
    +-- WithCancel()
    |       |
    |       +-- WithTimeout()
    |
    +-- WithValue()
            |
            +-- WithDeadline()
```

When a parent context is cancelled, all descendants are automatically cancelled.

## Best Practices

### Do:
- Pass context as the first parameter to functions: `func DoSomething(ctx context.Context, ...)`
- Use `context.Background()` at the top level (main, init, tests)
- Use `context.TODO()` when refactoring and unsure about context
- Always call the cancel function (use `defer cancel()`)
- Check `ctx.Done()` in long-running operations
- Use `ctx.Err()` to understand why context was cancelled

### Don't:
- Store contexts in structs (except for bridging goroutines)
- Pass nil context (use `context.TODO()` instead)
- Use context for optional parameters (use proper function arguments)
- Put important data in context values (use them sparingly)
- Ignore the cancel function (causes resource leaks)
- Use context values for function parameters

## Context Values: When to Use

Context values should be used sparingly for request-scoped data that crosses API boundaries:

**Good use cases:**
- Request IDs for tracing
- Authentication tokens
- User information for request context

**Bad use cases:**
- Optional function parameters
- Function configuration
- Data that affects program logic

```go
// BAD - using context for optional parameters
func ProcessData(ctx context.Context, data []byte) {
    if timeout := ctx.Value("timeout"); timeout != nil {
        // Don't do this
    }
}

// GOOD - explicit parameters
func ProcessData(ctx context.Context, data []byte, timeout time.Duration) {
    // Clear function signature
}
```

## Common Errors

### Resource Leak

```go
// BAD - context leak
ctx, cancel := context.WithCancel(context.Background())
// Forgot to call cancel()

// GOOD - always defer cancel
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

### Not Checking Done Channel

```go
// BAD - ignoring cancellation
func worker(ctx context.Context) {
    for {
        doWork()  // Never checks if cancelled
    }
}

// GOOD - checking for cancellation
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

## Resources

- [Go Blog: Context](https://go.dev/blog/context)
- [context package documentation](https://pkg.go.dev/context)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
