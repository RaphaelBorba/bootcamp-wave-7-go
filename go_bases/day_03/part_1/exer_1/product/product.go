package product

import (
	"errors"
	"fmt"
)

type ProductSA struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var products = []ProductSA{}

func (p ProductSA) Save() {
	products = append(products, p)
}

func (p ProductSA) GetAll() {
	for _, prod := range products {
		fmt.Printf("ID: %d | Nome: %s | Preço: R$ %.2f | Descrição: %s | Categoria: %s\n",
			prod.ID, prod.Name, prod.Price, prod.Description, prod.Category)
	}
}

func GetById(id int) (ProductSA, error) {

	for _, prd := range products {
		if prd.ID == id {
			return prd, nil
		}
	}
	return ProductSA{}, errors.New("produto não encontrado")
}
