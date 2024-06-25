package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)

		if i == 4 {
			for j := 0; j <= 10; j++ {

				if j == 5 {
					character := "САШАРВО"
					for i, char := range character {
						fmt.Printf("character %U %q starts at byte position %d\n", char, char, i)
					}

					continue
				}

				fmt.Println("Nilai j =", j)
			}
		}
	}
}