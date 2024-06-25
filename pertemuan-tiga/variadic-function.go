package main

import "fmt"

func main() {

	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Total =", total)

	number := []int{11, 12, 13, 14, 15}
	total = sumAll(number...)
	fmt.Println("Total =", total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}