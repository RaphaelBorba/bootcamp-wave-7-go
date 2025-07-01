package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type router struct {
}

func (router *router) MapRoutes() http.Handler {

	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.Heartbeat("/health"),
	)

	r.Route("/api/v1", func(rp chi.Router) {
		rp.Route("/tickets", func(rp chi.Router) {
			rp.Mount("/", buildTicketsRoutes())
		})
	})

	return r
}

func NewRouter() *router {
	return &router{}
}
