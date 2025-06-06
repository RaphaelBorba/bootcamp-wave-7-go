package main

import (
	"struct_emp/employee"
	"struct_emp/person"
)

func main() {
	emp1 := employee.Employee{
		ID:       4,
		Position: "Software Developer",
		Person: person.Person{
			ID:          2,
			Name:        "Fulano Ciclano",
			DateOfBirth: "01/01/1990",
		},
	}

	emp1.PrintEmployee()
}
