package main

import (
	"fmt"
)

func main() {

	tokenize := func (str string) {
		charCounter := make(map[string]int)

		for _, char := range str {
			fmt.Println(string(char))

			if val, key := charCounter[string(char)]; key {
				charCounter[string(char)] = val + 1
			} else {
				charCounter[string(char)] = 1
			}
		}

		fmt.Println(charCounter)
	}

	tokenize("selamat malam")
}