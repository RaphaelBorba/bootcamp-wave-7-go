package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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

	r.Post("/products", PostCreateProductHandler(repo))

	r.Get("/products/{id}", GetProductByIDHandler(repo))

	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("ListenAndServe: %v", err)
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

func PostCreateProductHandler(repo *products.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req products.CreateProductRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(req.Name) == "" {
			http.Error(w, `{"error":"campo 'name' é obrigatório"}`, http.StatusBadRequest)
			return
		}
		if req.Quantity < 0 {
			http.Error(w, `{"error":"campo 'quantity' não pode ser negativo"}`, http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(req.CodeValue) == "" {
			http.Error(w, `{"error":"campo 'code_value' é obrigatório"}`, http.StatusBadRequest)
			return
		}
		if req.Price < 0 {
			http.Error(w, `{"error":"campo 'price' não pode ser negativo"}`, http.StatusBadRequest)
			return
		}
		const layout = "02/01/2006"
		parsed, err := time.Parse(layout, req.Expiration)
		if err != nil || parsed.Format(layout) != req.Expiration {
			http.Error(w,
				`{"error":"campo 'expiration' inválido, use DD/MM/YYYY"}`,
				http.StatusBadRequest,
			)
			return
		}

		created, err := repo.CreateProduct(req)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"%v"}`, err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	}
}
