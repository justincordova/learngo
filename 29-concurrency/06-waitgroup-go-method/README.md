# WaitGroup.Go() Method (Go 1.25+)

**Requires Go 1.25 or later**

The `WaitGroup.Go()` method is a new addition in Go 1.25 that simplifies the WaitGroup pattern by automatically handling `Add` and `Done` operations. It makes concurrent code safer and more concise.

## Key Concepts

### What is WaitGroup.Go()?

`Go()` is a convenience method that combines three operations into one:

```go
// OLD WAY (pre-1.25)
wg.Add(1)
go func() {
    defer wg.Done()
    doWork()
}()

// NEW WAY (Go 1.25+)
wg.Go(doWork)  // Add, go, and defer Done all in one!
```

### Method Signature

```go
func (wg *WaitGroup) Go(f func())
```

**What it does:**
1. Calls `wg.Add(1)`
2. Launches `go f()`
3. Automatically calls `defer wg.Done()` in the goroutine

## Why Use WaitGroup.Go()?

### Problem with Traditional Pattern

The traditional WaitGroup pattern is error-prone:

```go
// Mistake 1: Forgot to Add
go func() {
    defer wg.Done()  // Panic! Counter goes negative
    doWork()
}()

// Mistake 2: Forgot to Done
wg.Add(1)
go func() {
    doWork()  // Deadlock! Counter never reaches zero
}()

// Mistake 3: Variable capture
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println(i)  // Bug! All print 5
    }()
}
```

### Solution with Go()

```go
// Impossible to forget Add/Done
for i := 0; i < 5; i++ {
    wg.Go(func() {
        doWork()  // Add and Done handled automatically!
    })
}

// Can pass function directly
wg.Go(doWork)

// Variable capture still needs care, but Add/Done are guaranteed
for i := 0; i < 5; i++ {
    id := i  // Still need to capture properly
    wg.Go(func() {
        process(id)
    })
}
```

## Example Breakdown

### Basic Comparison

**OLD PATTERN:**
```go
var wg sync.WaitGroup

wg.Add(1)  // Step 1: Increment counter

go func() {  // Step 2: Launch goroutine
    defer wg.Done()  // Step 3: Defer decrement
    fmt.Println("Hello")
}()

wg.Wait()
```

**NEW PATTERN:**
```go
var wg sync.WaitGroup

wg.Go(func() {  // All three steps in one!
    fmt.Println("Hello")
})

wg.Wait()
```

### Worker Pool Pattern

**OLD WAY:**
```go
var wg sync.WaitGroup

for i := 0; i < numWorkers; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        worker(id, jobs)
    }(i)
}

wg.Wait()
```

**NEW WAY:**
```go
var wg sync.WaitGroup

for i := 0; i < numWorkers; i++ {
    workerID := i
    wg.Go(func() {
        worker(workerID, jobs)
    })
}

wg.Wait()
```

### Passing Functions Directly

```go
var wg sync.WaitGroup

tasks := []func(){
    initDatabase,
    loadConfig,
    startServer,
}

// OLD: Required wrapper
for _, task := range tasks {
    wg.Add(1)
    go func(t func()) {
        defer wg.Done()
        t()
    }(task)
}

// NEW: Can pass directly
for _, task := range tasks {
    wg.Go(task)  // Much cleaner!
}

wg.Wait()
```

### Pipeline Pattern

```go
var wg sync.WaitGroup

// Stage 1: Generate
numbers := make(chan int)
wg.Go(func() {
    for i := 0; i < 10; i++ {
        numbers <- i
    }
    close(numbers)
})

// Stage 2: Process
results := make(chan int)
wg.Go(func() {
    for n := range numbers {
        results <- n * n
    }
    close(results)
})

// Stage 3: Consume
wg.Go(func() {
    for r := range results {
        fmt.Println(r)
    }
})

wg.Wait()
```

## Running the Example

**Requires Go 1.25+**

```bash
go version  # Ensure Go 1.25 or later
go run main.go
```

## Migration Guide

### Before (Go 1.24 and earlier)

```go
func processItems(items []Item) {
    var wg sync.WaitGroup

    for _, item := range items {
        wg.Add(1)                    // Manual Add
        go func(i Item) {
            defer wg.Done()          // Manual Done
            process(i)
        }(item)                      // Variable capture
    }

    wg.Wait()
}
```

### After (Go 1.25+)

```go
func processItems(items []Item) {
    var wg sync.WaitGroup

    for _, item := range items {
        i := item                    // Still need to capture
        wg.Go(func() {              // No Add/Done needed!
            process(i)
        })
    }

    wg.Wait()
}
```

## Best Practices

### Do:

**Use Go() for all new code:**
```go
wg.Go(func() {
    doWork()
})
```

**Still capture loop variables properly:**
```go
for _, item := range items {
    i := item  // Capture is still necessary
    wg.Go(func() {
        process(i)
    })
}
```

**Pass functions directly when possible:**
```go
wg.Go(processData)  // Clean and simple
```

**Use with function slices:**
```go
for _, fn := range functions {
    wg.Go(fn)
}
```

### Don't:

**Don't manually Add/Done when using Go():**
```go
// Wrong: Go() already does Add
wg.Add(1)
wg.Go(func() {
    doWork()
})

// Wrong: Go() already does Done
wg.Go(func() {
    defer wg.Done()  // Unnecessary, causes negative counter!
    doWork()
})
```

**Don't assume variable capture is fixed:**
```go
// Still wrong: i is captured incorrectly
for i := 0; i < 10; i++ {
    wg.Go(func() {
        fmt.Println(i)  // All print 10!
    })
}

// Correct: Capture properly
for i := 0; i < 10; i++ {
    id := i
    wg.Go(func() {
        fmt.Println(id)  // Prints 0-9
    })
}
```

## What Problems Does Go() Solve?

### 1. Forgotten Add

**Before:**
```go
// Bug: Forgot Add
go func() {
    defer wg.Done()  // Panic: negative counter
    doWork()
}()
```

**After:**
```go
// Impossible: Go() always does Add
wg.Go(func() {
    doWork()
})
```

### 2. Forgotten Done

**Before:**
```go
// Bug: Forgot Done
wg.Add(1)
go func() {
    doWork()  // Deadlock: counter never decrements
}()
```

**After:**
```go
// Impossible: Go() always does Done
wg.Go(func() {
    doWork()
})
```

### 3. Add/Done Mismatch

**Before:**
```go
// Bug: Mismatched Add/Done
wg.Add(2)  // Added 2
go func() {
    defer wg.Done()  // Only 1 Done
    doWork()
}()
// Deadlock!
```

**After:**
```go
// Impossible: Go() ensures 1:1 ratio
wg.Go(func() {
    doWork()
})
```

## What Problems Does Go() NOT Solve?

### Variable Capture

```go
// Still wrong with Go()
for i := 0; i < 5; i++ {
    wg.Go(func() {
        fmt.Println(i)  // All print 5!
    })
}

// Still need proper capture
for i := 0; i < 5; i++ {
    id := i  // Capture in loop
    wg.Go(func() {
        fmt.Println(id)  // Correct
    })
}

// Or pass as parameter (can't do with Go())
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println(id)
    }(i)
}
```

## Performance

`Go()` has **no performance overhead** compared to manual Add/Done:
- Same number of goroutines
- Same synchronization primitives
- Just syntactic sugar

## Backwards Compatibility

Go() is **purely additive**:
- Old code still works unchanged
- Can mix old and new patterns in same codebase
- Gradual migration is safe

```go
// Mixing is fine
var wg sync.WaitGroup

// Old way
wg.Add(1)
go func() {
    defer wg.Done()
    oldWork()
}()

// New way
wg.Go(func() {
    newWork()
})

wg.Wait()  // Works perfectly
```

## Common Patterns with Go()

### Fan-Out

```go
var wg sync.WaitGroup

for i := 0; i < numWorkers; i++ {
    wg.Go(func() {
        worker(jobs)
    })
}

wg.Wait()
```

### Parallel Processing

```go
var wg sync.WaitGroup

for _, item := range items {
    i := item
    wg.Go(func() {
        process(i)
    })
}

wg.Wait()
```

### Timeout with Context

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

var wg sync.WaitGroup

for i := 0; i < numTasks; i++ {
    wg.Go(func() {
        doWorkWithContext(ctx)
    })
}

wg.Wait()
```

### Error Collection

```go
var wg sync.WaitGroup
errors := make(chan error, numWorkers)

for i := 0; i < numWorkers; i++ {
    id := i
    wg.Go(func() {
        if err := doWork(id); err != nil {
            errors <- err
        }
    })
}

go func() {
    wg.Wait()
    close(errors)
}()

for err := range errors {
    log.Println("Error:", err)
}
```

## Key Takeaways

- `WaitGroup.Go()` combines `Add(1)`, `go`, and `defer Done()` into one call
- **Requires Go 1.25 or later**
- Prevents common mistakes: forgotten Add, forgotten Done, mismatched counts
- **Does NOT fix variable capture issues** - still need to capture properly
- Zero performance overhead - just cleaner syntax
- Fully backwards compatible - can mix with old pattern
- Use for all new code when on Go 1.25+
- Makes concurrent code safer and more maintainable

## Next Steps

- See [04-mutexes](../04-mutexes/) for protecting shared state
- See [05-waitgroups](../05-waitgroups/) for the traditional WaitGroup pattern
- Explore the context package for cancellation and timeouts
