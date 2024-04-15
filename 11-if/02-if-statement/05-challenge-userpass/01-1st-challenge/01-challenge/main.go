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
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

const usage = `
User/Password Storage
--------------
Used to store user/password and is protected

Usage:
go run main.go [username] [password]`

type User struct {
	username string
	password string
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(strings.TrimSpace(usage))

		return
	}
	args := os.Args

	storedUser := &User{username: "jack", password: "1888"}

	username, password := args[1], args[2]

	user := &User{username: username, password: password}

	if user.username == storedUser.username && user.password != storedUser.password {
		fmt.Printf("Invalid password for \"%v\"\n", user.username)

		return
	} else if user.username != storedUser.username {
		fmt.Printf("Access denied for \"%v\"\n", user.username)

		return
	}

	fmt.Printf("Access granted to \"%v\"", user.username)
}
