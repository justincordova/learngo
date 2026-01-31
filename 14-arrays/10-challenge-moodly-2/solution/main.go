package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("[your name] [positive|negative]")
		return
	}

	name, mood := args[0], args[1]

	moods := [...][3]string{
		{"happy 😀", "good 👍", "awesome 😎"},
		{"sad 😞", "bad 👎", "terrible 😩"},
	}

	n := rand.IntN(len(moods[0]))

	var mi int
	if mood != "positive" {
		mi = 1
	}
	fmt.Printf("%s feels %s\n", name, moods[mi][n])
}
