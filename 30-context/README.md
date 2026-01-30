# Context Package in Go

The `context` package is essential for managing request-scoped values, cancellation signals, and deadlines across API boundaries and goroutines. It's a fundamental tool for building robust concurrent applications.

## Overview

This section covers the context package and its patterns:

- **Context Basics**: Understanding Background, TODO, and why context exists
- **Context Cancellation**: Using WithCancel to stop goroutines gracefully
- **Context Timeout**: Automatic cancellation with WithTimeout and WithDeadline
- **Context Values**: Passing request-scoped data through the call chain
- **Best Practices**: Common patterns and anti-patterns

## Prerequisites

Before starting this section, you should be comfortable with:

- Basic Go syntax and control flow
- Goroutines and concurrent programming
- Channels (for understanding cancellation patterns)
- Interfaces (context is an interface)
- Error handling in Go
- Select statements

## Key Concepts

### The Context Interface

The `context.Context` interface defines four methods:

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

### Why Context Exists

Context solves several critical problems in concurrent Go programs:

1. **Cancellation Propagation**: Signal cancellation to all goroutines in an operation
2. **Timeouts**: Automatically cancel operations that take too long
3. **Request-Scoped Values**: Pass data like request IDs through call chains
4. **Resource Cleanup**: Ensure goroutines clean up when parent operations complete

### Context Rules

1. **Never store contexts**: Pass context as the first parameter, conventionally named `ctx`
2. **Context in function signatures**: `func DoSomething(ctx context.Context, ...)`
3. **Don't pass nil**: Use `context.TODO()` if unsure which context to use
4. **Context is immutable**: Creating a child context doesn't modify the parent

## Section Contents

1. **[Context Basics](01-context-basics/)** - Learn Background(), TODO(), and basic context concepts

2. **[Context Cancellation](02-context-cancellation/)** - Use WithCancel() to stop goroutines gracefully

3. **[Context Timeout](03-context-timeout/)** - Automatic cancellation with WithTimeout() and WithDeadline()

4. **[Context Values](04-context-values/)** - Pass request-scoped data safely

5. **[Exercises](exercises/)** - Practice context patterns in real scenarios

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
    }
}
```

### Worker Pool Pattern

```go
func worker(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case <-ctx.Done():
            return // Graceful shutdown
        case job := <-jobs:
            processJob(ctx, job)
        }
    }
}
```

### Timeout Pattern

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

result, err := database.QueryContext(ctx, query)
if err == context.DeadlineExceeded {
    // Operation timed out
}
```

## Best Practices

### Do:
- Pass context as the first parameter to functions
- Use `context.Background()` at the top level (main, init, tests)
- Use `context.TODO()` when refactoring and unsure about context
- Always call the cancel function returned by With* functions (use defer)
- Check `ctx.Done()` in long-running operations
- Use `ctx.Err()` to understand why context was cancelled

### Don't:
- Store contexts in structs (except for bridging goroutines)
- Pass nil context (use context.TODO() instead)
- Use context for optional parameters (use proper function arguments)
- Put important data in context values (use them sparingly)
- Ignore the cancel function from With* functions (causes resource leaks)

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

When a parent context is cancelled, all its descendants are automatically cancelled.

## Common Use Cases

1. **HTTP Handlers**: Automatically cancelled when client disconnects
2. **Database Queries**: Cancel long-running queries with timeouts
3. **API Calls**: Propagate cancellation to external services
4. **Worker Pools**: Graceful shutdown of background workers
5. **Microservices**: Pass trace IDs and request metadata

## Resources

- [Go Blog: Context](https://go.dev/blog/context)
- [context package documentation](https://pkg.go.dev/context)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)

## Next Steps

After completing this section, you'll be ready to:
- Build robust concurrent applications with proper cancellation
- Handle timeouts and deadlines effectively
- Pass request-scoped data through your application
- Integrate context with standard library packages (http, database/sql, etc.)
- Write production-grade concurrent code
