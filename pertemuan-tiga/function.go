package main

import "fmt"

func main() {

	sayHello()
	greet("Dody", 21)
}

// function tanpa parameter
func sayHello() {
	fmt.Println("Hello Dody!")
}

// function dengan parameter
func greet(name string, age int) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
}