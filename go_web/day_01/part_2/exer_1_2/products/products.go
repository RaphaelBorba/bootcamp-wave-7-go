package products

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type Repository struct {
	products []Product
}

func NewRepository(jsonPath string) (*Repository, error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return nil, fmt.Errorf("could not open JSON file %q: %w", jsonPath, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var products []Product

	if err := decoder.Decode(&products); err != nil {
		return nil, fmt.Errorf("Erro ao decodificar JSON %q: %w", jsonPath, err)
	}

	return &Repository{products: products}, nil
}

func (r *Repository) GetProducts() []Product {
	return r.products
}

func (r *Repository) GetProductById(id int) (*Product, error) {
	for _, prd := range r.products {
		if prd.Id == id {
			return &prd, nil
		}
	}
	return nil, fmt.Errorf("Nenhum produto encontrado com o id: %d", id)
}
