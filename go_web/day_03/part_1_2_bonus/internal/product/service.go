package product

import (
	"fmt"

	"github.com/raphaelBorba/api-chi/put-patch-delete/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetById(id int) (*domain.Product, error)
	Create(domain.CreateProductRequest) (*domain.Product, error)
	Update(id int, body domain.CreateProductRequest) (*domain.Product, error)
	Patch(id int, body map[string]any) (*domain.Product, error)
	Delete(id int) error
	GetConsumerPrice(productsids []int) (*domain.ConsumerPrice, error)
}

type service struct {
	repo *Repository
}

func NewService(repo *Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll() ([]domain.Product, error) {
	products := s.repo.GetAll()
	productsSlice := make([]domain.Product, 0, len(products))
	for _, v := range products {
		productsSlice = append(productsSlice, v)
	}
	return productsSlice, nil
}

func (s *service) GetById(id int) (*domain.Product, error) {
	prod, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("produto não encontrado: %w", err)
	}
	return prod, nil
}

func (s *service) Create(req domain.CreateProductRequest) (*domain.Product, error) {
	prod, err := s.repo.Create(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar produto: %w", err)
	}
	return prod, nil
}

func (s *service) Update(id int, req domain.CreateProductRequest) (*domain.Product, error) {
	prod, err := s.repo.Update(id, req)
	if err != nil {
		return nil, fmt.Errorf("erro ao atualizar produto: %w", err)
	}
	return prod, nil
}

func (s *service) Patch(id int, patchData map[string]any) (*domain.Product, error) {
	prod, err := s.repo.Patch(id, patchData)
	if err != nil {
		return nil, fmt.Errorf("erro ao aplicar patch no produto: %w", err)
	}
	return prod, nil
}

func (s *service) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("erro ao deletar produto: %w", err)
	}
	return nil
}

func (s *service) GetConsumerPrice(productsids []int) (*domain.ConsumerPrice, error) {

	prodQtd := make(map[int]int)
	products := s.repo.GetAll()
	selectedProds := []domain.Product{}
	totalPrice := 0.0

	if len(productsids) == 0 {
		for id := range products {
			selectedProds = append(selectedProds, products[id])
			totalPrice += products[id].Price
		}
	} else {
		for _, id := range productsids {
			prodQtd[id]++
		}
		for id, qtd := range prodQtd {
			if products[id].Quantity >= qtd {
				if products[id].IsPublished {
					for range qtd {
						selectedProds = append(selectedProds, products[id])
						totalPrice += products[id].Price
					}
				} else {
					return nil, fmt.Errorf("o produto id = %d, não está publicado", id)
				}
			} else {
				return nil, fmt.Errorf("o produto id = %d, tem uma quantidade inferior a solicitada", id)
			}
		}
	}

	selectedProdsQtd := len(selectedProds)
	switch {
	case selectedProdsQtd < 10:
		totalPrice *= 1.21
	case selectedProdsQtd < 20:
		totalPrice *= 1.17
	default:
		totalPrice *= 1.15
	}

	return &domain.ConsumerPrice{
		Products:   selectedProds,
		TotalPrice: totalPrice,
	}, nil
}
