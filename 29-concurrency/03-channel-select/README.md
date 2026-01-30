# Channel Select Statement

The `select` statement is Go's way of working with multiple channels simultaneously. It's similar to a `switch` statement but for channel operations, allowing a goroutine to wait on multiple channel operations at once.

## Key Concepts

### What is Select?

The `select` statement lets a goroutine wait on multiple communication operations:

```go
select {
case msg := <-ch1:
    // Received from ch1
case msg := <-ch2:
    // Received from ch2
case ch3 <- value:
    // Sent to ch3
default:
    // No channel operation ready
}
```

### How Select Works

1. **Evaluates all cases**: Each channel operation is evaluated
2. **Waits for ready operation**: Blocks until one channel is ready
3. **Chooses randomly**: If multiple channels are ready, picks one at random
4. **Executes case**: Runs the corresponding case block
5. **Default case**: If present, executes immediately if no channel is ready

## Example Breakdown

### Basic Select

```go
select {
case msg1 := <-ch1:
    fmt.Println(msg1)
case msg2 := <-ch2:
    fmt.Println(msg2)
}
```

Whichever channel receives a value first will have its case executed.

### Timeout Pattern

One of the most common uses of `select`:

```go
select {
case result := <-ch:
    // Got result
case <-time.After(1 * time.Second):
    // Timeout: took too long
}
```

This prevents indefinite blocking on slow operations.

### Non-Blocking Operations

Use the `default` case for non-blocking channel operations:

```go
select {
case msg := <-ch:
    fmt.Println(msg)
default:
    fmt.Println("No message available")
}
```

If the channel has no data, the default case runs immediately.

### Quit Channel Pattern

Allow goroutines to be stopped gracefully:

```go
select {
case job := <-jobs:
    // Process job
case <-quit:
    // Clean up and exit
    return
}
```

### Ticker with Timeout

Combine periodic operations with a timeout:

```go
ticker := time.NewTicker(100 * time.Millisecond)
timeout := time.After(1 * time.Second)

for {
    select {
    case <-ticker.C:
        // Do periodic work
    case <-timeout:
        ticker.Stop()
        return
    }
}
```

### Multiplexing (Fan-In)

Combine multiple channels into one:

```go
func fanIn(input1, input2 <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        for {
            select {
            case msg := <-input1:
                out <- msg
            case msg := <-input2:
                out <- msg
            }
        }
    }()
    return out
}
```

### Priority Select

Give priority to certain channels:

```go
// Check high priority first
select {
case msg := <-highPriority:
    handle(msg)
default:
    // Only check low priority if high priority empty
    select {
    case msg := <-lowPriority:
        handle(msg)
    default:
    }
}
```

### Send Operations

Select works with both send and receive:

```go
select {
case ch1 <- value1:
    // Sent to ch1
case ch2 <- value2:
    // Sent to ch2
}
```

### Graceful Shutdown

Coordinate shutdown across multiple goroutines:

```go
func worker(data <-chan int, stop <-chan bool) {
    for {
        select {
        case d := <-data:
            process(d)
        case <-stop:
            // Clean up
            return
        }
    }
}
```

## Running the Example

```bash
go run main.go
```

## Best Practices

### Do:
- Use `select` for timeout patterns
- Use `default` case for non-blocking operations
- Implement quit/stop channels for graceful shutdown
- Use `time.After()` for timeouts
- Use `time.NewTicker()` for periodic operations
- Handle all relevant cases in your select

### Don't:
- Use select when only one channel is involved (just use `<-ch`)
- Assume order of execution when multiple channels are ready
- Forget to stop tickers when done
- Leave goroutines running indefinitely without a quit mechanism
- Block in case statements (keep them quick)

## Common Patterns

### Timeout Pattern
```go
select {
case result := <-operation():
    // Success
case <-time.After(timeout):
    // Timeout
}
```

### Done Channel Pattern
```go
done := make(chan bool)
select {
case <-work():
    // Work completed
case <-done:
    // Cancelled
}
```

### Heartbeat Pattern
```go
heartbeat := time.NewTicker(1 * time.Second)
for {
    select {
    case <-heartbeat.C:
        sendHeartbeat()
    case <-done:
        return
    }
}
```

### Rate Limiting
```go
rate := time.Tick(100 * time.Millisecond)
for req := range requests {
    <-rate  // Wait for rate limiter
    go handle(req)
}
```

## Common Pitfalls

### Empty Select

```go
select {}  // Blocks forever - sometimes used to keep program running
```

### Forgetting Default in Non-Blocking

```go
// BAD: Will block if channel not ready
select {
case msg := <-ch:
    process(msg)
}

// GOOD: Won't block
select {
case msg := <-ch:
    process(msg)
default:
    // Continue without blocking
}
```

### Not Stopping Tickers

```go
// BAD: Ticker keeps running (goroutine leak)
ticker := time.NewTicker(1 * time.Second)
select {
case <-ticker.C:
    // ...
}

// GOOD: Stop ticker when done
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()
select {
case <-ticker.C:
    // ...
}
```

### Random Selection

```go
// Both channels ready - either case might execute
ch1 <- "ready"
ch2 <- "ready"

select {
case <-ch1:
    // Might execute
case <-ch2:
    // Might execute
}
```

Selection is random when multiple cases are ready.

## Select vs If-Else

**Select**:
- For channel operations
- Can wait on multiple channels
- Random choice when multiple ready
- Blocks until a case is ready (unless default)

**If-Else**:
- For boolean conditions
- Sequential evaluation
- Deterministic choice
- No blocking behavior

## Key Takeaways

- `select` enables concurrent waiting on multiple channels
- Timeout pattern prevents indefinite blocking
- `default` case provides non-blocking operations
- Random selection when multiple channels are ready
- Essential for implementing graceful shutdown
- Combine with tickers for periodic operations
- Use quit channels to stop goroutines cleanly
- Fan-in pattern merges multiple channels

## Next Steps

See [04-mutexes](../04-mutexes/) to learn about protecting shared state with mutual exclusion locks, or [05-waitgroups](../05-waitgroups/) to learn about coordinating multiple goroutines.
