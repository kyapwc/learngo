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
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//	Print the number of days in a given month.
//
// RESTRICTIONS
//
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
// EXPECTED OUTPUT
//
//	-----------------------------------------
//	Your solution should not accept invalid months
//	-----------------------------------------
//	go run main.go
//	  Give me a month name
//
//	go run main.go sheep
//	  "sheep" is not a month.
//
//	-----------------------------------------
//	Your solution should handle the leap years
//	-----------------------------------------
//	go run main.go january
//	  "january" has 31 days.
//
//	go run main.go february
//	  "february" has 28 days.
//
//	go run main.go march
//	  "march" has 31 days.
//
//	go run main.go april
//	  "april" has 30 days.
//
//	go run main.go may
//	  "may" has 31 days.
//
//	go run main.go june
//	  "june" has 30 days.
//
//	go run main.go july
//	  "july" has 31 days.
//
//	go run main.go august
//	  "august" has 31 days.
//
//	go run main.go september
//	  "september" has 30 days.
//
//	go run main.go october
//	  "october" has 31 days.
//
//	go run main.go november
//	  "november" has 30 days.
//
//	go run main.go december
//	  "december" has 31 days.
//
//	-----------------------------------------
//	Your solution should be case insensitive
//	-----------------------------------------
//	go run main.go DECEMBER
//	  "DECEMBER" has 31 days.
//
//	go run main.go dEcEmBeR
//	  "dEcEmBeR" has 31 days.
//
// ---------------------------------------------------------
type MonthAndDays struct {
	name string
	days int32
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Give me a month name")
		return
	}

	year, _, _ := time.Now().Date()

	months := map[string]int32{
		"january":   31,
		"february":  28,
		"march":     31,
		"april":     30,
		"may":       31,
		"june":      30,
		"july":      31,
		"august":    31,
		"september": 30,
		"october":   31,
		"november":  30,
		"december":  31,
	}

	value, found := months[strings.ToLower(args[1])]

	if !found {
		fmt.Printf("\"%v\" is not a month", args[1])
		return
	}

	isLeapYear := year%4 == 0

	if isLeapYear && strings.ToLower(args[1]) == "february" {
		value += 1
	}

	fmt.Printf("\"%v\" has %v days", args[1], value)
}
