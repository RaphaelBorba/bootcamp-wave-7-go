package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/RaphaelBorba/challenge_web/internal/ticket"
	"github.com/RaphaelBorba/challenge_web/pkg/apperrors"
	"github.com/RaphaelBorba/challenge_web/pkg/web/response"
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
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{"data": tickets, "message": "ok"})
	}
}

func (h *TicketHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "ticketId")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, apperrors.ErrValidation.Error())
			return
		}

		tkt, err := h.svc.GetById(id)
		if err != nil {
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
				response.Error(w, http.StatusNotFound, err.Error())
			case errors.Is(err, apperrors.ErrValidation):
				response.Error(w, http.StatusBadRequest, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, apperrors.ErrInternalError.Error())
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{"data": tkt, "message": "ok"})
	}
}
