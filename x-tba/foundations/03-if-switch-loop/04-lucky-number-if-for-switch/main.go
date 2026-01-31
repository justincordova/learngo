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

(Provide -v flag to see the picked numbers.)
`
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var verbose bool
	if args[0] == "-v" {
		verbose = true
	}

	guess, err := strconv.Atoi(args[len(args)-1])
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

		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess {
			switch rand.IntN(3) {
			case 0:
				fmt.Println("🎉  YOU WIN!")
			case 1:
				fmt.Println("🎉  YOU'RE AWESOME!")
			case 2:
				fmt.Println("🎉  PERFECT!")
			}
			return
		}
	}

	// msg, n := "%s Try again?\n", rand.IntN(5)
	// if msg, n := "%s Try again?\n", rand.IntN(5); n <= 2 {
	// 	fmt.Printf(msg, "☠️  YOU LOST...")
	// } else if n < 3 {
	// 	fmt.Printf(msg, "☠️  JUST A BAD LUCK...")
	// } else if n == 4 {
	// 	fmt.Printf(msg, "☠️  TRY NEXT TIME...")
	// }

	// var msg string
	// switch rand.IntN(10) {
	// // more probability
	// case 0, 1, 2, 3, 4, 5:
	// 	msg = "☠️  YOU LOST..."
	// case 6, 7, 8:
	// 	msg = "☠️  JUST A BAD LUCK..."
	// default:
	// 	msg = "☠️  TRY NEXT TIME..."
	// }
	// fmt.Printf("%s Try again?\n", msg)

	var msg string
	switch n := rand.IntN(10); {
	// more probability
	case n <= 5:
		msg = "☠️  YOU LOST..."
	case n <= 8:
		msg = "☠️  JUST A BAD LUCK..."
	default:
		msg = "☠️  TRY NEXT TIME..."
	}
	fmt.Printf("%s Try again?\n", msg)
}
