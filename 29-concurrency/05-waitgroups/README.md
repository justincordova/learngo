# WaitGroups

WaitGroups are a synchronization primitive that allow you to wait for a collection of goroutines to finish executing. They provide a simple way to coordinate goroutine completion without using channels.

## Key Concepts

### What is a WaitGroup?

A WaitGroup maintains a counter that tracks how many goroutines are still running:

```go
var wg sync.WaitGroup

wg.Add(1)    // Increment counter
go func() {
    defer wg.Done()  // Decrement counter when done
    // ... do work ...
}()

wg.Wait()    // Block until counter reaches 0
```

### The Three Operations

**Add(delta int)** - Increment the counter
```go
wg.Add(1)   // Add 1 to the counter
wg.Add(5)   // Add 5 to the counter
```

**Done()** - Decrement the counter (equivalent to Add(-1))
```go
defer wg.Done()  // Decrement when function exits
```

**Wait()** - Block until counter reaches zero
```go
wg.Wait()  // Blocks until all Done() calls complete
```

## Example Breakdown

### Basic Pattern

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)  // BEFORE starting goroutine

    go func(id int) {
        defer wg.Done()  // ALWAYS use defer
        fmt.Printf("Worker %d\n", id)
    }(i)
}

wg.Wait()  // Wait for all to complete
fmt.Println("All done")
```

**Flow:**
1. Counter starts at 0
2. Add(1) increments to 1 (repeated 5 times → counter = 5)
3. Each goroutine calls Done() → counter decrements
4. Wait() blocks until counter reaches 0
5. Program continues after Wait()

### Worker Pool with WaitGroup

```go
const numWorkers = 3
jobs := make(chan int, 10)
var wg sync.WaitGroup

// Start workers
for i := 0; i < numWorkers; i++ {
    wg.Add(1)
    go worker(i, jobs, &wg)
}

// Send jobs
for j := 0; j < 10; j++ {
    jobs <- j
}
close(jobs)

wg.Wait()  // Wait for all workers to finish
```

**Worker function:**
```go
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()  // Decrement when worker exits

    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        // ... process job ...
    }
}
```

### Collecting Results with Channels

```go
results := make(chan int, numWorkers)
var wg sync.WaitGroup

// Start workers
for i := 0; i < numWorkers; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        results <- n * n  // Send result
    }(i)
}

// Close results when all workers done
go func() {
    wg.Wait()
    close(results)
}()

// Collect results
for result := range results {
    fmt.Println(result)
}
```

**Pattern:**
1. Workers send results to channel
2. Separate goroutine waits for all workers and closes channel
3. Main goroutine ranges over channel until it's closed

### Nested WaitGroups

```go
var outerWG sync.WaitGroup

for _, batch := range batches {
    outerWG.Add(1)

    go func(b Batch) {
        defer outerWG.Done()

        var innerWG sync.WaitGroup

        // Process tasks in batch
        for _, task := range b.Tasks {
            innerWG.Add(1)
            go func(t Task) {
                defer innerWG.Done()
                process(t)
            }(task)
        }

        innerWG.Wait()  // Wait for batch tasks
    }(batch)
}

outerWG.Wait()  // Wait for all batches
```

**Use case:** When you have hierarchical work (batches of tasks).

### Error Handling

```go
errors := make(chan error, numWorkers)
var wg sync.WaitGroup

for i := 0; i < numWorkers; i++ {
    wg.Add(1)

    go func(id int) {
        defer wg.Done()

        if err := doWork(id); err != nil {
            errors <- err
        }
    }(i)
}

// Close errors channel when done
go func() {
    wg.Wait()
    close(errors)
}()

// Collect errors
for err := range errors {
    fmt.Println("Error:", err)
}
```

**Pattern:** Combine WaitGroup with error channel to handle failures.

## Running the Example

```bash
go run main.go
```

## Best Practices

### Do:

**Always Add before starting goroutine:**
```go
// Good
wg.Add(1)
go func() { ... }()

// Bad - race condition!
go func() {
    wg.Add(1)  // Might race with Wait()
    // ...
}()
```

**Always use defer with Done:**
```go
go func() {
    defer wg.Done()  // Ensures Done is called even on panic
    // ... work ...
}()
```

**Pass WaitGroup by pointer:**
```go
func worker(wg *sync.WaitGroup) {  // Correct: pointer
    defer wg.Done()
    // ...
}

// NOT: func worker(wg sync.WaitGroup)  // Wrong: copies WaitGroup
```

**Add outside the loop when count is known:**
```go
// Efficient
wg.Add(10)
for i := 0; i < 10; i++ {
    go work()
}

// Less efficient but okay
for i := 0; i < 10; i++ {
    wg.Add(1)
    go work()
}
```

### Don't:

**Don't Add inside goroutine:**
```go
// Bad - race between Add and Wait
for i := 0; i < 10; i++ {
    go func() {
        wg.Add(1)  // WRONG
        defer wg.Done()
    }()
}
wg.Wait()
```

**Don't forget Done:**
```go
go func() {
    // WRONG: no Done() call
    doWork()
}()
wg.Wait()  // Will deadlock!
```

**Don't reuse WaitGroup without resetting:**
```go
var wg sync.WaitGroup

// First use
wg.Add(5)
// ... goroutines call Done() ...
wg.Wait()

// Second use - OK, counter is back to 0
wg.Add(3)
// ... goroutines call Done() ...
wg.Wait()
```

**Don't call Add with 0:**
```go
wg.Add(0)  // No-op, but confusing
```

## Common Patterns

### Fan-Out/Fan-In

**Fan-Out:** Launch multiple workers
```go
var wg sync.WaitGroup
for i := 0; i < numWorkers; i++ {
    wg.Add(1)
    go worker(i, &wg)
}
```

**Fan-In:** Collect results
```go
results := make(chan Result, numWorkers)

go func() {
    wg.Wait()
    close(results)
}()

for result := range results {
    process(result)
}
```

### Batch Processing

```go
const batchSize = 100
items := getAllItems()

for i := 0; i < len(items); i += batchSize {
    end := i + batchSize
    if end > len(items) {
        end = len(items)
    }

    wg.Add(1)
    go func(batch []Item) {
        defer wg.Done()
        processBatch(batch)
    }(items[i:end])
}

wg.Wait()
```

### Timeout with WaitGroup

```go
done := make(chan struct{})

go func() {
    wg.Wait()
    close(done)
}()

select {
case <-done:
    fmt.Println("All workers completed")
case <-time.After(5 * time.Second):
    fmt.Println("Timeout waiting for workers")
}
```

### Early Exit on Error

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

errors := make(chan error, 1)
var wg sync.WaitGroup

for i := 0; i < numWorkers; i++ {
    wg.Add(1)

    go func(id int) {
        defer wg.Done()

        select {
        case <-ctx.Done():
            return  // Stop if another worker failed
        default:
        }

        if err := doWork(id); err != nil {
            select {
            case errors <- err:
                cancel()  // Signal other workers to stop
            default:
            }
        }
    }(i)
}

wg.Wait()
close(errors)

if err := <-errors; err != nil {
    fmt.Println("Failed:", err)
}
```

## Common Pitfalls

### Race Between Add and Wait

```go
// Wrong: Wait might execute before Add
for i := 0; i < 5; i++ {
    go func() {
        wg.Add(1)  // Race!
        defer wg.Done()
    }()
}
wg.Wait()

// Correct: Add before goroutine
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
    }()
}
wg.Wait()
```

### Forgetting to Pass Pointer

```go
// Wrong: WaitGroup copied by value
func worker(wg sync.WaitGroup) {
    defer wg.Done()  // Decrements copy, not original!
}

var wg sync.WaitGroup
wg.Add(1)
worker(wg)  // Pass by value - BUG!
wg.Wait()   // Deadlock!

// Correct: Pass pointer
func worker(wg *sync.WaitGroup) {
    defer wg.Done()
}

wg.Add(1)
worker(&wg)  // Pass by pointer
wg.Wait()    // Works!
```

### Mismatched Add/Done

```go
// Wrong: More Done than Add
wg.Add(1)
go func() {
    defer wg.Done()
    doWork()
    wg.Done()  // Called twice!
}()
wg.Wait()  // Panic: negative WaitGroup counter

// Wrong: More Add than Done
wg.Add(2)  // Added 2
go func() {
    defer wg.Done()  // Only 1 Done
    doWork()
}()
wg.Wait()  // Deadlock!
```

### Closure Variable Gotcha

```go
// Wrong: All goroutines see final value
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println(i)  // All print 5!
    }()
}

// Correct: Pass as parameter
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println(id)  // Prints 0,1,2,3,4
    }(i)
}
```

## When to Use WaitGroup

### Use WaitGroup when:
- You need to wait for a known number of goroutines
- You don't need to collect results (or use channel alongside)
- You want simple goroutine coordination
- Fan-out/fan-in pattern

### Don't use WaitGroup when:
- You need complex coordination (use channels/context)
- You need to pass data between goroutines (use channels)
- You need cancellation (use context)
- Only one goroutine (just call it directly)

## Key Takeaways

- WaitGroup coordinates goroutine completion with Add/Done/Wait
- Always call Add before starting the goroutine
- Always use defer with Done to handle panics
- Pass WaitGroup by pointer to functions
- Combine with channels to collect results
- Combine with context for cancellation
- Use the race detector (`-race`) to find bugs
- WaitGroup is for synchronization, channels are for communication

## Next Steps

See [06-waitgroup-go-method](../06-waitgroup-go-method/) to learn about the new WaitGroup.Go() method introduced in Go 1.25.
