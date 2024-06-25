package main

import "fmt"

func main() {

	// cara pertama
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	// cara kedua
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
}