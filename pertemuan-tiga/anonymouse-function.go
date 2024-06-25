package main

import "fmt"

func main() {

	// anonymouse function #1
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("admin", blacklist)

	/* 
	anonymouse function #2
	dengan memasukan function langsung pada parameter dalam pemanggilan sebuah function
	yang membutuhkan function paramter
	*/
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}