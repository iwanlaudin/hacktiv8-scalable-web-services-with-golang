package main

import "fmt"

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)

	recursive := factorialFunction(5)
	fmt.Println(recursive)
}

// menggunakan looping
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// menggunakan recursive function
func factorialFunction(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialFunction(value-1)
	}
}