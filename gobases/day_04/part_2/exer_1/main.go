package main

import (
	"errors"
	"fmt"
	"os"
)

type Client struct {
	Name    string
	ID      int
	Phone   string
	Address string
}

var clients = []Client{
	{Name: "Alice", ID: 1, Phone: "123456789", Address: "123 Main St"},
	{Name: "Bob", ID: 2, Phone: "987654321", Address: "456 Oak Ave"},
}

func checkExists(c Client) {
	for _, existing := range clients {
		if existing.ID == c.ID {
			panic("Error: client already exists")
		}
	}
}

func validateClient(c Client) error {
	if c.Name == "" {
		return fmt.Errorf("Error: Name is empty")
	}
	if c.ID == 0 {
		return fmt.Errorf("Error: ID is zero")
	}
	if c.Phone == "" {
		return fmt.Errorf("Error: Phone is empty")
	}
	if c.Address == "" {
		return fmt.Errorf("Error: Address is empty")
	}
	return nil
}

func appendToFile(c Client) error {
	f, err := os.OpenFile("customers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("Error: unable to open or create file")
	}
	defer f.Close()

	line := fmt.Sprintf("%s,%d,%s,%s\n", c.Name, c.ID, c.Phone, c.Address)
	if _, err := f.WriteString(line); err != nil {
		return errors.New("Error: unable to write to file")
	}
	return nil
}

func main() {
	defer fmt.Println("Several errors were detected at runtime")
	defer fmt.Println("End of execution")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	newClient := Client{
		Name:    "Charlie",
		ID:      2,
		Phone:   "555123456",
		Address: "789 Pine Rd",
	}

	checkExists(newClient)

	if err := validateClient(newClient); err != nil {
		fmt.Println(err)
		return
	}

	clients = append(clients, newClient)

	if err := appendToFile(newClient); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Client registered successfully")
}
