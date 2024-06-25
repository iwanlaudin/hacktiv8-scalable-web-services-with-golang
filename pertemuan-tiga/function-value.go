package main

import "fmt"

func main() {

	sayGoodBy := getGoodBy
	fmt.Println(sayGoodBy("Dody"))

	getNameFilter("Dody", spamFilter)
	getNameFilter("Anjing", spamFilter)

	getEmailFilter("dody@gmail.com", spamFilter)
}

// function variadic
func getGoodBy(name string) string {
	return "Good bay " + name
}

// function as parameter
func getNameFilter(name string, filter func(string)string) {
	filtered := filter(name)
	fmt.Println("Hello", filtered)
}

// function Type Declarations
type Filter func(string)string

func getEmailFilter(email string, filter Filter) {
	fmt.Println("Email", filter(email))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}