package product

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/raphaelBorba/api-chi/put-patch-delete/internal/domain"
)

type Repository struct {
	products      map[int]domain.Product
	nextID        int
	codeValueToID map[string]int
}

func NewRepository(jsonPath string) (*Repository, error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo %s: %w", jsonPath, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var products []domain.Product
	if err := decoder.Decode(&products); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}
	productMap := make(map[int]domain.Product, len(products))
	codeValueToID := make(map[string]int, len(products))
	maxID := 0
	for _, p := range products {
		productMap[p.Id] = p
		codeValueToID[p.CodeValue] = p.Id
		if p.Id > maxID {
			maxID = p.Id
		}
	}

	repo := &Repository{
		products:      productMap,
		nextID:        maxID + 1,
		codeValueToID: codeValueToID,
	}
	return repo, nil
}

func (r Repository) GetAll() map[int]domain.Product {
	return r.products
}

func (r *Repository) Create(body domain.CreateProductRequest) (*domain.Product, error) {
	if _, exists := r.codeValueToID[body.CodeValue]; exists {
		return nil, fmt.Errorf("code_value %q já está em uso", body.CodeValue)
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

	r.products[newProd.Id] = newProd
	r.codeValueToID[newProd.CodeValue] = newProd.Id
	r.nextID++
	return &newProd, nil
}

func (r Repository) GetByID(id int) (*domain.Product, error) {
	prod := r.products[id]
	if reflect.DeepEqual(prod, domain.Product{}) {
		return nil, fmt.Errorf("nenhum produto encontrado com o Id: %d", id)
	}
	return &prod, nil
}

func (r *Repository) Update(id int, body domain.CreateProductRequest) (*domain.Product, error) {

	old := r.products[id]

	if reflect.DeepEqual(old, domain.Product{}) {
		return nil, fmt.Errorf("nenhum produto encontrado com o Id: %d", id)
	}

	if body.CodeValue != old.CodeValue {
		if _, exists := r.codeValueToID[body.CodeValue]; exists {
			return nil, fmt.Errorf("code_value %q já está em uso", body.CodeValue)
		}
		delete(r.codeValueToID, old.CodeValue)
		r.codeValueToID[body.CodeValue] = id
	}

	r.products[id] = domain.Product{
		Id:          id,
		Name:        body.Name,
		Quantity:    body.Quantity,
		CodeValue:   body.CodeValue,
		IsPublished: body.IsPublished,
		Expiration:  body.Expiration,
		Price:       body.Price,
	}

	updProd := r.products[id]

	return &updProd, nil
}

func (r *Repository) Patch(id int, patchData map[string]any) (*domain.Product, error) {

	prod := r.products[id]

	if reflect.DeepEqual(prod, domain.Product{}) {
		return nil, fmt.Errorf("nenhum produto encontrado com o Id: %d", id)
	}

	if v, ok := patchData["name"].(string); ok {
		prod.Name = v
	}
	if v, ok := patchData["quantity"].(float64); ok {
		prod.Quantity = int(v)
	}
	if v, ok := patchData["code_value"].(string); ok {
		if _, exists := r.codeValueToID[v]; exists && v != prod.CodeValue {
			return nil, fmt.Errorf("code_value %q já está em uso", v)
		}
		delete(r.codeValueToID, prod.CodeValue)
		prod.CodeValue = v
		r.codeValueToID[v] = id
	}
	if v, ok := patchData["is_published"].(bool); ok {
		prod.IsPublished = v
	}
	if v, ok := patchData["expiration"].(string); ok {
		prod.Expiration = v
	}
	if v, ok := patchData["price"].(float64); ok {
		prod.Price = v
	}

	r.products[id] = prod

	return &prod, nil
}

func (r *Repository) Delete(id int) error {

	prod := r.products[id]

	if reflect.DeepEqual(prod, domain.Product{}) {
		return fmt.Errorf("nenhum produto encontrado com o Id: %d", id)
	}

	delete(r.codeValueToID, prod.CodeValue)
	delete(r.products, id)

	return nil
}
