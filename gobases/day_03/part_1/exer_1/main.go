package main

import (
	"fmt"
	"stuct/product"
)

func main() {

	product1 := product.ProductSA{ID: 1, Name: "Camisa", Price: 59.90, Description: "Camisa de algodão", Category: "Vestuário"}
	product2 := product.ProductSA{ID: 2, Name: "Calça Jeans", Price: 129.90, Description: "Calça jeans masculina", Category: "Vestuário"}

	product1.Save()

	product1.GetAll()

	product2.Save()

	product2.GetAll()

	prod, err := product.GetById(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Produto encontrado: ID: %d | Nome: %s | Preço: R$ %.2f | Categoria: %s\n",
			prod.ID, prod.Name, prod.Price, prod.Category)
	}

}
