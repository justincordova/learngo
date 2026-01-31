// ---------------------------------------------------------
// EXERCISE: Rate Limiter
//
//  Create a program that processes tasks with rate limiting
//  to avoid overwhelming a service or API.
//
//  1- Create a list of 10 "API requests" (just strings like "Request 1", "Request 2", etc.)
//
//  2- Implement a rate limiter that:
//     - Allows only 2 requests per second
//     - Uses time.Ticker to control the rate
//     - Processes requests concurrently but respects the rate limit
//
//  3- Create a function processRequest(id string) that:
//     - Simulates processing by sleeping for 100ms
//     - Prints when a request starts and completes
//
//  4- In main:
//     - Create a ticker that ticks every 500ms (2 per second)
//     - Use a channel to send requests
//     - Use a goroutine to process requests from the channel
//     - Print timing information to show rate limiting is working
//
//  5- Bonus: Add a burst capacity that allows 3 requests
//     to be processed immediately, then rate limits
//
//
// EXPECTED OUTPUT (timing will vary):
//
//  Starting rate-limited request processor...
//  Rate limit: 2 requests per second
//
//  [0.000s] Processing Request 1
//  [0.000s] Processing Request 2
//  [0.100s] Completed Request 1
//  [0.100s] Completed Request 2
//  [0.500s] Processing Request 3
//  [0.500s] Processing Request 4
//  [0.600s] Completed Request 3
//  [0.600s] Completed Request 4
//  [1.000s] Processing Request 5
//  [1.000s] Processing Request 6
//  ...
//
// ---------------------------------------------------------

package main

func main() {
	// TODO: Implement the rate limiter
}
