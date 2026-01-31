// ---------------------------------------------------------
// EXERCISE: Worker with Cancellation
//
//  Create a program with multiple workers that respect
//  context cancellation for graceful shutdown.
//
//  1- Create a worker function that:
//     - Accepts a context, worker ID, and a jobs channel
//     - Processes jobs from the channel
//     - Respects context cancellation using select
//     - Prints when it starts a job and when it completes
//     - Prints when it shuts down due to cancellation
//
//  2- Each job should:
//     - Simulate work by sleeping for a random duration (100-500ms)
//     - Check the context in the middle of processing
//     - Stop immediately if context is cancelled
//
//  3- In main:
//     - Create a cancellable context
//     - Start 3 worker goroutines
//     - Send 10 jobs to the workers
//     - After 1 second, cancel the context
//     - Wait for all workers to shut down gracefully
//
//  4- Use a WaitGroup to track when all workers finish
//
//
// EXPECTED OUTPUT (order may vary):
//
//  Starting 3 workers...
//  Worker 1: Processing job 1
//  Worker 2: Processing job 2
//  Worker 3: Processing job 3
//  Worker 1: Completed job 1
//  Worker 1: Processing job 4
//  Worker 2: Completed job 2
//  Worker 2: Processing job 5
//  Worker 3: Completed job 3
//  Worker 3: Processing job 6
//
//  Cancelling context...
//  Worker 1: Shutting down (context cancelled)
//  Worker 2: Shutting down (context cancelled)
//  Worker 3: Shutting down (context cancelled)
//
//  All workers stopped. Jobs processed: 6/10
//
// ---------------------------------------------------------

package main

func main() {
	// TODO: Implement the worker pool with cancellation
}
