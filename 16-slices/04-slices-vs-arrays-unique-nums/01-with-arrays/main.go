package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	const max = 5
	var uniques [max]int

	// It's harder to make a program dynamic using arrays
	// max, _ := strconv.Atoi(os.Args[1])
	// var uniques [10]int

loop:
	for found := 0; found < max; {
		n := rand.IntN(max) + 1
		fmt.Print(n, " ")

		for _, u := range uniques {
			if u == n {
				continue loop
			}
		}

		uniques[found] = n
		found++
	}

	fmt.Println("\n\nuniques:", uniques)
}
