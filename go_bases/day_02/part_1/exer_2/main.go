package main

import "fmt"

func main() {

	fmt.Println("A média do aluno é: ", calcAvg(10, 20, 20))

}

func calcAvg(grades ...uint) float32 {

	var totalSum uint = 0

	for _, grade := range grades {
		totalSum += grade
	}

	return float32(totalSum) / float32(len(grades))
}
