package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RaphaelBorba/challenge_web/internal/ticket"
)

type TicketHandler struct {
	svc ticket.Service
}

func NewTicketHandler(svc ticket.Service) *TicketHandler {
	return &TicketHandler{svc: svc}
}

func (h *TicketHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickets, err := h.svc.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tickets)
	}
}
