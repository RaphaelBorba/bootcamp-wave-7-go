package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/raphaelBorba/api-chi/internal/domain"
	"github.com/raphaelBorba/api-chi/internal/product"
)

type ProductHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *ProductHandler {
	return &ProductHandler{
		service,
	}
}

func (h *ProductHandler) FindOneById() http.HandlerFunc {
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

		prod, err := h.service.FindOneById(id)

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

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req domain.CreateProductRequest

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

		created, err := h.service.Create(req)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"%v"}`, err), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	}
}
