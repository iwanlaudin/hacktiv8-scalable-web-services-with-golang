package main

import "fmt"

func main() {

	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}


	// cara mengakses isi array multidimensi
	for i, arrs := range balances {
		for _, value := range arrs {
			fmt.Println("Array ke", i, "value =", value)
		}
	}
}