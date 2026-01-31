package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter demonstrates unsafe concurrent access
type Counter struct {
	value int
}

// SafeCounter demonstrates thread-safe access with Mutex
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// RWCounter demonstrates read/write locking with RWMutex
type RWCounter struct {
	mu    sync.RWMutex
	value int
}

func main() {
	fmt.Println("Mutexes")
	fmt.Println("=======")
	fmt.Println()

	// Example 1: Race condition without mutex
	fmt.Println("1. Race condition without mutex:")
	fmt.Println("   (Run with 'go run -race main.go' to detect the race)")
	unsafeExample()
	fmt.Println()

	// Example 2: Thread-safe with sync.Mutex
	fmt.Println("2. Thread-safe access with sync.Mutex:")
	safeMutexExample()
	fmt.Println()

	// Example 3: Using RWMutex for read-heavy workloads
	fmt.Println("3. Using sync.RWMutex for read-heavy workloads:")
	rwMutexExample()
	fmt.Println()

	// Example 4: Comparing Mutex vs RWMutex performance
	fmt.Println("4. Performance comparison: Mutex vs RWMutex")
	performanceComparison()
	fmt.Println()

	// Example 5: Critical section example
	fmt.Println("5. Protecting critical sections:")
	criticalSectionExample()
	fmt.Println()

	// Example 6: Common pitfall - forgetting to unlock
	fmt.Println("6. Best practice: Using defer to unlock:")
	deferUnlockExample()
}

// unsafeExample demonstrates a race condition
func unsafeExample() {
	counter := &Counter{}
	var wg sync.WaitGroup

	// Start 100 goroutines that increment the counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// RACE CONDITION: Multiple goroutines accessing counter.value
			counter.value++
		}()
	}

	wg.Wait()
	fmt.Printf("   Unsafe counter value: %d (expected 100, likely wrong due to race)\n", counter.value)
}

// safeMutexExample shows proper mutex usage
func safeMutexExample() {
	counter := &SafeCounter{}
	var wg sync.WaitGroup

	// Start 100 goroutines that safely increment the counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.mu.Lock()
			counter.value++
			counter.mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("   Safe counter value: %d (always correct)\n", counter.value)
}

// rwMutexExample demonstrates RWMutex with multiple readers
func rwMutexExample() {
	counter := &RWCounter{}
	var wg sync.WaitGroup

	// Start multiple readers (can run concurrently)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			counter.mu.RLock() // Read lock - multiple readers allowed
			value := counter.value
			time.Sleep(10 * time.Millisecond) // Simulate read operation
			fmt.Printf("   Reader %d read value: %d\n", id, value)
			counter.mu.RUnlock()
		}(i)
	}

	// Start a writer (exclusive access)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Millisecond) // Let readers start first
		counter.mu.Lock() // Write lock - exclusive access
		counter.value = 42
		fmt.Println("   Writer updated value to 42")
		counter.mu.Unlock()
	}()

	wg.Wait()
}

// performanceComparison compares Mutex vs RWMutex
func performanceComparison() {
	const numReads = 1000
	const numWrites = 10

	// Test with Mutex
	mutexCounter := &SafeCounter{}
	mutexStart := time.Now()
	var wg1 sync.WaitGroup

	for i := 0; i < numReads; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			mutexCounter.mu.Lock()
			_ = mutexCounter.value
			mutexCounter.mu.Unlock()
		}()
	}

	for i := 0; i < numWrites; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			mutexCounter.mu.Lock()
			mutexCounter.value++
			mutexCounter.mu.Unlock()
		}()
	}

	wg1.Wait()
	mutexDuration := time.Since(mutexStart)

	// Test with RWMutex
	rwCounter := &RWCounter{}
	rwStart := time.Now()
	var wg2 sync.WaitGroup

	for i := 0; i < numReads; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			rwCounter.mu.RLock()
			_ = rwCounter.value
			rwCounter.mu.RUnlock()
		}()
	}

	for i := 0; i < numWrites; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			rwCounter.mu.Lock()
			rwCounter.value++
			rwCounter.mu.Unlock()
		}()
	}

	wg2.Wait()
	rwDuration := time.Since(rwStart)

	fmt.Printf("   Mutex time: %v\n", mutexDuration)
	fmt.Printf("   RWMutex time: %v\n", rwDuration)
	fmt.Println("   Note: RWMutex is typically faster for read-heavy workloads")
}

// criticalSectionExample shows protecting multiple operations
func criticalSectionExample() {
	type BankAccount struct {
		mu      sync.Mutex
		balance int
	}

	account := &BankAccount{balance: 1000}
	var wg sync.WaitGroup

	// Simulate multiple transactions
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Critical section: check balance and withdraw
			account.mu.Lock()
			if account.balance >= 300 {
				fmt.Printf("   Transaction %d: Current balance: $%d\n", id, account.balance)
				time.Sleep(10 * time.Millisecond) // Simulate processing
				account.balance -= 300
				fmt.Printf("   Transaction %d: Withdrew $300, new balance: $%d\n", id, account.balance)
			} else {
				fmt.Printf("   Transaction %d: Insufficient funds (balance: $%d)\n", id, account.balance)
			}
			account.mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Printf("   Final balance: $%d\n", account.balance)
}

// deferUnlockExample demonstrates the defer pattern
func deferUnlockExample() {
	type DataStore struct {
		mu   sync.Mutex
		data map[string]string
	}

	store := &DataStore{
		data: make(map[string]string),
	}

	// Good practice: Use defer to ensure unlock happens
	updateData := func(key, value string) {
		store.mu.Lock()
		defer store.mu.Unlock() // Ensures unlock even if panic occurs

		// Complex operations that might fail
		if key == "" {
			// Even if we return early, unlock still happens
			fmt.Println("   Error: empty key")
			return
		}

		store.data[key] = value
		fmt.Printf("   Updated: %s = %s\n", key, value)
	}

	var wg sync.WaitGroup
	keys := []string{"name", "", "age", "city"}
	values := []string{"Alice", "invalid", "30", "NYC"}

	for i := range keys {
		wg.Add(1)
		go func(k, v string) {
			defer wg.Done()
			updateData(k, v)
		}(keys[i], values[i])
	}

	wg.Wait()

	// Read final data
	store.mu.Lock()
	defer store.mu.Unlock()
	fmt.Printf("   Final data: %v\n", store.data)
}
