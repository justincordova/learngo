# Worker Pool Pattern

This example demonstrates the worker pool pattern in Go, a fundamental concurrency pattern for processing a queue of jobs with a fixed number of workers.

## What is a Worker Pool?

A worker pool is a concurrency pattern where:
- A fixed number of worker goroutines process jobs from a shared queue
- Jobs are sent to a channel
- Workers read from the channel and process jobs concurrently
- Results are collected through another channel

## Key Benefits

### 1. Bounded Parallelism
Control the maximum number of concurrent operations:
```go
const numWorkers = 5  // Only 5 goroutines, regardless of job count
```

### 2. Resource Management
- Prevents overwhelming system resources
- Limits concurrent network connections, file handles, etc.
- Provides backpressure when jobs arrive faster than processing

### 3. Graceful Shutdown
- Cleanly finish in-flight work before termination
- Cancel work when timeout occurs
- Handle errors without stopping the entire pool

### 4. Better Performance
- Reuse goroutines instead of creating one per job
- Reduce context switching overhead
- Process large job queues efficiently

## Worker Pool Components

### Jobs Channel
```go
jobs := make(chan Job, bufferSize)
```
- Buffered channel holds pending work
- Buffer size controls backpressure
- Closed when no more jobs will be sent

### Results Channel
```go
results := make(chan Result, bufferSize)
```
- Collects worker outputs
- Can include errors
- Closed after all workers finish

### Worker Goroutines
```go
for w := 1; w <= numWorkers; w++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        for job := range jobs {
            // Process job
            result := processJob(job)
            results <- result
        }
    }(w)
}
```

### WaitGroup for Synchronization
```go
var wg sync.WaitGroup
wg.Add(numWorkers)
// ... start workers ...
wg.Wait()
close(results)
```

## Examples in This Directory

### Example 1: Basic Worker Pool
Simple number processing demonstrating core pattern:
- Fixed worker count
- Job distribution
- Result collection

### Example 2: Error Handling
Shows how to handle errors in worker pools:
- Job and Result structs
- Error propagation
- Success/failure counting

### Example 3: URL Fetcher
Practical example fetching URLs concurrently:
- HTTP client in each worker
- Timeout handling
- Size and status tracking
- Real-world error handling

### Example 4: Image Processor
Simulates batch image processing:
- Variable processing times
- Duration tracking
- Performance metrics

### Example 5: Timeout and Cancellation
Advanced pattern with context:
- Context-based cancellation
- Graceful shutdown
- Timeout handling
- Partial result collection

## Running the Example

```bash
go run main.go
```

## Common Patterns

### Basic Pattern
```go
func workerPool() {
    const numWorkers = 3
    jobs := make(chan Job, 100)
    results := make(chan Result, 100)

    // Start workers
    var wg sync.WaitGroup
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // Send jobs
    for _, job := range allJobs {
        jobs <- job
    }
    close(jobs)

    // Close results after all workers done
    go func() {
        wg.Wait()
        close(results)
    }()

    // Collect results
    for result := range results {
        processResult(result)
    }
}
```

### With Context Cancellation
```go
func workerPoolWithCancel(ctx context.Context) {
    jobs := make(chan Job)
    results := make(chan Result)

    for w := 1; w <= numWorkers; w++ {
        go func() {
            for {
                select {
                case <-ctx.Done():
                    return
                case job, ok := <-jobs:
                    if !ok {
                        return
                    }
                    processJob(job, results)
                }
            }
        }()
    }
}
```

### With Error Handling
```go
type Result struct {
    Value interface{}
    Error error
}

func worker(jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        result := Result{}
        value, err := processJob(job)
        if err != nil {
            result.Error = err
        } else {
            result.Value = value
        }
        results <- result
    }
}
```

## Best Practices

### 1. Choose Worker Count Wisely
- **CPU-bound tasks**: `runtime.NumCPU()`
- **I/O-bound tasks**: Higher count (10-100s)
- **Memory-intensive**: Lower count to avoid OOM
- **External APIs**: Match rate limits

### 2. Use Buffered Channels
```go
jobs := make(chan Job, 100)  // Buffer prevents blocking
```
- Decouples job producers from workers
- Provides some backpressure
- Reduces goroutine blocking

### 3. Handle Errors Properly
- Don't panic in workers
- Return errors in result structs
- Log errors with context
- Consider retry logic

### 4. Graceful Shutdown
```go
// Close jobs channel first
close(jobs)

// Wait for workers to finish
wg.Wait()

// Then close results
close(results)
```

### 5. Use Context for Cancellation
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// In workers:
select {
case <-ctx.Done():
    return
case job := <-jobs:
    // Process job
}
```

### 6. Monitor and Metrics
```go
processed := atomic.Int64{}
errors := atomic.Int64{}

// In worker:
processed.Add(1)
if err != nil {
    errors.Add(1)
}
```

## When to Use Worker Pools

✅ **Good for:**
- Batch processing large datasets
- Concurrent API calls
- Image/video processing
- File I/O operations
- Database queries
- Web scraping

❌ **Not ideal for:**
- Very few jobs (overhead not worth it)
- Jobs with vastly different durations
- Real-time streaming (consider pipelines instead)
- Single job that needs all resources

## Performance Considerations

### Worker Count Impact
```
Too Few Workers:
  - Underutilizes resources
  - Longer total processing time

Too Many Workers:
  - Context switching overhead
  - Resource contention
  - Potential OOM issues

Sweet Spot:
  - CPU-bound: ~NumCPU()
  - I/O-bound: 10-100x NumCPU()
  - Benchmark your specific workload!
```

### Channel Buffer Size
```
No Buffer (0):
  - Job sender blocks until worker available
  - Tight coupling
  - Good for rate limiting

Small Buffer (10-100):
  - Smooths out bursts
  - Reduces blocking
  - Moderate memory usage

Large Buffer (1000+):
  - Maximum throughput
  - Higher memory usage
  - Risk of OOM with large jobs
```

## Advanced Patterns

### Priority Queue
Use separate channels for different priority jobs:
```go
highPriority := make(chan Job)
lowPriority := make(chan Job)

// Worker prefers high priority
select {
case job := <-highPriority:
    process(job)
case job := <-lowPriority:
    process(job)
}
```

### Rate Limiting
Add rate limiter per worker:
```go
limiter := time.NewTicker(100 * time.Millisecond)
defer limiter.Stop()

for job := range jobs {
    <-limiter.C  // Wait for rate limit
    process(job)
}
```

### Dynamic Worker Pool
Adjust worker count based on load:
```go
// Start with minimum workers
// Monitor queue depth
// Add/remove workers as needed
```

## Common Pitfalls

### ❌ Forgetting to Close Channels
```go
close(jobs)     // Required!
wg.Wait()
close(results)  // Required!
```

### ❌ Not Using WaitGroup
Results channel will close prematurely:
```go
// WRONG:
close(results)  // Closes before workers finish!

// RIGHT:
go func() {
    wg.Wait()
    close(results)
}()
```

### ❌ Sending After Close
```go
close(jobs)
jobs <- newJob  // PANIC!
```

### ❌ Deadlock on Unbuffered Results
```go
// If results is unbuffered and no consumer:
results <- result  // Workers block forever
```

## Testing Worker Pools

```go
func TestWorkerPool(t *testing.T) {
    jobs := make(chan int, 10)
    results := make(chan int, 10)

    // Start pool
    var wg sync.WaitGroup
    for w := 0; w < 3; w++ {
        wg.Add(1)
        go worker(jobs, results, &wg)
    }

    // Send test jobs
    for i := 1; i <= 10; i++ {
        jobs <- i
    }
    close(jobs)

    // Collect results
    go func() {
        wg.Wait()
        close(results)
    }()

    // Verify
    count := 0
    for range results {
        count++
    }

    if count != 10 {
        t.Errorf("Expected 10 results, got %d", count)
    }
}
```

## Key Takeaways

1. **Worker pools limit concurrency** to a fixed number of goroutines
2. **Use buffered channels** for jobs and results
3. **Close channels properly** in the right order
4. **WaitGroup coordinates** worker completion
5. **Context enables cancellation** and timeouts
6. **Handle errors** in result structs, don't panic
7. **Choose worker count** based on workload type
8. **Monitor performance** and adjust parameters

## Learn More

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- Try modifying the examples to process different workloads
- Experiment with different worker counts and buffer sizes
