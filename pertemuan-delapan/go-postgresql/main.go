package main

import (
	"fmt"
	"pertemuan-delapan/go-postgresql/config"
	"pertemuan-delapan/go-postgresql/services"
)

func main() {
	db := config.CreateConnection()
	defer db.Close()

	fmt.Printf("Created: %v\n", services.CreateEmployee(db, "Budy", "budi@gmail.com", 21, "Tech Data"))

	// fmt.Printf("services.GetEmployee(db): %v\n", services.GetEmployee(db))

	// fmt.Printf("Status: %v\n", services.UpdateEmployee(db, 3, "Erick", "erick@gmail.com", 23, "Tech Data"))

	// fmt.Printf("Status: %v\n", services.DeleteEmployee(db, 1))
}
