# Goroutines Basics

Goroutines are lightweight threads managed by the Go runtime. They enable concurrent execution and are one of Go's most powerful features for building efficient, concurrent programs.

## Key Concepts

### What is a Goroutine?

A goroutine is a function that runs concurrently with other functions. Creating a goroutine is as simple as prefixing a function call with the `go` keyword:

```go
go functionName()
```

### Goroutine Characteristics

1. **Lightweight**: Goroutines start with only a few KB of stack space
2. **Cheap**: You can run thousands or millions of goroutines
3. **Managed**: The Go runtime handles scheduling and execution
4. **Concurrent**: They run concurrently with the main function and other goroutines

## Example Breakdown

### Basic Goroutine

```go
go printMessage("Hello from goroutine")
```

The `go` keyword starts the function in a new goroutine. The main function continues immediately without waiting.

### Anonymous Function Goroutines

```go
go func() {
    fmt.Println("Hello from anonymous goroutine")
}()
```

You can launch anonymous functions as goroutines. Note the `()` at the end to call the function.

### Multiple Concurrent Goroutines

```go
for i := 1; i <= 5; i++ {
    go printNumber(i)
}
```

Multiple goroutines run concurrently. Their execution order is not guaranteed.

### Closures and Variable Capture

```go
message := "Captured message"
go func() {
    fmt.Printf("Closure says: %s\n", message)
}()
```

Anonymous goroutines can capture variables from their enclosing scope.

### Synchronization with WaitGroup

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Do work
}()
wg.Wait() // Wait for all goroutines
```

`sync.WaitGroup` provides a way to wait for multiple goroutines to complete:
- `Add(n)`: Increment the counter by n
- `Done()`: Decrement the counter by 1
- `Wait()`: Block until counter is zero

### Race Conditions

Race conditions occur when multiple goroutines access shared data concurrently and at least one modifies it:

```go
// BAD: Race condition
counter := 0
go func() {
    counter++ // Multiple goroutines modifying counter
}()
```

### Detecting Race Conditions

Use the `-race` flag to detect race conditions:

```bash
go run -race main.go
```

The race detector will warn you about concurrent access to shared variables.

### Fixing Race Conditions

Use a `sync.Mutex` to protect shared data:

```go
var mu sync.Mutex
mu.Lock()
counter++
mu.Unlock()
```

Only one goroutine can hold the lock at a time, preventing concurrent access.

## Running the Example

Basic run:
```bash
go run main.go
```

With race detection:
```bash
go run -race main.go
```

## Best Practices

### Do:
- Use `sync.WaitGroup` to wait for goroutines to complete
- Protect shared data with mutexes or use channels
- Always test concurrent code with the `-race` flag
- Use `defer` to ensure locks are released
- Pass data to goroutines through function parameters when possible

### Don't:
- Assume execution order of goroutines
- Use `time.Sleep` as a primary synchronization mechanism
- Ignore race condition warnings
- Access shared variables without synchronization
- Create unbounded numbers of goroutines

## Common Pitfalls

### Goroutines and Loop Variables

```go
// BAD: Closure captures loop variable
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i) // All goroutines may print same value
    }()
}

// GOOD: Pass value as parameter
for i := 0; i < 5; i++ {
    go func(id int) {
        fmt.Println(id)
    }(i)
}
```

### Main Function Exits

If the main function exits, all goroutines are terminated immediately:

```go
// BAD: May not print anything
go fmt.Println("Hello")
// main exits immediately

// GOOD: Wait for goroutine
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    fmt.Println("Hello")
}()
wg.Wait()
```

## Key Takeaways

- Goroutines are the foundation of concurrent programming in Go
- Use the `go` keyword to launch functions concurrently
- Always synchronize goroutines properly (WaitGroup, channels, etc.)
- Race conditions are dangerous - use `-race` flag to detect them
- Mutexes protect shared data from concurrent access
- The main function won't wait for goroutines unless you tell it to

## Next Steps

See [02-channels](../02-channels/) to learn how to communicate safely between goroutines using channels.
