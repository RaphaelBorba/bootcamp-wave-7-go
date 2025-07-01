package ticket

import (
	"github.com/RaphaelBorba/challenge_web/internal/domain"
	"github.com/RaphaelBorba/challenge_web/pkg/apperrors"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetById(id int) (*domain.Ticket, error)
	CountTicketsByDestiny(destiny string) (int, error)
	GetAverage(destiny string) (float64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]domain.Ticket, error) {
	return s.repo.GetAll()
}

func (s *service) GetById(id int) (*domain.Ticket, error) {
	if id <= 0 {
		return nil, apperrors.ErrValidation
	}
	t, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *service) CountTicketsByDestiny(destiny string) (int, error) {
	return s.repo.CountTicketsByDestiny(destiny)
}

func (s *service) GetAverage(destiny string) (float64, error) {
	return s.repo.GetAverage(destiny)
}
