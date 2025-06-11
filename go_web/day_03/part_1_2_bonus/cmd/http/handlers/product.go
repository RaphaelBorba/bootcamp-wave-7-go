package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/raphaelBorba/api-chi/put-patch-delete/internal/domain"
	"github.com/raphaelBorba/api-chi/put-patch-delete/internal/product"
)

type ProductHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *ProductHandler {
	return &ProductHandler{service}
}

func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prods, err := h.service.GetAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		response.JSON(w, http.StatusOK, prods)
	}
}

func (h *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido, deve ser um número"})
			return
		}

		prod, err := h.service.GetById(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{"error": "produto não encontrado"})
			return
		}

		response.JSON(w, http.StatusOK, prod)
	}
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req domain.CreateProductRequest
		if err := request.JSON(r, &req); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
			return
		}

		if strings.TrimSpace(req.Name) == "" {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "campo 'name' é obrigatório"})
			return
		}
		if req.Quantity < 0 {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "campo 'quantity' não pode ser negativo"})
			return
		}
		if strings.TrimSpace(req.CodeValue) == "" {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "campo 'code_value' é obrigatório"})
			return
		}
		if req.Price < 0 {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "campo 'price' não pode ser negativo"})
			return
		}

		const layout = "02/01/2006"
		parsed, err := time.Parse(layout, req.Expiration)
		if err != nil || parsed.Format(layout) != req.Expiration {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "campo 'expiration' inválido, use DD/MM/YYYY"})
			return
		}

		created, err := h.service.Create(req)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		response.JSON(w, http.StatusCreated, created)
	}
}

func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido, deve ser um número"})
			return
		}

		var req domain.CreateProductRequest
		if err := request.JSON(r, &req); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
			return
		}

		if strings.TrimSpace(req.Name) == "" || req.Quantity < 0 || strings.TrimSpace(req.CodeValue) == "" || req.Price < 0 {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "payload inválido"})
			return
		}

		const layout = "02/01/2006"
		if parsed, err := time.Parse(layout, req.Expiration); err != nil || parsed.Format(layout) != req.Expiration {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "campo 'expiration' inválido, use DD/MM/YYYY"})
			return
		}

		updated, err := h.service.Update(id, req)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		response.JSON(w, http.StatusOK, updated)
	}
}

func (h *ProductHandler) Patch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido, deve ser um número"})
			return
		}

		var patchData map[string]interface{}
		if err := request.JSON(r, &patchData); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
			return
		}

		patched, err := h.service.Patch(id, patchData)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		response.JSON(w, http.StatusOK, patched)
	}
}

func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": "ID inválido, deve ser um número"})
			return
		}

		if err := h.service.Delete(id); err != nil {
			response.JSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

func (h *ProductHandler) GetConsumerPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		raw := r.URL.Query().Get("list")

		var productsIds []int
		if raw != "" {
			if err := json.Unmarshal([]byte(raw), &productsIds); err != nil {
				response.JSON(w, http.StatusBadRequest, map[string]string{"error": "parâmetro 'list' inválido"})
				return
			}
		}

		result, err := h.service.GetConsumerPrice(productsIds)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]string{"error": fmt.Sprintf("%v", err)})
		}

		response.JSON(w, http.StatusOK, result)
	}
}
