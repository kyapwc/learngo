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
// EXERCISE: Simplify the Leap Year
//
//  1. Look at the solution of "the previous exercise".
//
//  2. And simplify the code (especially the if statements!).
//
// EXPECTED OUTPUT
//  It's the same as the previous exercise.
// ---------------------------------------------------------

func main() {
	args := os.Args

	if len(args) < 2 || len(args[1]) < 0 {
		fmt.Println("Give me a year number")
		return
	}

	value, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Printf("\"%v\" is not a valid year\n", args[1])
		return
	}

	remainder := value % 4 // 4 as the modulus, since leap year happens every 4 years
	// if year%4 == 0 && (year%100 != 0 || year%400 == 0) {

	if remainder == 0 && (value%100 != 0 || value%400 == 0) {
		fmt.Printf("%v is a leap year.\n", value)
	} else {
		fmt.Printf("%v is not a leap year.\n", value)
	}
}
