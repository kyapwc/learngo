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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Richter Scale #2
//
//  Repeat the previous exercise.
//
//  But, this time, get the description and print the
//  corresponding richter scale.
//
//  See the expected outputs.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 to less than 3.0         very minor
// 3.0 to less than 4.0         minor
// 4.0 to less than 5.0         light
// 5.0 to less than 6.0         moderate
// 6.0 to less than 7.0         strong
// 7.0 to less than 8.0         major
// 8.0 to less than 10.0        great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//   Tell me the magnitude of the earthquake in human terms.
//
//  go run main.go alien
//   alien's richter scale is unknown
//
//  go run main.go micro
//   micro's richter scale is less than 2.0
//
//  go run main.go "very minor"
//   very minor's richter scale is 2 - 2.9
//
//  go run main.go minor
//   minor's richter scale is 3 - 3.9
//
//  go run main.go light
//   light's richter scale is 4 - 4.9
//
//  go run main.go moderate
//   moderate's richter scale is 5 - 5.9
//
//  go run main.go strong
//   strong's richter scale is 6 - 6.9
//
//  go run main.go major
//   major's richter scale is 7 - 7.9
//
//  go run main.go great
//   great's richter scale is 8 - 9.9
//
//  go run main.go massive
//   massive's richter scale is 10+
// ---------------------------------------------------------

// func main() {
// 	args := os.Args
//
// 	if len(args) < 2 {
// 		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
// 		return
// 	}
//
// 	// less resource intensive is to use a `map`
// 	// but its hard for error handling and the list will keep getting larger
// 	// it would do better to use a switch case
// 	rictherScale := map[string]string{
// 		"micro":      "micro's richter scale is less than 2.0",
// 		"very minor": "very minor's richter scale is 2 - 2.9",
// 		"minor":      "minor's richter scale is 3 - 3.9",
// 		"light":      "light's richter scale is 4 - 4.9",
// 		"moderate":   "moderate's richter scale is 5 - 5.9",
// 		"strong":     "strong's richter scale is 6 - 6.9",
// 		"major":      "major's richter scale is 7 - 7.9",
// 		"great":      "great's richter scale is 8 - 9.9",
// 		"massive":    "massive's richter scale is 10+",
// 	}
//
// 	value, found := rictherScale[strings.ToLower(args[1])]
//
// 	if !found {
// 		fmt.Printf("%v's richter scale is unknown", args[1])
// 		return
// 	}
//
// 	fmt.Println(value)
// }

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	value := strings.ToLower(args[1])

	if len(value) == 0 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	var desc string

	switch value {
	case "micro":
		desc = "less than 2.0"
	case "very minor":
		desc = "2 - 2.9"
	case "minor":
		desc = "3 - 3.9"
	case "light":
		desc = "4 - 4.9"
	case "moderate":
		desc = "5 - 5.9"
	case "strong":
		desc = "6 - 6.9"
	case "major":
		desc = "5 - 5.9"
	case "great":
		desc = "8 - 9.9"
	case "massive":
		desc = "10+"
	default:
		desc = "unknown"
	}

	fmt.Printf("%v's richter scale is %v", args[1], desc)
}
