package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// rand.Seed(10)
	// rand.Seed(100)

	// t := time.Now()
	// rand.Seed(t.UnixNano())

	// ^-- same:

	guess := 10

	for n := 0; n != guess; {
		n = rand.IntN(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
