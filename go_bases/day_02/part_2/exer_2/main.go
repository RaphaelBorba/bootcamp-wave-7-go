package main

import (
	"fmt"
)

func average(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	var sum int
	for _, v := range nums {
		sum += v
	}
	return float64(sum) / float64(len(nums))
}

func main() {
	a := []int{5, 4, 3, 6, 7, 8}
	fmt.Println("Média de", a, "=", average(a))

	b := []int{}
	fmt.Println("Média de slice vazio =", average(b))
}
