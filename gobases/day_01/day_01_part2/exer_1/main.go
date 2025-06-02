package main

import "fmt"

func main() {

	var word string = "Escola"

	for _, val := range word {
		fmt.Printf("%c\n", val)
	}
	fmt.Println(len(word))
}
