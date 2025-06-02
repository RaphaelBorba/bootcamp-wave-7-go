package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Salário do empregado será: ", calcSalary(1000, "C"))

}

func calcSalary(workMinutes int, category string) float64 {
	workHours := float64(workMinutes) / 60.0
	category = strings.ToUpper(category)

	switch category {
	case "C":
		return workHours * 1000
	case "B":
		base := workHours * 1500
		return base + base*0.2
	case "A":
		base := workHours * 3000
		return base + base*0.5
	default:
		fmt.Println("Categoria inválida. Use A, B ou C.")
		return 0
	}
}
