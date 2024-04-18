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
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//
//	Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//
//	Let's suppose that the user runs it like this:
//
//	  go run main.go 1 10
//
//	Then it should print:
//
//	  2 + 4 + 6 + 8 + 10 = 30
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
	var sum int

	if len(args) < 3 {
		fmt.Println("Please provide 2 arguments, min and max")
		return
	}

	min := convertToInt(args[1], "min")
	max := convertToInt(args[2], "max")

	if min == -1 || max == -1 {
		return
	}

	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}

		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}

	fmt.Printf(" = %d\n", sum)
}
