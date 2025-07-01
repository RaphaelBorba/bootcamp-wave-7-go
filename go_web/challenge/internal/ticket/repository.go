package ticket

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/RaphaelBorba/challenge_web/internal/domain"
	"github.com/RaphaelBorba/challenge_web/pkg/apperrors"
)

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetById(id int) (*domain.Ticket, error)
}

type repository struct {
	tickets map[int]domain.Ticket
}

func NewRepository(path string) (*repository, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("repository: failed to open CSV file %q: %v", path, err)
		return nil, apperrors.ErrInternalError
	}
	defer f.Close()

	rdr := csv.NewReader(f)

	m := make(map[int]domain.Ticket)
	for {
		rec, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("repository: error reading CSV record: %v", err)
			return nil, apperrors.ErrInternalError
		}

		id, err := strconv.Atoi(rec[0])
		if err != nil {
			return nil, apperrors.ErrValidation
		}
		price, err := strconv.ParseFloat(rec[5], 64)
		if err != nil {
			return nil, apperrors.ErrValidation
		}

		m[id] = domain.Ticket{
			Id:      id,
			Name:    rec[1],
			Email:   rec[2],
			Country: rec[3],
			Hour:    rec[4],
			Price:   price,
		}
	}

	return &repository{tickets: m}, nil
}

func (r *repository) GetAll() ([]domain.Ticket, error) {
	out := make([]domain.Ticket, 0, len(r.tickets))
	for _, t := range r.tickets {
		out = append(out, t)
	}
	return out, nil
}

func (r *repository) GetById(id int) (*domain.Ticket, error) {
	t, ok := r.tickets[id]
	if !ok {
		return nil, apperrors.ErrNotFound
	}
	return &t, nil
}
