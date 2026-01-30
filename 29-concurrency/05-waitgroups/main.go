// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("WaitGroups")
	fmt.Println("==========")
	fmt.Println()

	// Example 1: Basic WaitGroup usage
	fmt.Println("1. Basic WaitGroup usage:")
	basicWaitGroup()
	fmt.Println()

	// Example 2: WaitGroup with multiple workers
	fmt.Println("2. Multiple workers with WaitGroup:")
	multipleWorkers()
	fmt.Println()

	// Example 3: WaitGroup with return values via channels
	fmt.Println("3. WaitGroup with channels for results:")
	waitGroupWithChannels()
	fmt.Println()

	// Example 4: Nested WaitGroups
	fmt.Println("4. Nested WaitGroups:")
	nestedWaitGroups()
	fmt.Println()

	// Example 5: WaitGroup with error handling
	fmt.Println("5. WaitGroup with error handling:")
	waitGroupWithErrors()
	fmt.Println()

	// Example 6: Common mistake - forgetting to pass pointer
	fmt.Println("6. Common mistake demonstration:")
	commonMistakes()
}

// basicWaitGroup demonstrates the fundamental WaitGroup pattern
func basicWaitGroup() {
	var wg sync.WaitGroup

	// Launch 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter before starting goroutine

		go func(id int) {
			defer wg.Done() // Decrement counter when done

			fmt.Printf("   Worker %d starting\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("   Worker %d done\n", id)
		}(i)
	}

	wg.Wait() // Block until counter reaches zero
	fmt.Println("   All workers completed")
}

// multipleWorkers shows a worker pool pattern with WaitGroup
func multipleWorkers() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("   All jobs processed")
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("   Worker %d processing job %d\n", id, job)
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("   Worker %d finished job %d\n", id, job)
	}
}

// waitGroupWithChannels demonstrates collecting results
func waitGroupWithChannels() {
	const numWorkers = 5
	results := make(chan int, numWorkers)
	var wg sync.WaitGroup

	// Start workers that compute squares
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
			result := n * n
			results <- result
			fmt.Printf("   Computed %d^2 = %d\n", n, result)
		}(i)
	}

	// Wait and close results channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	sum := 0
	for result := range results {
		sum += result
	}
	fmt.Printf("   Sum of squares: %d\n", sum)
}

// nestedWaitGroups shows using WaitGroups at different levels
func nestedWaitGroups() {
	var outerWG sync.WaitGroup

	// Simulate processing multiple batches
	batches := []string{"Batch-A", "Batch-B"}

	for _, batch := range batches {
		outerWG.Add(1)

		go func(batchName string) {
			defer outerWG.Done()

			fmt.Printf("   Starting %s\n", batchName)

			// Inner WaitGroup for tasks within this batch
			var innerWG sync.WaitGroup

			for i := 1; i <= 3; i++ {
				innerWG.Add(1)

				go func(taskID int) {
					defer innerWG.Done()
					fmt.Printf("   %s: Task %d running\n", batchName, taskID)
					time.Sleep(30 * time.Millisecond)
					fmt.Printf("   %s: Task %d complete\n", batchName, taskID)
				}(i)
			}

			innerWG.Wait() // Wait for all tasks in this batch
			fmt.Printf("   Completed %s\n", batchName)
		}(batch)
	}

	outerWG.Wait() // Wait for all batches
	fmt.Println("   All batches completed")
}

// waitGroupWithErrors demonstrates error handling with WaitGroups
func waitGroupWithErrors() {
	const numWorkers = 5
	errors := make(chan error, numWorkers)
	var wg sync.WaitGroup

	// Start workers that might fail
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Printf("   Worker %d starting\n", id)
			time.Sleep(30 * time.Millisecond)

			// Simulate error for worker 3
			if id == 3 {
				err := fmt.Errorf("worker %d failed", id)
				errors <- err
				fmt.Printf("   Worker %d: ERROR\n", id)
				return
			}

			fmt.Printf("   Worker %d: success\n", id)
		}(i)
	}

	// Wait and close error channel
	go func() {
		wg.Wait()
		close(errors)
	}()

	// Collect errors
	var errorList []error
	for err := range errors {
		errorList = append(errorList, err)
	}

	if len(errorList) > 0 {
		fmt.Printf("   Encountered %d error(s):\n", len(errorList))
		for _, err := range errorList {
			fmt.Printf("   - %v\n", err)
		}
	} else {
		fmt.Println("   All workers succeeded")
	}
}

// commonMistakes demonstrates common WaitGroup pitfalls
func commonMistakes() {
	// Mistake 1: Correct way - Add before goroutine starts
	fmt.Println("   Correct: Add before starting goroutine")
	var wg1 sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg1.Add(1) // CORRECT: Add before 'go'

		go func(id int) {
			defer wg1.Done()
			fmt.Printf("   Worker %d running\n", id)
		}(i)
	}
	wg1.Wait()

	// Mistake 2: What happens if we Add inside goroutine (potential race)
	fmt.Println("\n   Warning: Adding inside goroutine is risky")
	var wg2 sync.WaitGroup

	for i := 1; i <= 3; i++ {
		go func(id int) {
			wg2.Add(1) // RISKY: Race with Wait()
			defer wg2.Done()
			fmt.Printf("   Worker %d running\n", id)
		}(i)
	}
	time.Sleep(100 * time.Millisecond) // Give goroutines time
	wg2.Wait()

	// Mistake 3: Must pass WaitGroup as pointer
	fmt.Println("\n   Correct: Pass WaitGroup by pointer to functions")
	var wg3 sync.WaitGroup

	wg3.Add(1)
	processTask(1, &wg3) // CORRECT: Pass pointer

	wg3.Wait()
	fmt.Println("   Done")
}

func processTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("   Processing task %d\n", id)
}
