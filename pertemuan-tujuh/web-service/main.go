package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var PORT = ":5001"

type Employee struct {
	Id       int
	Name     string
	Age      int
	Devision string
}

var employees = []Employee{
	{Id: 1, Name: "Dody", Age: 23, Devision: "Tech Data"},
	{Id: 2, Name: "Joko", Age: 21, Devision: "Tech Data"},
	{Id: 3, Name: "Erick", Age: 22, Devision: "Tech Data"},
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid method!", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		devision := r.FormValue("devision")

		ageConv, err := strconv.Atoi(age)
		if err != nil {
			http.Error(w, "Invalid Age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			Id:       len(employees) + 1,
			Name:     name,
			Age:      ageConv,
			Devision: devision,
		}

		employees = append(employees, newEmployee)
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/employees", getEmployees)

	fmt.Println("Application is listenning on port", PORT)
	http.ListenAndServe(PORT, nil)
}
