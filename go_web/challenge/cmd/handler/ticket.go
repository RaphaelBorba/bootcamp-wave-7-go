package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/RaphaelBorba/challenge_web/internal/ticket"
	"github.com/RaphaelBorba/challenge_web/pkg/apperrors"
	"github.com/go-chi/chi/v5"
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
			http.Error(w, apperrors.ErrInternalError.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tickets)
	}
}

func (h *TicketHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "ticketId")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, apperrors.ErrValidation.Error(), http.StatusBadRequest)
			return
		}

		t, err := h.svc.GetById(id)
		if err != nil {
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
				http.Error(w, err.Error(), http.StatusNotFound)
			case errors.Is(err, apperrors.ErrValidation):
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, apperrors.ErrInternalError.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(t)
	}
}
