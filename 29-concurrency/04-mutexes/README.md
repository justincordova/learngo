# Mutexes

Mutexes (mutual exclusion locks) are synchronization primitives that protect shared data from concurrent access. They ensure that only one goroutine can access a critical section of code at a time.

## Key Concepts

### What is a Mutex?

A mutex is a lock that provides mutually exclusive access to shared resources:

```go
var mu sync.Mutex

mu.Lock()    // Acquire the lock
// critical section - only one goroutine at a time
mu.Unlock()  // Release the lock
```

### Why Do We Need Mutexes?

**Without mutex - Race Condition:**
```go
counter := 0
// Multiple goroutines
counter++  // UNSAFE: read, increment, write are not atomic
```

**With mutex - Thread-Safe:**
```go
var mu sync.Mutex
counter := 0
// Multiple goroutines
mu.Lock()
counter++  // SAFE: protected by mutex
mu.Unlock()
```

## Types of Mutexes

### sync.Mutex

Standard mutex providing exclusive access:

```go
type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}
```

**Use when:**
- You have both reads and writes
- Operations are quick
- Write frequency is high

### sync.RWMutex

Read/Write mutex allowing multiple concurrent readers:

```go
type RWCounter struct {
    mu    sync.RWMutex
    value int
}

func (c *RWCounter) Read() int {
    c.mu.RLock()    // Multiple readers allowed
    defer c.mu.RUnlock()
    return c.value
}

func (c *RWCounter) Write(v int) {
    c.mu.Lock()     // Exclusive write access
    defer c.mu.Unlock()
    c.value = v
}
```

**Use when:**
- Reads are much more frequent than writes
- Read operations take significant time
- You need better read performance

## Example Breakdown

### Race Condition Without Mutex

```go
counter := 0
var wg sync.WaitGroup

for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        counter++  // RACE: Not thread-safe!
    }()
}

wg.Wait()
// counter will likely be < 100 due to lost updates
```

**What happens:**
1. Goroutine A reads `counter` (value: 5)
2. Goroutine B reads `counter` (value: 5)
3. Goroutine A increments and writes 6
4. Goroutine B increments and writes 6
5. Result: Two increments, but counter only increased by 1!

### Fixed with Mutex

```go
var mu sync.Mutex
counter := 0
var wg sync.WaitGroup

for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        mu.Lock()
        counter++  // SAFE: Protected by mutex
        mu.Unlock()
    }()
}

wg.Wait()
// counter will always be 100
```

### Critical Sections

Protect multiple operations that must execute atomically:

```go
type BankAccount struct {
    mu      sync.Mutex
    balance int
}

func (a *BankAccount) Withdraw(amount int) bool {
    a.mu.Lock()
    defer a.mu.Unlock()

    // Critical section: check and modify must be atomic
    if a.balance >= amount {
        a.balance -= amount
        return true
    }
    return false
}
```

### RWMutex for Read-Heavy Workloads

```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

// Read operation - allows concurrent reads
func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

// Write operation - exclusive access
func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}
```

**Performance benefits:**
- Multiple `Get()` calls can run simultaneously
- `Set()` calls wait for exclusive access
- Ideal when reads >> writes

## Running the Example

```bash
go run main.go
```

### Detect Race Conditions

Run with the race detector to find concurrency bugs:

```bash
go run -race main.go
```

The race detector will report:
- Where races occur
- Which goroutines are involved
- The conflicting memory access

## Best Practices

### Do:

**Always use defer to unlock:**
```go
mu.Lock()
defer mu.Unlock()  // Ensures unlock even if panic occurs
// ... critical section ...
```

**Keep critical sections small:**
```go
// Good: minimal locked section
mu.Lock()
value := counter
counter = value + 1
mu.Unlock()

// Bad: unnecessary locking
mu.Lock()
value := counter
processData()      // Don't hold lock during slow operations!
counter = value + 1
mu.Unlock()
```

**Embed mutex in structs:**
```go
type Counter struct {
    mu    sync.Mutex
    value int
}
```

**Use RWMutex for read-heavy scenarios:**
```go
// Many reads, few writes
cache.mu.RLock()    // Better than Lock() for reads
value := cache.data[key]
cache.mu.RUnlock()
```

### Don't:

**Don't copy mutexes:**
```go
// Bad: copying a mutex copies its state
counter1 := Counter{value: 0}
counter2 := counter1  // BUG: copies the mutex!

// Good: use pointers
counter1 := &Counter{value: 0}
counter2 := counter1  // OK: same mutex
```

**Don't lock twice in same goroutine:**
```go
// Bad: deadlock!
mu.Lock()
mu.Lock()  // Deadlock: already locked
mu.Unlock()
mu.Unlock()

// Good: lock once
mu.Lock()
// ... all work ...
mu.Unlock()
```

**Don't hold locks during slow operations:**
```go
// Bad: holding lock during I/O
mu.Lock()
data := readFromDatabase()  // Slow!
cache[key] = data
mu.Unlock()

// Good: lock only for cache access
data := readFromDatabase()
mu.Lock()
cache[key] = data
mu.Unlock()
```

**Don't forget to unlock:**
```go
// Bad: forget to unlock on error
mu.Lock()
if err != nil {
    return err  // BUG: mutex still locked!
}
mu.Unlock()

// Good: use defer
mu.Lock()
defer mu.Unlock()
if err != nil {
    return err  // OK: defer ensures unlock
}
```

## When to Use Mutex vs Channels

### Use Mutex when:
- Protecting access to shared state
- Simple read/write operations
- Performance-critical code
- Working with existing data structures

### Use Channels when:
- Passing ownership of data
- Coordinating between goroutines
- Implementing pipelines
- Broadcasting to multiple goroutines

**Remember:** "Don't communicate by sharing memory; share memory by communicating." But mutexes are still the right tool for protecting shared state.

## Common Patterns

### Singleton with Mutex

```go
var (
    instance *Database
    mu       sync.Mutex
)

func GetDatabase() *Database {
    mu.Lock()
    defer mu.Unlock()

    if instance == nil {
        instance = &Database{}
    }
    return instance
}
```

### Read-Write Map

```go
type SafeMap struct {
    mu   sync.RWMutex
    data map[string]int
}

func (m *SafeMap) Get(key string) (int, bool) {
    m.mu.RLock()
    defer m.mu.RUnlock()
    val, ok := m.data[key]
    return val, ok
}

func (m *SafeMap) Set(key string, value int) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.data[key] = value
}
```

### Conditional Update

```go
func (c *Counter) IncrementIfPositive() bool {
    c.mu.Lock()
    defer c.mu.Unlock()

    if c.value > 0 {
        c.value++
        return true
    }
    return false
}
```

## Common Pitfalls

### Deadlock

```go
// Circular wait
mu1.Lock()
mu2.Lock()  // Another goroutine has mu2 and wants mu1
// DEADLOCK!
```

**Solution:** Always acquire locks in the same order.

### Copying a Locked Mutex

```go
type Counter struct {
    mu    sync.Mutex
    value int
}

// Bad: passing by value copies the mutex
func increment(c Counter) {  // BUG!
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}

// Good: pass by pointer
func increment(c *Counter) {
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}
```

### Forgetting defer

```go
// Risky: might forget to unlock on early return
mu.Lock()
if condition {
    mu.Unlock()
    return
}
// ... more code ...
mu.Unlock()

// Safe: defer handles all cases
mu.Lock()
defer mu.Unlock()
if condition {
    return
}
// ... more code ...
```

## Performance Considerations

### Mutex vs RWMutex

```go
// Benchmark results (typical)
// 90% reads, 10% writes:
// sync.Mutex:    ~1000 ns/op
// sync.RWMutex:  ~100 ns/op  (10x faster for reads)

// 50% reads, 50% writes:
// sync.Mutex:    ~1000 ns/op
// sync.RWMutex:  ~1200 ns/op  (slightly slower due to overhead)
```

**Choose RWMutex when:**
- Read-to-write ratio > 10:1
- Critical sections are small
- Contention is high

## Key Takeaways

- Mutexes prevent race conditions by ensuring exclusive access
- Always use `defer` to unlock mutexes
- sync.Mutex: exclusive access for all operations
- sync.RWMutex: allows concurrent reads, exclusive writes
- Keep critical sections as small as possible
- Use the race detector (`-race`) to find bugs
- Don't copy mutexes (use pointers to structs)
- RWMutex is faster for read-heavy workloads
- Prefer channels for communication, mutexes for protecting state

## Next Steps

See [05-waitgroups](../05-waitgroups/) to learn how to coordinate multiple goroutines.
