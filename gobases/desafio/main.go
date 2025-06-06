package main

import (
	"fmt"
	"log"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	repo, err := tickets.NewRepository("tickets.csv")

	if err != nil {
		log.Fatalf("Não foi possivel iniciar o repositório de tickets: %v", err)
	}

	// Ejemplo 1
	country := "china"
	respCountries, err := repo.GetTotalTickets(country)

	if err != nil {
		fmt.Printf("Erro: %+v", err)
	} else {
		fmt.Printf("Total de voos para %s: %d", country, respCountries)
	}

	// Ejemplo 2
	periodTime := "manhã"
	respMornings, err := repo.GetMornings(periodTime)

	if err != nil {
		fmt.Printf("Erro: %+v", err)
	} else {
		fmt.Printf("Total de voos pela %s: %d", periodTime, respMornings)
	}

	// Ejemplo 3
	hour := 2
	country = "china"

	avg, err := repo.AverageDestination(country, hour)

	if err != nil {
		fmt.Printf("Erro: %+v", err)
	} else {
		fmt.Printf("Total de voos para %s na hora %d: %f", country, hour, avg)
	}

}
