package main

import (
	"fmt"
	"os"
	"strconv"
)

type Persons struct {
	Name, Address, Job, Reason string
}

func datas() map[int]Persons {
	person := make(map[int]Persons, 2)
	person[1] = Persons{
		Name:    "Dody Kurniawan",
		Address: "Kota Bandung",
		Job:     "Software Engineer",
		Reason:  "Memiliki kemampuan konkurensi yang baik",
	}
	person[2] = Persons{
		Name:    "Eryck Kurniawan",
		Address: "Jakarta Selatan",
		Job:     "Web Designer",
		Reason:  "Performa yang tinggi dan mudah dipelajari",
	}
	person[3] = Persons{
		Name:    "Iwan La Udin",
		Address: "Ternate, Maluku Utara",
		Job:     "Software Engineer",
		Reason:  "Performa yang tinggi dan mudah dipelajari",
	}

	return person
}

func getPersonByKey(key int, person map[int]Persons) {
	fmt.Printf("Nama: %s\n", person[key].Name)
	fmt.Printf("Alamat: %s\n", person[key].Address)
	fmt.Printf("Pekerjaan: %s\n", person[key].Job)
	fmt.Printf("Alasan Memilih Golang: %s\n", person[key].Reason)
}

func main() {

	// person := make(map[int]Persons, 2)

	// person[1] = Persons{
	// 	Name:    "Dody Kurniawan",
	// 	Address: "Kota Bandung",
	// 	Job:     "Software Engineer",
	// 	Reason:  "Memiliki kemampuan konkurensi yang baik",
	// }
	// person[2] = Persons{
	// 	Name:    "Eryck Kurniawan",
	// 	Address: "Jakarta Selatan",
	// 	Job:     "Web Designer",
	// 	Reason:  "Performa yang tinggi dan mudah dipelajari",
	// }
	// person[3] = Persons{
	// 	Name:    "Iwan La Udin",
	// 	Address: "Ternate, Maluku Utara",
	// 	Job:     "Software Engineer",
	// 	Reason:  "Performa yang tinggi dan mudah dipelajari",
	// }

	// getData := func (number int) {
	// 	fmt.Printf("Nama: %s\n", person[number].Name)
	// 	fmt.Printf("Alamat: %s\n", person[number].Address)
	// 	fmt.Printf("Pekerjaan: %s\n", person[number].Job)
	// 	fmt.Printf("Alasan Memilih Golang: %s\n", person[number].Reason)
	// }

	// _ = getData


	if len(os.Args) != 2{
		fmt.Println("Usage: go run namafile.go <nomor>")
		os.Exit(1)
	}

	number, err := strconv.Atoi(os.Args[1])
	if number <= 0 || err != nil{
		fmt.Println("Nomor harus berupa angka dan tidak boleh kurang atau sama dengan 0!")
		os.Exit(2)
	}

	datas := datas()

	if number > 0 && number <= len(datas) {
		getPersonByKey(number, datas)
		// getData(number)
	} else {
		fmt.Println("Data absen tidak tersedia")
	}
}
