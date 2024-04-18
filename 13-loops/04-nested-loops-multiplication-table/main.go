// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

// EXERCISE: Get the `max` from the user
//           And create the table according to that.

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide the maximum")
		return
	}

	max, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Printf("Max (%v) provided is not a number", args[1])
		return
	}

	// print the header
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
