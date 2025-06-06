package main

import (
	"fmt"
)

type Students struct {
	Name    string
	Surname string
	DNI     string
	Date    string
}

func (s Students) Detail() {
	fmt.Printf("Nome: %s\n\n", s.Name)
	fmt.Printf("Sobrenome: %s\n\n", s.Surname)
	fmt.Printf("ID: %s\n\n", s.DNI)
	fmt.Printf("Data: %s\n", s.Date)
}

func main() {
	aluno := Students{
		Name:    "Mariana",
		Surname: "Almeida",
		DNI:     "202500123",
		Date:    "2025-02-15",
	}

	aluno.Detail()
}
