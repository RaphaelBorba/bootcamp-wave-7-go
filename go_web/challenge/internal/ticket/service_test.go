package ticket_test

import (
	"testing"

	"github.com/RaphaelBorba/challenge_web/internal/domain"
	"github.com/RaphaelBorba/challenge_web/internal/ticket"
	"github.com/RaphaelBorba/challenge_web/pkg/apperrors"
	"github.com/stretchr/testify/require"
)

// createTestCSV is a helper function to create a temporary CSV for testing.
// It returns the file path and a cleanup function to remove the file.

func TestServiceGetAll(t *testing.T) {
	t.Run("success: returns all tickets from repository", func(t *testing.T) {
		// Arrange
		csvContent := "1,John Doe,johndoe@example.com,USA,10:00,150.50\n" +
			"2,Jane Smith,janesmith@example.com,Brazil,12:30,250.75\n"
		path := createTestCSV(t, csvContent)
		repo, err := ticket.NewRepository(path)
		require.NoError(t, err)
		service := ticket.NewService(repo)

		expectedTickets := []domain.Ticket{
			{Id: 1, Name: "John Doe", Email: "johndoe@example.com", Country: "USA", Hour: "10:00", Price: 150.50},
			{Id: 2, Name: "Jane Smith", Email: "janesmith@example.com", Country: "Brazil", Hour: "12:30", Price: 250.75},
		}

		// Act
		result, err := service.GetAll()

		// Assert
		require.NoError(t, err)
		require.ElementsMatch(t, expectedTickets, result)
	})
}

func TestServiceGetById(t *testing.T) {
	// Arrange for repository-dependent tests
	csvContent := "5,Peter Jones,peter@example.com,Canada,15:00,450.0\n" +
		"10,Clark Kent,clark@dailyplanet.com,USA,09:00,300.0\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)
	service := ticket.NewService(repo)

	t.Run("success: returns ticket by id", func(t *testing.T) {
		// Arrange
		expectedTicket := &domain.Ticket{Id: 5, Name: "Peter Jones", Email: "peter@example.com", Country: "Canada", Hour: "15:00", Price: 450.0}

		// Act
		result, err := service.GetById(5)

		// Assert
		require.NoError(t, err)
		require.Equal(t, expectedTicket, result)
	})

	t.Run("error: invalid id (zero)", func(t *testing.T) {
		// Act
		result, err := service.GetById(0)

		// Assert
		require.Error(t, err)
		require.Nil(t, result)
		require.Equal(t, apperrors.ErrValidation, err)
	})

	t.Run("error: invalid id (negative)", func(t *testing.T) {
		// Act
		result, err := service.GetById(-10)

		// Assert
		require.Error(t, err)
		require.Nil(t, result)
		require.Equal(t, apperrors.ErrValidation, err)
	})

	t.Run("error: repository returns not found", func(t *testing.T) {
		// Act
		result, err := service.GetById(99)

		// Assert
		require.Error(t, err)
		require.Nil(t, result)
		require.Equal(t, apperrors.ErrNotFound, err)
	})
}

func TestServiceCountTicketsByDestiny(t *testing.T) {
	// Arrange
	csvContent := "1,A,a@a.com,Brazil,10:00,100\n" +
		"2,B,b@b.com,USA,11:00,200\n" +
		"3,C,c@c.com,Brazil,12:00,150\n" +
		"4,D,d@d.com,Brazil,13:00,120\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)
	service := ticket.NewService(repo)

	t.Run("success: returns count for a destination", func(t *testing.T) {
		// Act
		count, err := service.CountTicketsByDestiny("Brazil")

		// Assert
		require.NoError(t, err)
		require.Equal(t, 3, count)
	})

	t.Run("error: destination not found", func(t *testing.T) {
		// Act
		count, err := service.CountTicketsByDestiny("Argentina")

		// Assert
		require.Error(t, err)
		require.Equal(t, 0, count)
		require.Equal(t, apperrors.ErrNotFound, err)
	})
}

func TestServiceGetAverage(t *testing.T) {
	// Arrange
	csvContent := "1,A,a@a.com,Chile,10:00,100\n" + // Total 4 tickets
		"2,B,b@b.com,Chile,11:00,100\n" +
		"3,C,c@c.com,Chile,12:00,100\n" +
		"4,D,d@d.com,Mexico,13:00,200\n"
	path := createTestCSV(t, csvContent)
	repo, err := ticket.NewRepository(path)
	require.NoError(t, err)
	service := ticket.NewService(repo)

	t.Run("success: returns average for a destination", func(t *testing.T) {
		// Act: 3 tickets to Chile out of 4 total tickets -> (3 / 4) * 100 = 75
		avg, err := service.GetAverage("Chile")

		// Assert
		require.NoError(t, err)
		require.Equal(t, 75.0, avg)
	})

	t.Run("error: destination not found", func(t *testing.T) {
		// Act
		avg, err := service.GetAverage("Peru")

		// Assert
		require.Error(t, err)
		require.Equal(t, 0.0, avg)
		require.Equal(t, apperrors.ErrNotFound, err)
	})
}
