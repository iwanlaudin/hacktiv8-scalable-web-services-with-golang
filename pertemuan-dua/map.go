package main

import (
	"fmt"
)

func main() {

	// function yang dapat digunakan dalam map
	// len(map) -> untuk mendapatkan jumlah data map
	// map[key] -> mengambil data di map dengan key
	// map[key] = value -> mengubah data di map dengan key
	// make(map[TypeKey]TypeValue) -> membuat map baru
	// delete(map, key) -> menghapus data di map dengan key

	person := map[string]string{
		"name":    "Dody",
		"address": "Jakarta Selatan",
	}

	// menambah data map
	person["title"] = "Software Engineer"

	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Dodyd Kurniawan"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}