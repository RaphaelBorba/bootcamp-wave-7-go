package main

import (
	"errors"
	"fmt"
)

type SalaryError1 struct{}

func (SalaryError1) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}

type SalaryError2 struct{}

func (SalaryError2) Error() string {
	return "Error: salary is less than 10000"
}

var (
	ErrSalaryTooLow      = SalaryError2{}
	ErrSalaryTooLowNew   = errors.New("Error: salary is less than 10000")
	ErrHoursWorkedTooLow = errors.New("Error: the worker cannot have worked less than 80 hours per month")
)

func calculateSalary(hours float64, rate float64) (float64, error) {
	if hours < 80 {
		return 0, ErrHoursWorkedTooLow
	}
	salary := hours * rate
	if salary >= 150000 {
		salary *= 0.90
	}
	return salary, nil
}

func main() {

	salary1 := 140000
	if salary1 < 150000 {
		err1 := SalaryError1{}
		fmt.Println(err1)
	} else {
		fmt.Println("Must pay tax")
	}

	salary2 := 9000
	if salary2 <= 10000 {
		err2 := ErrSalaryTooLow
		if errors.Is(err2, ErrSalaryTooLow) {
			fmt.Println(err2)
		}
	}

	salary3 := 8000
	var err3 error
	if salary3 <= 10000 {
		err3 = ErrSalaryTooLowNew
		if errors.Is(err3, ErrSalaryTooLowNew) {
			fmt.Println(err3)
		}
	}

	salary4 := 130000
	if salary4 < 150000 {
		err4 := fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary4)
		fmt.Println(err4)
	}

	h1, r1 := 100.0, 2000.0
	sal5, err5 := calculateSalary(h1, r1)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Printf("Salário calculado: R$ %.2f\n", sal5)
	}

	h2, r2 := 70.0, 1000.0
	sal6, err6 := calculateSalary(h2, r2)
	if err6 != nil {
		fmt.Println(err6)
	} else {
		fmt.Printf("Salário calculado: R$ %.2f\n", sal6)
	}
}
