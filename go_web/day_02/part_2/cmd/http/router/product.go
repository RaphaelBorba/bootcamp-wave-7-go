package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphaelBorba/api-chi/cmd/http/handlers"
	"github.com/raphaelBorba/api-chi/internal/product"
)

func buildProductsRoutes() http.Handler {

	r := chi.NewRouter()
	repository := product.NewRepository("../../products.json")
	service := product.NewService(*repository)
	handler := handlers.NewProductHandler(service)

	r.Get("/{id}", handler.FindOneById())
	r.Post("/", handler.Create())

	return r

}
