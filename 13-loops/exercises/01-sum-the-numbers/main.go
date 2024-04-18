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
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//
//	Sum: 55
//
// ---------------------------------------------------------

// Pair type and sortByHeight is just testing
type Pair struct {
	name   string
	height int
}

func sortByHeight(names []string, heights []int) []string {
	var pairs []Pair

	for i := 0; i < len(names); i++ {
		pairs = append(pairs, Pair{name: names[i], height: heights[i]})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].height > pairs[j].height
	})

	sortedNames := make([]string, len(names))
	for i, pair := range pairs {
		sortedNames[i] = pair.name
	}

	return sortedNames
}

func main() {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}

	names := []string{"Yap", "Wei", "Chun"}
	heights := []int{170, 165, 180}

	t := sortByHeight(names, heights)
	fmt.Println(t)

	fmt.Println("Sum:", sum)
}
