package main

import "fmt"

func main() {

	fmt.Println("Salário pós taxas:", returnSalaryAfterTax(150001))

}

func returnSalaryAfterTax(salary float32) float32 {

	if salary > 150000 {
		salary *= 0.73
		return salary
	} else if salary > 50000 {
		return salary * 0.87
	}
	return salary
}
