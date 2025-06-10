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
	Patch(id int, body map[string]interface{}) (*domain.Product, error)
	Delete(id int) error
}

type service struct {
	repo *Repository
}

func NewService(repo *Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll() ([]domain.Product, error) {
	products := s.repo.GetAll()
	return products, nil
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
