package product

import (
	"fmt"

	"github.com/raphaelBorba/api-chi/internal/domain"
)

type service struct {
	repostory Repository
}

type Service interface {
	GetById(id int) (*domain.Product, error)
	Create(body domain.CreateProductRequest) (*domain.Product, error)
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetById(id int) (*domain.Product, error) {

	prod, err := s.repostory.GetById(id)

	if err != nil {
		return nil, fmt.Errorf("Produto não encontrado: %+v", err)
	}
	return prod, nil
}

func (s *service) Create(body domain.CreateProductRequest) (*domain.Product, error) {

	prod, err := s.repostory.Create(body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao criar o produto: %+v", err)
	}

	return prod, nil
}
