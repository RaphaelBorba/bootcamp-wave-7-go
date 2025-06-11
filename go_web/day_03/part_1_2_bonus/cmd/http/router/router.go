package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raphaelBorba/api-chi/put-patch-delete/cmd/http/middlewares"
)

type router struct {
}

func (router *router) MapRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		middleware.Timeout(5*time.Second),
		middleware.Heartbeat("/ping"),

		middlewares.Auth,
	)

	r.Route("/api/v1", func(rp chi.Router) {
		rp.Mount("/", buildProductsRoutes())
	})

	return r
}

func NewRouter() *router {
	return &router{}
}
