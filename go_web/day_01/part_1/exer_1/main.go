package main

import (
	"fmt"
	"net/http"
)

func pingPongHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "pong")
}

func main() {

	http.HandleFunc("/ping", pingPongHandler)

	fmt.Println("Server running in port: 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)

	}
}
