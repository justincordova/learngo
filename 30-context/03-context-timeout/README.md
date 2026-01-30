# Context Timeout and Deadline

This example demonstrates how to use `context.WithTimeout()` and `context.WithDeadline()` to create contexts that automatically cancel after a specified time period or at a specific point in time.

## Key Concepts

### context.WithTimeout()

Creates a context that cancels after a relative duration:

```go
// Cancel after 5 seconds from now
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel()
```

### context.WithDeadline()

Creates a context that cancels at an absolute point in time:

```go
// Cancel at 3:00 PM
deadline := time.Date(2024, 1, 1, 15, 0, 0, 0, time.UTC)
ctx, cancel := context.WithDeadline(parentCtx, deadline)
defer cancel()
```

## When to Use

### Use Timeouts For:

- **API calls**: Prevent hanging on slow external services
- **Database queries**: Avoid blocking on slow queries
- **Network operations**: Set maximum wait time for responses
- **User requests**: Ensure responsive user experience
- **Resource allocation**: Prevent indefinite resource holds

### Common Timeout Durations:

- API calls: 5-30 seconds
- Database queries: 1-10 seconds
- HTTP requests: 10-60 seconds
- Microservice calls: 3-10 seconds

## Timeout vs Deadline

| Feature | WithTimeout | WithDeadline |
|---------|-------------|--------------|
| **Time specification** | Relative (duration) | Absolute (time.Time) |
| **Example** | "5 seconds from now" | "at 3:00 PM" |
| **Use when** | Duration-based limits | Specific time constraints |
| **Implementation** | Calls WithDeadline internally | Fundamental operation |

**Both achieve the same result** - just different ways to express when the context should cancel.

## Examples in This Program

1. **Basic timeout** - Simple timeout demonstration
2. **Basic deadline** - Deadline with specific time
3. **Timeout vs Deadline** - Understanding the difference
4. **API call timeout** - Simulating slow API with timeout protection
5. **Database query timeout** - Query operations with time limits
6. **Success within timeout** - Operation completing before timeout
7. **Checking remaining time** - Monitoring time until deadline
8. **Nested timeouts** - Child inherits parent's shorter deadline

## Best Practices

### Always defer cancel()

```go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel() // Releases resources even if timeout not reached
```

**Why?** Even if the operation completes before timeout, you must call `cancel()` to:
- Release associated timer resources
- Cancel any goroutines waiting on the context
- Prevent resource leaks

### Call cancel() early when done

```go
ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
defer cancel()

result, err := fastOperation(ctx)
if err != nil {
    return err
}

cancel() // Release resources immediately, don't wait for defer
return result
```

### Check context before starting work

```go
func doWork(ctx context.Context) error {
    // Check if already cancelled before starting
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }

    // Proceed with work...
}
```

### Use select to respect context

```go
func query(ctx context.Context) (Result, error) {
    resultChan := make(chan Result, 1)

    go func() {
        // Perform actual work
        resultChan <- doActualWork()
    }()

    select {
    case result := <-resultChan:
        return result, nil
    case <-ctx.Done():
        return Result{}, ctx.Err()
    }
}
```

### Check deadline to decide if work is worth starting

```go
deadline, ok := ctx.Deadline()
if ok {
    if time.Until(deadline) < minimumWorkTime {
        return errors.New("not enough time to complete work")
    }
}
```

## Error Handling

When a context times out, `ctx.Err()` returns `context.DeadlineExceeded`:

```go
<-ctx.Done()
err := ctx.Err()

switch err {
case context.DeadlineExceeded:
    // Timeout occurred
    log.Println("operation timed out")
case context.Canceled:
    // Manual cancellation
    log.Println("operation was cancelled")
}
```

## Common Patterns

### HTTP Handler with Timeout

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Add 5s timeout to request context
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    result, err := fetchData(ctx)
    if err == context.DeadlineExceeded {
        http.Error(w, "Request timeout", http.StatusGatewayTimeout)
        return
    }
    // ... handle result
}
```

### Database Query with Timeout

```go
func queryUser(ctx context.Context, db *sql.DB, id int) (*User, error) {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    var user User
    err := db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", id).Scan(&user)
    return &user, err
}
```

### API Call with Timeout

```go
func callExternalAPI(ctx context.Context, url string) ([]byte, error) {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := http.DefaultClient.Do(req)
    // ... handle response
}
```

## Running the Example

```bash
go run main.go
```

The program demonstrates various timeout and deadline scenarios, showing both successful operations and timeouts.

## Key Takeaways

1. **WithTimeout** is for relative durations, **WithDeadline** is for absolute times
2. **Always defer cancel()** to prevent resource leaks
3. **Call cancel() early** when operation completes before timeout
4. Use **select** with `ctx.Done()` to respect cancellation
5. Child contexts **inherit parent deadlines** (shortest deadline wins)
6. Timeouts are essential for **responsive, resilient applications**
7. Different operations need **different timeout values** based on expected duration
