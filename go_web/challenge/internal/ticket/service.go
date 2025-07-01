package ticket

import "github.com/RaphaelBorba/challenge_web/internal/domain"

type Service interface {
	GetAll() (map[int]domain.Ticket, error)
	GetById(id int) (*domain.Ticket, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{repo: r}
}

func (s *service) GetAll() (map[int]domain.Ticket, error) {
	return s.repo.GetAll()
}

func (s *service) GetById(id int) (*domain.Ticket, error) {
	return s.repo.GetById(id)
}
