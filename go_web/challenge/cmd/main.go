package main

import (
	"log"
	"net/http"

	"github.com/RaphaelBorba/challenge_web/cmd/routes"
)

func main() {

	r := routes.NewRouter()

	log.Println("👉 Listening on :8080")
	if err := http.ListenAndServe(":8080", r.MapRoutes()); err != nil {
		log.Fatal(err)
	}
}
