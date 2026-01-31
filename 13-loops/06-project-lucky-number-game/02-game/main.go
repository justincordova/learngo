package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?
`
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		// fmt.Println("Pick a number.")
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.IntN(guess + 1)

		if n == guess {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
