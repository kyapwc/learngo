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
)

// ---------------------------------------------------------
// CHALLENGE #2
//  Add one more user to the PassMe program below.
//
// EXAMPLE USERS
//  username: jack
//  password: 1888
//
//  username: inanc
//  password: 1879
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 1888
//    Access granted to "jack".
//
//  go run main.go inanc 1879
//    Access granted to "inanc".
//
//  go run main.go jack 1879
//    Invalid password for "jack".
//
//  go run main.go inanc 1888
//    Invalid password for "inanc".
// ---------------------------------------------------------

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
	user1    = "jack"
	pass1    = "1888"
	user2    = "inanc"
	pass2    = "1879"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	user, password := args[1], args[2]

	if user != user1 && user != user2 {
		fmt.Printf(errUser, user)
	} else if user == user1 && password != pass1 {
		fmt.Printf(errPwd, user)
	} else if user == user2 && password != pass2 {
		fmt.Printf(errPwd, user)
	} else {
		fmt.Printf(accessOK, user)
	}
}
