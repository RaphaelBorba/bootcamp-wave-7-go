package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operation(op string) (func(...int) int, error) {
	switch op {
	case minimum:
		return func(nums ...int) int {
			if len(nums) == 0 {
				return 0
			}
			min := nums[0]
			for _, n := range nums {
				if n < min {
					min = n
				}
			}
			return min
		}, nil

	case average:
		return func(nums ...int) int {
			if len(nums) == 0 {
				return 0
			}
			sum := 0
			for _, n := range nums {
				sum += n
			}
			return sum / len(nums)
		}, nil

	case maximum:
		return func(nums ...int) int {
			if len(nums) == 0 {
				return 0
			}
			max := nums[0]
			for _, n := range nums {
				if n > max {
					max = n
				}
			}
			return max
		}, nil

	default:
		return nil, errors.New("Operação inválida. Use: minimum, maximum ou average")
	}
}

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	avgFunc, _ := operation(average)
	maxFunc, _ := operation(maximum)

	min := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	avg := avgFunc(2, 3, 3, 4, 1, 2, 4, 5)
	max := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Minimum:", min)
	fmt.Println("Average:", avg)
	fmt.Println("Maximum:", max)
}
