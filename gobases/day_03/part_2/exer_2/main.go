package main

import (
	"fmt"
	"inter_prod/product"
)

func main() {
	p1 := product.NewProduct("small", 100)
	p2 := product.NewProduct("medium", 200)
	p3 := product.NewProduct("large", 300)

	fmt.Printf("Preço total do produto pequeno: R$ %.2f\n", p1.Price())
	fmt.Printf("Preço total do produto médio: R$ %.2f\n", p2.Price())
	fmt.Printf("Preço total do produto grande: R$ %.2f\n", p3.Price())
}
