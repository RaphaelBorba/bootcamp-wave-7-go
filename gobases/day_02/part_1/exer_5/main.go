package main

import (
	"errors"
	"fmt"
)

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func animal(ani string) (func(qtd int64) float64, error) {

	switch ani {
	case dog:
		return func(qtd int64) float64 {
			return float64(qtd) * 10
		}, nil
	case cat:
		return func(qtd int64) float64 {
			return float64(qtd) * 5
		}, nil
	case hamster:
		return func(qtd int64) float64 {
			return float64(qtd) * 0.25
		}, nil
	case spider:
		return func(qtd int64) float64 {
			return float64(qtd) * 0.15
		}, nil
	default:
		return nil, errors.New("Animal not registred!")
	}
}

func main() {

	var (
		dogFunc, msgDog         = animal(dog)
		catFunc, msgCat         = animal(cat)
		hamsterFunc, msgHamster = animal(hamster)
		spiderFunc, msgSpider   = animal(spider)
	)

	fmt.Println(msgDog, msgCat, msgHamster, msgSpider)

	totalFood := 0.0

	totalFood += dogFunc(1)
	totalFood += catFunc(1)
	totalFood += hamsterFunc(1)
	totalFood += spiderFunc(1)

	fmt.Println(totalFood)

}
