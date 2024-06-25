package main

import "fmt"

func main() {

	sayHelloTo("Dody", "Kurniawan")

	result := getHello("Dody Kurniawan")
	fmt.Println(result)

	firtsName, _, lastName := getName("Joko", " ", "Irawan")
	fmt.Println(firtsName, lastName)

	fName, lName := getNames()
	fmt.Println(fName, lName)
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function dengan return value
func getHello(name string) string {
	return "Hello " + name
}

// function dengan multiple return values
func getName(firstName string, middleName string, lastName string) (string, string, string) {
	return firstName, middleName, lastName
}

// named return value
func getNames() (firstName string, lastName string) {

	firstName = "Erick"
	lastName = "Kurniawan"
	return
}