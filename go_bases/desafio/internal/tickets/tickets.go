package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ticket struct {
	ID             int
	Name           string
	Email          string
	DestinyCountry string
	FlightHour     time.Time
	Price          float64
}

type Repository struct {
	tickets []Ticket
}

func NewRepository(csvPath string) (*Repository, error) {
	f, err := os.Open(csvPath)
	if err != nil {
		return nil, fmt.Errorf("could not open CSV file %q: %w", csvPath, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	var all []Ticket
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV: %w", err)
		}
		if len(record) < 6 {
			return nil, fmt.Errorf("unexpected record length: got %d fields, want ≥6", len(record))
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("invalid ID %q: %w", record[0], err)
		}

		name := record[1]
		email := record[2]
		country := strings.ToLower(record[3])

		flightHour, err := time.Parse("15:06", record[4])
		if err != nil {
			return nil, fmt.Errorf("invalid FlightHour %q: %w", record[4], err)
		}

		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid Price %q: %w", record[5], err)
		}

		all = append(all, Ticket{
			ID:             id,
			Name:           name,
			Email:          email,
			DestinyCountry: country,
			FlightHour:     flightHour,
			Price:          price,
		})
	}

	return &Repository{tickets: all}, nil
}

// ejemplo 1
func (r *Repository) GetTotalTickets(destination string) (int, error) {

	if len(destination) == 0 {
		return 0, errors.New("String is empty")
	}

	total := 0
	for _, tck := range r.tickets {
		if strings.ToLower(tck.DestinyCountry) == strings.ToLower(destination) {
			total++
		}
	}
	return total, nil
}

// ejemplo 2
func (r *Repository) GetMornings(periodTime string) (int, error) {

	if len(periodTime) == 0 {
		return 0, errors.New("String is empty")
	}

	var startHour, endHour int
	switch periodTime {
	case "início da manhã":
		startHour, endHour = 0, 6
	case "manhã":
		startHour, endHour = 7, 12
	case "tarde":
		startHour, endHour = 13, 19
	case "noite":
		startHour, endHour = 20, 23
	default:
		return 0, fmt.Errorf("período inválido: %q", periodTime)
	}

	total := 0
	for _, tck := range r.tickets {
		h := tck.FlightHour.Hour()
		if h >= startHour && h <= endHour {
			total++
		}
	}
	return total, nil
}

// ejemplo 3
func (r *Repository) AverageDestination(destination string, hour int) (float64, error) {
	if len(destination) == 0 {
		return 0, fmt.Errorf("String do país está vazia")
	}
	if hour < 0 || hour > 23 {
		return 0, fmt.Errorf("O horário precisa estar entre 0 e 23 (inclusive)")
	}
	totalTickets := len(r.tickets)
	if totalTickets == 0 {
		return 0, fmt.Errorf("nenhum ticket disponível")
	}

	destLower := strings.ToLower(destination)

	matchCount := 0
	for _, tck := range r.tickets {
		if strings.ToLower(tck.DestinyCountry) == destLower && tck.FlightHour.Hour() == hour {
			matchCount++
		}
	}

	percentage := (float64(matchCount) * 100.0) / float64(totalTickets)
	return percentage, nil
}
