package tickets

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func writeTempCSV(t *testing.T, dir, content string) string {
	t.Helper()
	path := filepath.Join(dir, "test_tickets.csv")
	err := os.WriteFile(path, []byte(content), 0644)
	require.NoError(t, err)
	return path
}

func TestGetTotalTickets(t *testing.T) {
	tmpDir := t.TempDir()

	csvContent := `1,Alice,alice@example.com,Brazil,08:00,100.0
2,Bob,bob@example.com,USA,14:00,200.0
3,Carol,carol@example.com,Brazil,20:00,150.0
4,David,david@example.com,Brazil,05:30,120.0
`
	csvPath := writeTempCSV(t, tmpDir, csvContent)

	repo, err := NewRepository(csvPath)
	require.NoError(t, err)

	t.Run("destino válido presente", func(t *testing.T) {
		count, err := repo.GetTotalTickets("Brazil")
		require.NoError(t, err)
		require.Equal(t, 3, count, "deve contar 3 tickets para Brazil (case-insensitive)")
	})

	t.Run("destino válido ausente", func(t *testing.T) {
		count, err := repo.GetTotalTickets("Argentina")
		require.NoError(t, err)
		require.Equal(t, 0, count, "deve retornar 0 para país não presente no CSV")
	})

	t.Run("destino case‐insensitive", func(t *testing.T) {
		countLower, errLower := repo.GetTotalTickets("brazil")
		require.NoError(t, errLower)
		countMixed, errMixed := repo.GetTotalTickets("BrAzIl")
		require.NoError(t, errMixed)
		require.Equal(t, 3, countLower)
		require.Equal(t, 3, countMixed)
	})

	t.Run("destino vazio retorna erro", func(t *testing.T) {
		_, err := repo.GetTotalTickets("")
		require.Error(t, err)
	})
}

func TestGetMornings(t *testing.T) {
	tmpDir := t.TempDir()

	csvContent := `1,Alice,alice@example.com,Brazil,08:00,100.0
2,Bob,bob@example.com,USA,14:00,200.0
3,Carol,carol@example.com,Brazil,20:00,150.0
4,David,david@example.com,Brazil,05:30,120.0
`
	csvPath := writeTempCSV(t, tmpDir, csvContent)

	repo, err := NewRepository(csvPath)
	require.NoError(t, err)

	t.Run("período início da manhã (00:00–06:00)", func(t *testing.T) {
		count, err := repo.GetMornings("início da manhã")
		require.NoError(t, err)
		require.Equal(t, 1, count)
	})

	t.Run("período manhã (07:00–12:00)", func(t *testing.T) {
		count, err := repo.GetMornings("manhã")
		require.NoError(t, err)
		require.Equal(t, 1, count)
	})

	t.Run("período tarde (13:00–19:00)", func(t *testing.T) {
		count, err := repo.GetMornings("tarde")
		require.NoError(t, err)
		require.Equal(t, 1, count)
	})

	t.Run("período noite (20:00–23:00)", func(t *testing.T) {
		count, err := repo.GetMornings("noite")
		require.NoError(t, err)
		require.Equal(t, 1, count)
	})

	t.Run("período inválido retorna erro", func(t *testing.T) {
		_, err := repo.GetMornings("madrugada")
		require.Error(t, err)
	})

	t.Run("período vazio retorna erro", func(t *testing.T) {
		_, err := repo.GetMornings("")
		require.Error(t, err)
	})
}

func TestAverageDestination(t *testing.T) {
	tmpDir := t.TempDir()

	csvContent := `1,Alice,alice@example.com,Brazil,08:00,100.0
2,Bob,bob@example.com,USA,14:00,200.0
3,Carol,carol@example.com,Brazil,20:00,150.0
4,David,david@example.com,Brazil,08:00,120.0
`
	csvPath := writeTempCSV(t, tmpDir, csvContent)

	repo, err := NewRepository(csvPath)
	require.NoError(t, err)

	t.Run("percentual para país e horário válidos", func(t *testing.T) {
		percentage, err := repo.AverageDestination("Brazil", 8)
		require.NoError(t, err)
		require.InDelta(t, 50.0, percentage, 0.0001)
	})

	t.Run("percentual case‐insensitive", func(t *testing.T) {
		p1, err1 := repo.AverageDestination("brazil", 8)
		require.NoError(t, err1)
		p2, err2 := repo.AverageDestination("BrAzIl", 8)
		require.NoError(t, err2)
		require.InDelta(t, p1, p2, 0.0001)
	})

	t.Run("sem correspondências resulta em 0%", func(t *testing.T) {
		percentage, err := repo.AverageDestination("Argentina", 10)
		require.NoError(t, err)
		require.InDelta(t, 0.0, percentage, 0.0001)
	})

	t.Run("hora inválida retorna erro", func(t *testing.T) {
		_, err := repo.AverageDestination("Brazil", 24)
		require.Error(t, err)
	})

	t.Run("destino vazio retorna erro", func(t *testing.T) {
		_, err := repo.AverageDestination("", 8)
		require.Error(t, err)
	})

	t.Run("nenhum ticket no repositório retorna erro", func(t *testing.T) {
		emptyCSV := ``
		emptyPath := writeTempCSV(t, tmpDir, emptyCSV)
		emptyRepo, err := NewRepository(emptyPath)
		require.NoError(t, err)
		_, err = emptyRepo.AverageDestination("Brazil", 8)
		require.Error(t, err)
	})
}

func TestNewRepositoryErrors(t *testing.T) {
	tmpDir := t.TempDir()

	t.Run("arquivo inexistente retorna erro", func(t *testing.T) {
		_, err := NewRepository(filepath.Join(tmpDir, "arquivo_inexistente.csv"))
		require.Error(t, err)
	})

	t.Run("comprimento de registro inválido retorna erro", func(t *testing.T) {
		badCSV := `1,Alice,alice@example.com,Brazil,08:00`
		path := writeTempCSV(t, tmpDir, badCSV)
		_, err := NewRepository(path)
		require.Error(t, err)
	})

	t.Run("ID inválido retorna erro", func(t *testing.T) {
		badCSV := `not_an_int,Alice,alice@example.com,Brazil,08:00,100.0`
		path := writeTempCSV(t, tmpDir, badCSV)
		_, err := NewRepository(path)
		require.Error(t, err)
	})

	t.Run("formato de FlightHour inválido retorna erro", func(t *testing.T) {
		badCSV := `1,Alice,alice@example.com,Brazil,8am,100.0`
		path := writeTempCSV(t, tmpDir, badCSV)
		_, err := NewRepository(path)
		require.Error(t, err)
	})

	t.Run("preço inválido retorna erro", func(t *testing.T) {
		badCSV := `1,Alice,alice@example.com,Brazil,08:00,not_a_price`
		path := writeTempCSV(t, tmpDir, badCSV)
		_, err := NewRepository(path)
		require.Error(t, err)
	})
}
