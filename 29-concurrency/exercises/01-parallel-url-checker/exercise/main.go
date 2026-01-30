// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Parallel URL Checker
//
//  Create a program that checks multiple URLs concurrently
//  and reports their status (reachable or unreachable).
//
//  1- Define a list of URLs to check (at least 5 URLs)
//     Include a mix of valid and invalid URLs
//
//  2- Create a Result struct to hold:
//     - URL (string)
//     - Status (string: "reachable" or "unreachable")
//     - Error (if any)
//
//  3- Create a function checkURL(url string) Result that:
//     - Makes an HTTP GET request to the URL
//     - Returns a Result with the appropriate status
//     - Use http.Get() with a timeout (use context.WithTimeout)
//
//  4- In main:
//     - Launch a goroutine for each URL check
//     - Use a channel to collect results
//     - Use a WaitGroup to wait for all checks to complete
//     - Print the results as they come in
//
//  5- Bonus: Add timing to show how much faster concurrent
//     checking is compared to sequential checking
//
//
// EXPECTED OUTPUT (timing will vary):
//
//  Checking 5 URLs concurrently...
//
//  https://www.google.com: reachable
//  https://www.github.com: reachable
//  https://invalid-url-xyz-123.com: unreachable (Get "https://invalid-url-xyz-123.com": dial tcp: lookup invalid-url-xyz-123.com: no such host)
//  https://www.golang.org: reachable
//  https://www.example.com: reachable
//
//  Completed in 1.234s
//
// ---------------------------------------------------------

package main

func main() {
	// TODO: Implement the parallel URL checker
}
