package main

import (
	"fmt"
	"net/http"

	"github.com/raphaelBorba/api-chi/put-patch-delete/cmd/http/router"
)

func main() {

	r := router.NewRouter()

	addr := fmt.Sprintf("%s:%s", "localhost", "8080")
	fmt.Println("Starting server on", addr)

	if err := http.ListenAndServe(addr, r.MapRoutes()); err != nil {
		panic(err)
	}
}
