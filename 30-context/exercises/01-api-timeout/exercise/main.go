// ---------------------------------------------------------
// EXERCISE: API Client with Timeout
//
//  Create a program that makes HTTP requests with timeouts
//  using context to prevent requests from hanging indefinitely.
//
//  1- Create a function fetchURL(ctx context.Context, url string) that:
//     - Creates an HTTP GET request with the provided context
//     - Returns the response status code and any error
//     - Uses http.NewRequestWithContext to attach the context
//
//  2- In main, create a list of URLs with different response times:
//     - https://httpbin.org/delay/1 (1 second delay)
//     - https://httpbin.org/delay/3 (3 second delay)
//     - https://httpbin.org/delay/5 (5 second delay)
//
//  3- For each URL:
//     - Create a context with a 2-second timeout
//     - Call fetchURL with the context
//     - Print whether the request succeeded or timed out
//     - Don't forget to call the cancel function
//
//  4- Handle both successful requests and timeout errors appropriately
//
//
// EXPECTED OUTPUT:
//
//  Fetching https://httpbin.org/delay/1 with 2s timeout...
//  Success: Status 200 OK
//
//  Fetching https://httpbin.org/delay/3 with 2s timeout...
//  Error: request timed out: context deadline exceeded
//
//  Fetching https://httpbin.org/delay/5 with 2s timeout...
//  Error: request timed out: context deadline exceeded
//
// ---------------------------------------------------------

package main

func main() {
	// TODO: Implement the API client with timeout
}
