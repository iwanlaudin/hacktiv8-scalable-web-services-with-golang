package main

import "fmt"

func main() {

	names := []string{"Joko", "Dody", "Erick", "Indra"}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	person := make(map[string]any)
	person["name"] = "Dody"
	person["title"] = "Software Engineer"
	person["age"] = 21

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}