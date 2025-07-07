package main

import (
	"app/internal/application"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// env
	// ...
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on real env vars")
	}

	// app
	// - config
	app := application.NewApplicationDefault("", "./docs/db/json/products.json")
	// - tear down
	defer app.TearDown()
	// - set up
	if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	}
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
