package product

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/raphaelBorba/api-chi/internal/domain"
)

type Repository struct {
	products     []domain.Product
	nextID       int
	codeValueSet map[string]struct{}
}

func NewRepository(jsonPath string) *Repository {
	file, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var products []domain.Product

	if err := decoder.Decode(&products); err != nil {
		return nil
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
	return repo
}

func (r *Repository) Create(body domain.CreateProductRequest) (*domain.Product, error) {

	if _, exists := r.codeValueSet[body.CodeValue]; exists {
		return nil, fmt.Errorf("CodeValue %q já está em uso", body.CodeValue)
	}

	newProd := domain.Product{
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

func (r *Repository) GetById(id int) (*domain.Product, error) {
	for _, prd := range r.products {
		if prd.Id == id {
			return &prd, nil
		}
	}
	return nil, fmt.Errorf("Nenhum produto encontrado com o id: %d", id)
}
