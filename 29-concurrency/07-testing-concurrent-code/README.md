# Testing Concurrent Code (Go 1.25+)

This example demonstrates how to test concurrent code using the `testing/synctest` package introduced in Go 1.25.

## What is synctest?

The `testing/synctest` package provides tools for testing concurrent code in a deterministic, fast, and isolated manner:

- **Deterministic**: Tests produce consistent results every run
- **Fast**: No real time passes - time is simulated
- **Isolated**: Each test has its own independent time context

## Key Benefits

### 1. No Real Waiting
Traditional concurrent tests require actual time to pass:
```go
time.Sleep(100 * time.Millisecond) // Actually waits 100ms
```

With synctest, time is fake and instant:
```go
synctest.Run(func() {
    time.Sleep(100 * time.Millisecond) // Returns immediately!
})
```

### 2. Deterministic Results
Real-time tests can be flaky due to:
- System load variations
- Scheduler unpredictability
- Race conditions

Synctest eliminates these issues with controlled time progression.

### 3. Easier Debugging
When tests fail, you can reason about the exact sequence of events without timing-related race conditions.

## Core APIs

### `synctest.Test(t *testing.T, func(*testing.T))`
Creates an isolated test environment with fake time:
```go
synctest.Test(t, func(t *testing.T) {
    // All time operations here use fake time
    time.Sleep(1 * time.Hour) // Instant!
})
```

### `synctest.Wait()`
Advances fake time until all goroutines are idle or blocked:
```go
synctest.Test(t, func(t *testing.T) {
    go func() {
        time.Sleep(100 * time.Millisecond)
        // Do work
    }()

    synctest.Wait() // Returns when goroutine completes
})
```

## Examples in This Directory

### TimeBasedWorker
A worker that processes tasks on a timer:
- Without synctest: Must wait real time (slow, flaky)
- With synctest: Instant, deterministic testing

### RateLimiter
Limits operations per time period:
- Without synctest: Tests take real time
- With synctest: Tests complete instantly

### ConcurrentCache
Thread-safe cache with TTL expiration:
- Without synctest: Must wait for expiration (slow)
- With synctest: Instant expiration testing

## Running the Examples

### Run the main program:
```bash
go run main.go
```

### Run tests (traditional + synctest):
```bash
go test -v
```

### Run only synctest tests:
```bash
go test -v -run WithSynctest
```

### Run benchmarks to see speed difference:
```bash
go test -bench=.
```

## Important Notes

**Requires Go 1.25 or later**
```bash
go version  # Should show 1.25.0 or higher
```

**Not all code works with synctest**
- External time sources (system clocks, databases)
- Network I/O with real timeouts
- Third-party libraries that don't use `time` package

**Best for**
- Internal business logic with time dependencies
- Worker pools with intervals
- Rate limiters
- Caches with TTL
- Timeout/retry logic

## When to Use Synctest

✅ **Use synctest for:**
- Testing time-based business logic
- Testing concurrent workers and pools
- Testing rate limiting and throttling
- Testing cache expiration
- Testing retry mechanisms

❌ **Don't use synctest for:**
- Integration tests with real external services
- Tests that measure actual performance
- Tests with true wall-clock requirements

## Pattern: Traditional vs Synctest

### Traditional Test (Slow)
```go
func TestWorker_Traditional(t *testing.T) {
    worker := NewWorker(100 * time.Millisecond)
    worker.Start()
    defer worker.Stop()

    worker.Schedule("task")
    time.Sleep(150 * time.Millisecond) // Actual wait!

    if worker.ProcessedCount() != 1 {
        t.Error("Task not processed")
    }
}
```

### Synctest Test (Fast)
```go
func TestWorker_WithSynctest(t *testing.T) {
    synctest.Test(t, func(t *testing.T) {
        worker := NewWorker(100 * time.Millisecond)
        worker.Start()
        defer worker.Stop()

        worker.Schedule("task")
        synctest.Wait() // Instant!

        if worker.ProcessedCount() != 1 {
            t.Error("Task not processed")
        }
    })
}
```

## Learn More

- [Go 1.25 Release Notes - synctest](https://go.dev/doc/go1.25#testing/synctest)
- [testing/synctest Package Documentation](https://pkg.go.dev/testing/synctest)
- Run the tests to see the difference: `go test -v`

## Key Takeaways

1. **Synctest makes concurrent tests fast and deterministic**
2. **Use `synctest.Run()` to create isolated time contexts**
3. **Use `synctest.Wait()` to advance fake time**
4. **Tests that took seconds now take milliseconds**
5. **Only works with code using the `time` package**
6. **Requires Go 1.25 or later**
