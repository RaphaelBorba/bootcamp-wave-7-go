package ticket_test

import (
	"os"
	"testing"

	"github.com/RaphaelBorba/challenge_web/internal/domain"
	"github.com/RaphaelBorba/challenge_web/internal/ticket"
	"github.com/RaphaelBorba/challenge_web/pkg/apperrors"
	"github.com/stretchr/testify/require"
)

func createTestCSV(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "tickets_*.csv")
	require.NoError(t, err)

	_, err = f.WriteString(content)
	require.NoError(t, err)

	err = f.Close()
	require.NoError(t, err)

	t.Cleanup(func() {
		os.Remove(f.Name())
	})

	return f.Name()
}

func TestNewRepository(t *testing.T) {
	t.Run("success: creates repository from valid CSV", func(t *testing.T) {
		csvContent := "1,John Doe,johndoe@example.com,Brazil,10:00,150.50\n" +
			"2,Jane Smith,janesmith@example.com,USA,12:30,250.75\n"
		path := createTestCSV(t, csvContent)

		repo, err := ticket.NewRepository(path)

		require.NoError(t, err)
		require.NotNil(t, repo)

		allTickets, err := repo.GetAll()
		require.NoError(t, err)
		require.Len(t, allTickets, 2)
	})

	t.Run("error: file does not exist", func(t *testing.T) {
		nonExistentPath := "non_existent_file.csv"

		repo, err := ticket.NewRepository(nonExistentPath)

		require.Error(t, err)
		require.Nil(t, repo)
		require.Equal(t, apperrors.ErrInternalError, err)
	})

	t.Run("error: invalid ID in CSV", func(t *testing.T) {
		csvContent := "not-an-id,John Doe,johndoe@example.com,Brazil,10:00,150.50\n"
		path := createTestCSV(t, csvContent)

		repo, err := ticket.NewRepository(path)

		require.Error(t, err)
		require.Nil(t, repo)
		require.Equal(t, apperrors.ErrValidation, err)
	})

	t.Run("error: invalid price in CSV", func(t *testing.T) {
		csvContent := "1,John Doe,johndoe@example.com,Brazil,10:00,not-a-price\n"
		path := createTestCSV(t, csvContent)

		repo, err := ticket.NewRepository(path)

		require.Error(t, err)
		require.Nil(t, repo)
		require.Equal(t, apperrors.ErrValidation, err)
	})
}

func TestGetAll(t *testing.T) {
	csvContent := "1,John Doe,johndoe@example.com,Brazil,10:00,150.50\n" +
		"2,Jane Smith,janesmith@example.com,USA,12:30,250.75\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)

	tickets, err := repo.GetAll()

	require.NoError(t, err)
	require.Len(t, tickets, 2)

	expected := []domain.Ticket{
		{Id: 1, Name: "John Doe", Email: "johndoe@example.com", Country: "Brazil", Hour: "10:00", Price: 150.50},
		{Id: 2, Name: "Jane Smith", Email: "janesmith@example.com", Country: "USA", Hour: "12:30", Price: 250.75},
	}
	require.ElementsMatch(t, expected, tickets)
}

func TestGetById(t *testing.T) {

	csvContent := "10,Clark Kent,clark@dailyplanet.com,USA,09:00,300.0\n" +
		"20,Bruce Wayne,bruce@wayne.com,USA,23:00,1000.0\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)

	t.Run("success: found", func(t *testing.T) {
		foundTicket, err := repo.GetById(10)

		require.NoError(t, err)
		require.NotNil(t, foundTicket)
		expected := &domain.Ticket{
			Id: 10, Name: "Clark Kent", Email: "clark@dailyplanet.com", Country: "USA", Hour: "09:00", Price: 300.0,
		}
		require.Equal(t, expected, foundTicket)
	})

	t.Run("error: not found", func(t *testing.T) {
		foundTicket, err := repo.GetById(99)

		require.Error(t, err)
		require.Nil(t, foundTicket)
		require.Equal(t, apperrors.ErrNotFound, err)
	})
}

func TestCountTicketsByDestiny(t *testing.T) {
	csvContent := "1,John Doe,johndoe@example.com,Brazil,10:00,150.50\n" +
		"2,Jane Smith,janesmith@example.com,USA,12:30,250.75\n" +
		"3,Pedro Álvares,pedro@brasil.com,Brazil,14:00,180.00\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)

	t.Run("success: counts tickets for a destination", func(t *testing.T) {
		count, err := repo.CountTicketsByDestiny("Brazil")

		require.NoError(t, err)
		require.Equal(t, 2, count)
	})

	t.Run("success: counts tickets case-insensitively", func(t *testing.T) {
		count, err := repo.CountTicketsByDestiny("brazil")

		require.NoError(t, err)
		require.Equal(t, 2, count)
	})

	t.Run("error: destination not found", func(t *testing.T) {
		count, err := repo.CountTicketsByDestiny("Argentina")

		require.Error(t, err)
		require.Equal(t, 0, count)
		require.Equal(t, apperrors.ErrNotFound, err)
	})
}

func TestGetAverage(t *testing.T) {
	csvContent := "1,A,a@a.com,Chile,10:00,100\n" +
		"2,B,b@b.com,Chile,11:00,100\n" +
		"3,C,c@c.com,Chile,12:00,100\n" +
		"4,D,d@d.com,Mexico,13:00,200\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)

	t.Run("success: calculates average for destination", func(t *testing.T) {
		avg, err := repo.GetAverage("Chile")

		require.NoError(t, err)
		require.Equal(t, 75.0, avg, "Expected average for Chile to be 75%")
	})

	t.Run("success: calculates average case-insensitively", func(t *testing.T) {
		avg, err := repo.GetAverage("mexico")

		require.NoError(t, err)
		require.Equal(t, 25.0, avg, "Expected average for Mexico to be 25%")
	})

	t.Run("error: destination not found", func(t *testing.T) {
		avg, err := repo.GetAverage("Peru")

		require.Error(t, err)
		require.Equal(t, 0.0, avg)
		require.Equal(t, apperrors.ErrNotFound, err)
	})
}
