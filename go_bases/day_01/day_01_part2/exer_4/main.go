package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Employees with more than 21 years old are:", returnEmployeesInfo(employees))
	fmt.Println(employees)
}

func returnEmployeesInfo(employees map[string]int) int {

	emp_older_than_21 := 0
	for key, val := range employees {
		if key == "Benjamin" {
			fmt.Println("Idade de Benjamin é:", val)
		}

		if val > 21 {
			emp_older_than_21++
		}
	}

	employees["Frederico"] = 25

	delete(employees, "Pedro")

	return emp_older_than_21
}
