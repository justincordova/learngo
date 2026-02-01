# Concurrency Notes

## Goroutines

Goroutines are lightweight threads managed by the Go runtime. They're much cheaper than OS threads.

### Starting a Goroutine

```go
go someFunction()

go func() {
    // anonymous function as goroutine
    fmt.Println("Running concurrently")
}()
```

### Goroutine Characteristics

- **Lightweight**: ~2KB initial stack size (vs ~1MB for OS threads)
- **Multiplexed**: Many goroutines run on fewer OS threads
- **Managed**: Go runtime handles scheduling
- **Cheap**: Can easily create thousands of goroutines

## Channels

Channels provide safe communication between goroutines. They're the recommended way to share data.

### Creating Channels

```go
ch := make(chan int)        // Unbuffered channel
ch := make(chan int, 10)    // Buffered channel with capacity 10
```

### Sending and Receiving

```go
ch <- 42      // Send to channel
value := <-ch // Receive from channel
```

### Channel Directions

Specify send-only or receive-only channels in function signatures:

```go
func sender(ch chan<- int) {
    ch <- 42  // Can only send
}

func receiver(ch <-chan int) {
    value := <-ch  // Can only receive
}
```

### Closing Channels

```go
close(ch)

// Check if channel is closed
value, ok := <-ch
if !ok {
    // Channel is closed
}

// Range over channel (stops when closed)
for value := range ch {
    fmt.Println(value)
}
```

## Select Statement

The `select` statement multiplexes multiple channel operations:

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case ch3 <- 42:
    fmt.Println("Sent to ch3")
default:
    fmt.Println("No channel ready")
}
```

### Timeout Pattern

```go
select {
case result := <-ch:
    fmt.Println("Got result:", result)
case <-time.After(5 * time.Second):
    fmt.Println("Timeout")
}
```

## Mutexes

Mutexes protect shared state when channels aren't appropriate.

### sync.Mutex

```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

### sync.RWMutex

Allows multiple readers or one writer:

```go
type Cache struct {
    mu    sync.RWMutex
    items map[string]string
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.items[key]
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.items[key] = value
}
```

## WaitGroups

Wait for multiple goroutines to complete:

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Printf("Worker %d done\n", id)
    }(i)
}

wg.Wait() // Wait for all goroutines to finish
```

## Worker Pool Pattern

A common pattern for limiting concurrent operations:

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 5; a++ {
        <-results
    }
}
```

## Race Conditions

Race conditions occur when multiple goroutines access shared data without synchronization.

### Detecting Races

Use the race detector during development:

```bash
go run -race main.go
go test -race ./...
go build -race
```

### Preventing Races

1. **Use channels** to communicate
2. **Use mutexes** to protect shared state
3. **Avoid shared state** when possible

## Best Practices

### Do:
- Use channels for communication between goroutines
- Use mutexes for protecting shared state
- Always call `defer wg.Done()` when using WaitGroups
- Close channels when you're done sending
- Use `-race` flag during development
- Prefer communication over shared memory

### Don't:
- Share memory without synchronization
- Close channels from the receiving end
- Send on closed channels (causes panic)
- Forget to call `wg.Done()` (causes deadlock)
- Ignore race detector warnings

## Goroutine Leaks

Prevent goroutine leaks:

```go
// BAD - goroutine may never finish
go func() {
    <-ch  // If nothing sends, goroutine leaks
}()

// GOOD - goroutine can be cancelled
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    select {
    case <-ch:
        // process
    case <-ctx.Done():
        return  // Cleanup on cancellation
    }
}()
```

## Common Patterns

### Fan-out, Fan-in

**Fan-out**: Multiple goroutines read from the same channel
```go
for i := 0; i < numWorkers; i++ {
    go worker(jobs)
}
```

**Fan-in**: Multiple goroutines send to the same channel
```go
results := make(chan Result)
go worker1(results)
go worker2(results)
go worker3(results)
```

### Pipeline Pattern

Chain goroutines together:

```go
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Usage
for n := range sq(sq(gen(2, 3))) {
    fmt.Println(n)  // Prints: 16, 81
}
```

## Resources

- "Don't communicate by sharing memory; share memory by communicating"
- [Go Blog: Concurrency is not parallelism](https://go.dev/blog/concurrency-is-not-parallelism)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
- Always test concurrent code with `-race` flag
