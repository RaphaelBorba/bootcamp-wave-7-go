package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphaelBorba/api-chi/put-patch-delete/cmd/http/handlers"
	"github.com/raphaelBorba/api-chi/put-patch-delete/internal/product"
)

func buildProductsRoutes() http.Handler {

	r := chi.NewRouter()
	repo, _ := product.NewRepository("../../products.json")

	service := product.NewService(repo)
	handler := handlers.NewProductHandler(service)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", handler.GetAll())
		r.Get("/{id}", handler.GetById())
		r.Post("/", handler.Create())
		r.Put("/{id}", handler.Update())
		r.Patch("/{id}", handler.Patch())
		r.Delete("/{id}", handler.Delete())
		r.Get("/consumer_price", handler.GetConsumerPrice())
	})

	return r

}
