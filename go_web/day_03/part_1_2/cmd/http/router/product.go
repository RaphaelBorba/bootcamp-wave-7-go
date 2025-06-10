package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raphaelBorba/api-chi/put-patch-delete/cmd/http/handlers"
	"github.com/raphaelBorba/api-chi/put-patch-delete/internal/product"
)

func buildProductsRoutes() http.Handler {

	r := chi.NewRouter()
	repo, _ := product.NewRepository("../../products.json")

	service := product.NewService(repo)
	handler := handlers.NewProductHandler(service)

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		middleware.Timeout(5*time.Second),
		middleware.Heartbeat("/ping"),
	)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", handler.GetAll())
		r.Get("/{id}", handler.GetById())
		r.Post("/", handler.Create())
		r.Put("/{id}", handler.Update())
		r.Patch("/{id}", handler.Patch())
		r.Delete("/{id}", handler.Delete())
	})

	return r

}
