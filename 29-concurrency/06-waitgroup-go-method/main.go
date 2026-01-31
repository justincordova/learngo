package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("WaitGroup.Go() Method (Go 1.25+)")
	fmt.Println("==================================")
	fmt.Println()

	// Example 1: Old vs New pattern comparison
	fmt.Println("1. Comparing old pattern vs new WaitGroup.Go():")
	oldVsNewComparison()
	fmt.Println()

	// Example 2: Simplified worker pool
	fmt.Println("2. Simplified worker pool with Go():")
	simplifiedWorkerPool()
	fmt.Println()

	// Example 3: Error-free patterns with Go()
	fmt.Println("3. Go() prevents common mistakes:")
	preventingMistakes()
	fmt.Println()

	// Example 4: Nested function calls
	fmt.Println("4. Using Go() with nested function calls:")
	nestedFunctionCalls()
	fmt.Println()

	// Example 5: Cleaner pipeline pattern
	fmt.Println("5. Cleaner pipeline with Go():")
	cleanerPipeline()
	fmt.Println()

	// Example 6: Collecting results with Go()
	fmt.Println("6. Collecting results with Go():")
	collectingResults()
}

// oldVsNewComparison shows the difference between old and new patterns
func oldVsNewComparison() {
	fmt.Println("   OLD PATTERN (pre-1.25):")
	var wgOld sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wgOld.Add(1) // Must remember to Add

		go func(id int) {
			defer wgOld.Done() // Must remember defer Done
			fmt.Printf("   Old way: Worker %d\n", id)
		}(i) // Must capture variable
	}

	wgOld.Wait()

	fmt.Println("\n   NEW PATTERN (Go 1.25+):")
	var wgNew sync.WaitGroup

	for i := 1; i <= 3; i++ {
		// Go() automatically calls Add(1) and schedules Done()
		wgNew.Go(func() {
			fmt.Printf("   New way: Worker %d\n", i) // Closure works safely
		})
	}

	wgNew.Wait()
	fmt.Println("\n   Benefits: Less boilerplate, impossible to forget Add/Done!")
}

// simplifiedWorkerPool demonstrates a cleaner worker pool
func simplifiedWorkerPool() {
	const numWorkers = 3
	jobs := make(chan int, 10)

	var wg sync.WaitGroup

	// Start workers - much cleaner!
	for i := 1; i <= numWorkers; i++ {
		workerID := i // Capture for closure
		wgNew.Go(func() {
			processJobs(workerID, jobs)
		})
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	fmt.Println("   All jobs completed")
}

var wgNew sync.WaitGroup

func processJobs(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("   Worker %d processing job %d\n", id, job)
		time.Sleep(30 * time.Millisecond)
	}
}

// preventingMistakes shows how Go() prevents common errors
func preventingMistakes() {
	fmt.Println("   Old way - Easy to make mistakes:")
	fmt.Println("   // Mistake 1: Forgot to Add")
	fmt.Println("   // go func() { defer wg.Done(); ... }()  // Panic!")
	fmt.Println("   // Mistake 2: Forgot to Done")
	fmt.Println("   // wg.Add(1); go func() { ... }()  // Deadlock!")
	fmt.Println("   // Mistake 3: Variable capture")
	fmt.Println("   // for i := ... { go func() { use i }()  // Wrong value!")

	fmt.Println("\n   New way - Mistakes are impossible:")
	var wg sync.WaitGroup

	// Add/Done handled automatically
	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			// Variable capture issue still exists with closures,
			// but Add/Done are guaranteed correct
			fmt.Printf("   Worker completed (iteration %d environment)\n", i)
		})
	}

	wg.Wait()
	fmt.Println("   Note: Go() ensures Add/Done are always balanced!")
}

// nestedFunctionCalls shows Go() with function pointers
func nestedFunctionCalls() {
	var wg sync.WaitGroup

	tasks := []func(){
		func() { fmt.Println("   Task 1: Initialize database") },
		func() { fmt.Println("   Task 2: Load configuration") },
		func() { fmt.Println("   Task 3: Start services") },
	}

	// OLD WAY:
	// for _, task := range tasks {
	//     wg.Add(1)
	//     go func(t func()) {
	//         defer wg.Done()
	//         t()
	//     }(task)
	// }

	// NEW WAY - much cleaner:
	for _, task := range tasks {
		wg.Go(task) // Pass function directly!
	}

	wg.Wait()
	fmt.Println("   All tasks completed")
}

// cleanerPipeline demonstrates pipeline pattern with Go()
func cleanerPipeline() {
	var wg sync.WaitGroup

	// Stage 1: Generate numbers
	numbers := make(chan int, 5)
	wg.Go(func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	})

	// Stage 2: Square numbers
	squares := make(chan int, 5)
	wg.Go(func() {
		for n := range numbers {
			squares <- n * n
		}
		close(squares)
	})

	// Stage 3: Print results
	wg.Go(func() {
		for sq := range squares {
			fmt.Printf("   Square: %d\n", sq)
		}
	})

	wg.Wait()
}

// collectingResults shows collecting results with Go()
func collectingResults() {
	var wg sync.WaitGroup
	results := make(chan string, 5)

	// Simulate fetching data from different sources
	sources := []string{"Database", "Cache", "API", "File", "Network"}

	for _, source := range sources {
		src := source // Capture for closure
		wg.Go(func() {
			time.Sleep(50 * time.Millisecond)
			data := fmt.Sprintf("Data from %s", src)
			results <- data
		})
	}

	// Close results when all goroutines complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and print results
	for result := range results {
		fmt.Printf("   Received: %s\n", result)
	}
}
