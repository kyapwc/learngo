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
// EXERCISE: Age Seasons
//
//  Let's start simple. Print the expected outputs,
//  depending on the age variable.
//
// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up
// ---------------------------------------------------------

var usage = `
Age Seasons
---------------------------------------------------------
Please provide an age variable.

Usage: go run main.go [age]
`

func main() {
	// Change this accordingly to produce the expected outputs
	args := os.Args

	if len(args) < 2 {
		fmt.Println(usage)
		return
	}

	age, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Println("Failed to convert age, with err: ", err)

		return
	}

	// Type your if statement here.
	if age > 10 && age < 20 {
		fmt.Println("Young blood")
	} else if age > 20 && age < 30 {
		fmt.Println("Adulthood")
	} else if age > 30 && age < 60 {
		fmt.Println("Getting wiser")
	} else if age > 60 {
		fmt.Println("Getting older")
	} else {
		fmt.Println("Booting up")
	}
}
