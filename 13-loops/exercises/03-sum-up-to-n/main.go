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

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//
//	For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//
//	Let's suppose that the user runs it like this:
//
//	  go run main.go 1 10
//
//	Then it should print:
//
//	  1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
//
// ---------------------------------------------------------
func convertToInt(value string, valueType string) int {
	output, err := strconv.Atoi(value)

	if err != nil {
		fmt.Printf("%q not a valid number", valueType)
		return -1
	}

	return output
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Please provide 2 arguments, min and max")
		return
	}

	min := convertToInt(args[1], "min")
	max := convertToInt(args[2], "max")

	if min == -1 || max == -1 {
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

// Interviewer name: Sohan Gunasekera, 10 years in mudah, engineer manager
//
