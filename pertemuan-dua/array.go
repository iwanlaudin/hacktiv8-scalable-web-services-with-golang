package main

import "fmt"

func main() {

	// membuat array secara manual
	var names [3]string

	names[0] = "Dody"
	names[1] = "Joko"
	names[2] = "Erick"

	fmt.Println(names)

	// cara mengakses isi dari array
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array secara langsung
	values := [5]int{90, 80, 75, 70, 65}

	// cara memodifikasi isi dari array
	values[0] = 100

	fmt.Println(values)

	// cara mendapatkan panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))
}