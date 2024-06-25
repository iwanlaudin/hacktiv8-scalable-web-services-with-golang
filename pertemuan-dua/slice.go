package main

import "fmt"

func main() {

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)
	daySlice := days[5:]
	daySlice[0] = "Sabtu Baru"
	daySlice[1] = "Minggu Baru"

	fmt.Println(daySlice)

	daySlice1 := append(daySlice, "Waktu Ahad")
	fmt.Println(daySlice1)
	daySlice1[1] = "Bukan Minggu"
	fmt.Println(daySlice1)

	// membuat make slice
	newSlice := make([]string, 2, 5) // len slice = 2 dengan capacity =5
	newSlice[0] = "Dody"
	newSlice[1] = "Kurniawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// cara melakukan copy pada slice
	// pastikan ukuran slicenya sama, jika tidak maka akan terpotong
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	// array vs slice
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
