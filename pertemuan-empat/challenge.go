package main

import "fmt"

func main() {

	liriks := []string{
		"Anak ayam turunlah",
		"Mati satu tinggallah",
		"Mati satu tinggallah induknya",
	}

	c := map[bool]int{true: 1, false: 0}[5 > 4]
	_ = c

	decrement := func(qty int) {

		for i := qty; i > 0; i-- {
			idx := 0
			fmt.Println(liriks[idx], i)

			if i <= 1 {
				fmt.Println(liriks[0], i)
				fmt.Println(liriks[2])
			}
		}
	}

	decrement(18)
}
