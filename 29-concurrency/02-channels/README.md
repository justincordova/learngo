# Channels

Channels are Go's way of allowing goroutines to communicate safely. They provide a typed conduit through which you can send and receive values, enabling synchronization and data sharing between concurrent operations.

## Key Concepts

### What is a Channel?

A channel is a communication mechanism that lets one goroutine send values to another goroutine:

```go
ch := make(chan int)  // Create a channel of integers
ch <- 42              // Send a value to the channel
value := <-ch         // Receive a value from the channel
```

### Channel Types

**Unbuffered Channels**
```go
ch := make(chan int)
```
- Synchronous communication
- Sender blocks until receiver is ready
- Receiver blocks until sender sends

**Buffered Channels**
```go
ch := make(chan int, 10)  // Buffer capacity of 10
```
- Asynchronous up to capacity
- Sender blocks only when buffer is full
- Receiver blocks only when buffer is empty

## Example Breakdown

### Basic Channel Operations

```go
ch := make(chan string)

// Send to channel (blocks until received)
ch <- "message"

// Receive from channel (blocks until available)
msg := <-ch
```

### Buffered vs Unbuffered

**Unbuffered**: Requires both sender and receiver to be ready simultaneously.

**Buffered**: Allows sending multiple values without blocking until buffer is full:

```go
ch := make(chan int, 3)
ch <- 1  // Doesn't block
ch <- 2  // Doesn't block
ch <- 3  // Doesn't block
// ch <- 4  // Would block - buffer full
```

### Closing Channels

```go
close(ch)  // Close the channel
```

Important rules:
- Only the sender should close a channel
- Closing indicates no more values will be sent
- Receiving from a closed channel returns zero value and `false`
- Sending to a closed channel causes a panic

### Range Over Channels

```go
for value := range ch {
    // Receives values until channel is closed
}
```

This automatically exits when the channel is closed.

### Checking if Channel is Closed

```go
value, ok := <-ch
if !ok {
    // Channel is closed
}
```

The second return value indicates whether the channel is open.

### Pipeline Pattern

Chain operations together using channels:

```go
numbers := generate(1, 2, 3)    // Stage 1: generate
squares := square(numbers)       // Stage 2: transform
// results := consume(squares)   // Stage 3: consume
```

Each stage:
- Receives values from an input channel
- Performs operations
- Sends results to an output channel
- Closes output when input is exhausted

### Fan-Out/Fan-In Pattern

**Fan-Out**: Distribute work across multiple goroutines
```go
worker1 := process(input)
worker2 := process(input)
worker3 := process(input)
```

**Fan-In**: Merge results from multiple goroutines
```go
results := merge(worker1, worker2, worker3)
```

This pattern enables parallel processing of data.

### Channel Direction

Restrict channel operations in function signatures:

```go
// Send-only channel
func send(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func receive(ch <-chan int) {
    value := <-ch
}
```

Benefits:
- Type safety: prevents misuse
- Clear intent: documents expected usage
- Compile-time checks

### Worker Pool Pattern

Create a pool of workers to process jobs concurrently:

```go
jobs := make(chan Job, 100)
results := make(chan Result, 100)

// Start workers
for w := 1; w <= 10; w++ {
    go worker(w, jobs, results)
}

// Send jobs
for j := range jobList {
    jobs <- j
}
close(jobs)

// Collect results
for r := range results {
    // Process result
}
```

## Running the Example

```bash
go run main.go
```

## Best Practices

### Do:
- Close channels when done sending (sender's responsibility)
- Use buffered channels for known capacity scenarios
- Use channel direction in function signatures
- Range over channels when consuming all values
- Use the comma-ok idiom to check if channel is closed

### Don't:
- Close a channel from the receiver side
- Send on a closed channel (causes panic)
- Close a channel multiple times (causes panic)
- Use channels when a simple mutex would suffice
- Leave goroutines running indefinitely

## Common Patterns

### Generator Pattern
```go
func generate() <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
        close(ch)
    }()
    return ch
}
```

### Pipeline Stage
```go
func stage(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for v := range in {
            out <- v * 2  // Transform
        }
        close(out)
    }()
    return out
}
```

### Quit Channel
```go
quit := make(chan bool)
go func() {
    // Do work
    quit <- true
}()
<-quit  // Wait for signal
```

## Common Pitfalls

### Deadlock

```go
// BAD: Deadlock - nobody to receive
ch := make(chan int)
ch <- 1  // Blocks forever on unbuffered channel
```

### Forgetting to Close

```go
// BAD: Range will wait forever
ch := make(chan int)
go func() {
    ch <- 1
    // Forgot to close(ch)
}()
for v := range ch {  // Blocks forever after receiving 1
    fmt.Println(v)
}
```

### Closing from Receiver

```go
// BAD: Receiver shouldn't close
go func() {
    value := <-ch
    close(ch)  // Wrong: sender should close
}()
```

## Key Takeaways

- Channels enable safe communication between goroutines
- Unbuffered channels synchronize; buffered channels allow asynchrony
- Always close channels from the sender side
- Use channel direction to make intent clear
- Pipelines and fan-out/fan-in are powerful concurrency patterns
- Range over channels to consume all values
- Check `ok` value to detect closed channels

## Next Steps

See [03-channel-select](../03-channel-select/) to learn how to work with multiple channels using the `select` statement.
