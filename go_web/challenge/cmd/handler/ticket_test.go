package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/RaphaelBorba/challenge_web/cmd/handler"
	"github.com/RaphaelBorba/challenge_web/internal/ticket"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func createTestCSV(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "tickets_*.csv")
	require.NoError(t, err)
	t.Cleanup(func() { os.Remove(f.Name()) })

	_, err = f.WriteString(content)
	require.NoError(t, err)

	err = f.Close()
	require.NoError(t, err)
	return f.Name()
}

func setupTestServer(h *handler.TicketHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/tickets/getAll", h.GetAll())
	r.Get("/tickets/getById/{ticketId}", h.GetByID())
	r.Get("/tickets/getByCountry/{dest}", h.GetByCountry())
	r.Get("/tickets/getAverage/{dest}", h.GetAverage())
	return r
}

func TestTicketHandler_GetAll(t *testing.T) {
	csvContent := "1,John Doe,johndoe@example.com,USA,10:00,150.50\n" +
		"2,Jane Smith,janesmith@example.com,Brazil,12:30,250.75\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)
	service := ticket.NewService(repo)
	h := handler.NewTicketHandler(service)
	r := setupTestServer(h)

	req := httptest.NewRequest(http.MethodGet, "/tickets/getAll", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	var body map[string]any
	err = json.Unmarshal(rr.Body.Bytes(), &body)
	require.NoError(t, err)

	require.Equal(t, "ok", body["message"])
	data, ok := body["data"].([]any)
	require.True(t, ok)
	require.Len(t, data, 2)
}

func TestTicketHandler_GetByID(t *testing.T) {
	csvContent := "10,Clark Kent,clark@dailyplanet.com,USA,09:00,300.0\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)
	service := ticket.NewService(repo)
	h := handler.NewTicketHandler(service)
	r := setupTestServer(h)

	t.Run("success: get by id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tickets/getById/10", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		var body map[string]any
		err = json.Unmarshal(rr.Body.Bytes(), &body)
		require.NoError(t, err)
		data, ok := body["data"].(map[string]any)
		require.True(t, ok)
		require.Equal(t, float64(10), data["id"])
		require.Equal(t, "Clark Kent", data["name"])
	})

	t.Run("error: invalid id format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tickets/getById/not-a-number", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("error: validation error for negative id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tickets/getById/-5", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("error: ticket not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tickets/getById/999", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestTicketHandler_GetByCountry(t *testing.T) {
	csvContent := "1,A,a@a.com,Brazil,10:00,100\n" +
		"2,B,b@b.com,USA,11:00,200\n" +
		"3,C,c@c.com,Brazil,12:00,150\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)
	service := ticket.NewService(repo)
	h := handler.NewTicketHandler(service)
	r := setupTestServer(h)

	t.Run("success: get by country", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tickets/getByCountry/Brazil", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		var body map[string]any
		err = json.Unmarshal(rr.Body.Bytes(), &body)
		require.NoError(t, err)
		require.Equal(t, float64(2), body["data"])
	})

	t.Run("error: country not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tickets/getByCountry/Argentina", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestTicketHandler_GetAverage(t *testing.T) {
	csvContent := "1,A,a@a.com,Chile,10:00,100\n" +
		"2,B,b@b.com,Chile,11:00,100\n" +
		"3,C,c@c.com,Mexico,12:00,100\n" +
		"4,D,d@d.com,USA,13:00,200\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)
	service := ticket.NewService(repo)
	h := handler.NewTicketHandler(service)
	r := setupTestServer(h)

	t.Run("success: get average", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tickets/getAverage/Chile", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		var body map[string]any
		err = json.Unmarshal(rr.Body.Bytes(), &body)
		require.NoError(t, err)
		require.Equal(t, float64(50), body["data"], "Expected average for Chile to be 50%")
	})

	t.Run("error: country not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/tickets/getAverage/Peru", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusNotFound, rr.Code)
	})
}
