package routes

import (
	"log"
	"net/http"

	"github.com/RaphaelBorba/challenge_web/cmd/handler"
	"github.com/RaphaelBorba/challenge_web/internal/ticket"
	"github.com/go-chi/chi/v5"
)

func buildTicketsRoutes() http.Handler {

	r := chi.NewRouter()

	repo, err := ticket.NewRepository("db/tickets.csv")
	if err != nil {
		log.Fatalf("failed to load tickets: %v", err)
	}

	svc := ticket.NewService(repo)
	h := handler.NewTicketHandler(svc)

	r.Get("/", h.GetAll())
	r.Get("/{ticketId}", h.GetByID())

	return r

}
