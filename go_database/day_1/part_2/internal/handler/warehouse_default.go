package handler

import (
	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// NewProductsDefault returns a new instance of ProductsDefault
func NewWarehouseDefault(rp_w internal.RepositoryWarehouse) *WarehouseDefault {
	return &WarehouseDefault{
		rp_w: rp_w,
	}
}

// ProductsDefault is a struct that represents the default product handler
type WarehouseDefault struct {
	// rp_p is the product repository
	rp_w internal.RepositoryWarehouse
}

// ProductJSON is a struct that represents a product in JSON
type WarehouseJSON struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Adress    string `json:"adress"`
	Telephone string `json:"telephone"`
	Capacity  int    `json:"capacity"`
}

// GetOne returns a product by id
func (h *WarehouseDefault) GetOneWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		ware, err := h.rp_w.GetOneWarehouse(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Error(w, http.StatusNotFound, "warehouse not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize
		data := WarehouseJSON{
			ID:        ware.ID,
			Name:      ware.Name,
			Adress:    ware.Adress,
			Telephone: ware.Telephone,
			Capacity:  ware.Capacity,
		}
		response.JSON(w, http.StatusOK, map[string]any{"message": "product found", "data": data})
	}
}

// RequestBodyProductCreate is a struct that represents the request body of a product to create
type RequestBodyWarehouseJSONCreate struct {
	Name      string `json:"name"`
	Adress    string `json:"adress"`
	Telephone string `json:"telephone"`
	Capacity  int    `json:"capacity"`
}

// Create creates a product
func (h *WarehouseDefault) CreateWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var body RequestBodyWarehouseJSONCreate
		if err := request.JSON(r, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request body")
			return
		}

		// process
		ware := internal.Warehouse{
			Name:      body.Name,
			Adress:    body.Adress,
			Telephone: body.Telephone,
			Capacity:  body.Capacity,
		}
		if err := h.rp_w.StoreWarehouse(&ware); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotUnique):
				response.Error(w, http.StatusConflict, "product not unique")
			case errors.Is(err, internal.ErrProductRelation):
				response.Error(w, http.StatusConflict, "product relation error")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize
		data := WarehouseJSON{
			ID:        ware.ID,
			Name:      ware.Name,
			Adress:    ware.Adress,
			Telephone: ware.Telephone,
			Capacity:  ware.Capacity,
		}
		response.JSON(w, http.StatusCreated, map[string]any{"message": "product created", "data": data})
	}
}

// RequestBodyProductUpdate is a struct that represents the request body of a product to update
type RequestBodyWarehouseUpdate struct {
	Name      string `json:"name"`
	Adress    string `json:"adress"`
	Telephone string `json:"telephone"`
	Capacity  int    `json:"capacity"`
}

// Update updates a product
func (h *WarehouseDefault) UpdateWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - get product
		ware, err := h.rp_w.GetOneWarehouse(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Error(w, http.StatusNotFound, "product not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		// - patch product
		body := RequestBodyWarehouseUpdate{
			Name:      ware.Name,
			Adress:    ware.Adress,
			Telephone: ware.Telephone,
			Capacity:  ware.Capacity,
		}
		if err := request.JSON(r, &body); err != nil {
			response.Error(w, http.StatusBadRequest, "invalid request body")
			return
		}

		ware.Name = body.Name
		ware.Adress = body.Adress
		ware.Telephone = body.Telephone
		ware.Capacity = body.Capacity

		if err := h.rp_w.UpdateWarehouse(&ware); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotUnique):
				response.Error(w, http.StatusConflict, "warehouse not unique")
			case errors.Is(err, internal.ErrProductRelation):
				response.Error(w, http.StatusConflict, "warehouse relation error")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize
		data := WarehouseJSON{
			ID:        ware.ID,
			Name:      ware.Name,
			Adress:    ware.Adress,
			Telephone: ware.Telephone,
			Capacity:  ware.Capacity,
		}
		response.JSON(w, http.StatusOK, map[string]any{"message": "product updated", "data": data})
	}
}

// Delete deletes a product by id
func (h *WarehouseDefault) DeleteWarehouse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		if err := h.rp_w.DeleteWarehouse(id); err != nil {
			response.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{"message": "product deleted", "data": id})
	}
}
