package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	for {
		var c string

		switch rand.IntN(4) {
		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		case 3:
			c = "-"
		}
		fmt.Printf("\r%s Please Wait. Processing....", c)
		time.Sleep(time.Millisecond * 250)
	}
}
