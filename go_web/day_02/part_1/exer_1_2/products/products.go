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

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type Repository struct {
	products     []Product
	nextID       int
	codeValueSet map[string]struct{}
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

	set := make(map[string]struct{}, len(products))
	maxID := 0
	for _, p := range products {
		set[p.CodeValue] = struct{}{}
		if p.Id > maxID {
			maxID = p.Id
		}
	}

	repo := &Repository{
		products:     products,
		nextID:       maxID + 1,
		codeValueSet: set,
	}
	return repo, nil
}

func (r *Repository) CreateProduct(body CreateProductRequest) (*Product, error) {

	if _, exists := r.codeValueSet[body.CodeValue]; exists {
		return nil, fmt.Errorf("CodeValue %q já está em uso", body.CodeValue)
	}

	newProd := Product{
		Id:          r.nextID,
		Name:        body.Name,
		Quantity:    body.Quantity,
		CodeValue:   body.CodeValue,
		IsPublished: body.IsPublished,
		Expiration:  body.Expiration,
		Price:       body.Price,
	}

	r.products = append(r.products, newProd)
	r.nextID++
	r.codeValueSet[newProd.CodeValue] = struct{}{}
	return &newProd, nil

}

func (r *Repository) GetProductById(id int) (*Product, error) {
	for _, prd := range r.products {
		if prd.Id == id {
			return &prd, nil
		}
	}
	return nil, fmt.Errorf("Nenhum produto encontrado com o id: %d", id)
}
