package ticket

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/RaphaelBorba/challenge_web/internal/domain"
)

type Repository interface {
	GetAll() (map[int]domain.Ticket, error)
	GetById(id int) (*domain.Ticket, error)
}

type repository struct {
	tickets map[int]domain.Ticket
}

func NewRepository(path string) (*repository, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rdr := csv.NewReader(f)

	ticketsMap := make(map[int]domain.Ticket)
	for {
		rec, err := rdr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		id, err := strconv.Atoi(rec[0])
		if err != nil {
			return nil, err
		}
		price, err := strconv.ParseFloat(rec[5], 64)
		if err != nil {
			return nil, err
		}

		ticketsMap[id] = domain.Ticket{
			Id:      id,
			Name:    rec[1],
			Email:   rec[2],
			Country: rec[3],
			Hour:    rec[4],
			Price:   price,
		}
	}

	return &repository{tickets: ticketsMap}, nil
}

func (r *repository) GetAll() (map[int]domain.Ticket, error) {
	return r.tickets, nil
}

func (r *repository) GetById(id int) (*domain.Ticket, error) {
	for _, t := range r.tickets {
		if t.Id == id {
			return &t, nil
		}
	}
	return nil, nil
}
