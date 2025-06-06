package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raphaelBorba/api-chi/products"
)

func main() {

	repo, err := products.NewRepository("products.json")

	if err != nil {
		fmt.Println("Erro", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", GetPingPong())

	r.Get("/products", GetProductsHandler(repo))

	r.Get("/products/{id}", GetProductByIDHandler(repo))

	r.Get("/products/search", GetProductSearchHandler(repo))

	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("ListenAndServe: %v", err)
	}
}

func GetPingPong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content_Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "pong"})
	}
}

func GetProductsHandler(repo *products.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		all := repo.GetProducts()
		json.NewEncoder(w).Encode(all)
	}
}

func GetProductByIDHandler(repo *products.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w,
				`{"error":"ID inválido, deve ser um número"}`,
				http.StatusBadRequest,
			)
			return
		}

		prod, err := repo.GetProductById(id)
		if err != nil {
			http.Error(w,
				`{"error":"produto não encontrado"}`,
				http.StatusNotFound,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(prod)
	}
}

func GetProductSearchHandler(repo *products.Repository) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		priceParam := r.URL.Query().Get("priceGt")
		if priceParam == "" {
			http.Error(w,
				`{"error":"query param 'price' é obrigatório"}`,
				http.StatusBadRequest,
			)
			return
		}

		price, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			http.Error(w,
				`{"error":"price inválido, use um número (ex: 12.23)"}`,
				http.StatusBadRequest,
			)
			return
		}

		var result []products.Product
		for _, p := range repo.GetProducts() {
			if p.Price > price {
				result = append(result, p)
			}
		}

		if len(result) == 0 {
			msg := fmt.Sprintf(`{"error":"Nenhum produto encontrado com preço acima de %.2f"}`, price)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(msg))
			return
		}

		json.NewEncoder(w).Encode(result)

	}
}
