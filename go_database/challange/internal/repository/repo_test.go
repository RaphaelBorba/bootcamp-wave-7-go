package repository

import (
	"database/sql"
	"testing"
	"time"

	"app/internal"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

const dsn = "root:rootpass@tcp(localhost:3306)/fantasy_products_test?parseTime=true"
const txdbDriverName = "txdb-fantasy-products"

func init() {
	txdb.Register(txdbDriverName, "mysql", dsn)
}

func seedData(t *testing.T, db *sql.DB) {
	mustExec := func(query string, args ...interface{}) {
		_, err := db.Exec(query, args...)
		if err != nil {
			t.Fatalf("Failed to execute query: %s. Error: %v", query, err)
		}
	}

	mustExec("SET FOREIGN_KEY_CHECKS = 0;")
	mustExec("TRUNCATE TABLE sales;")
	mustExec("TRUNCATE TABLE invoices;")
	mustExec("TRUNCATE TABLE products;")
	mustExec("TRUNCATE TABLE customers;")
	mustExec("SET FOREIGN_KEY_CHECKS = 1;")

	mustExec("INSERT INTO customers (id, first_name, last_name, `condition`) VALUES (1, 'John', 'Doe', 1), (2, 'Jane', 'Smith', 0), (3, 'Peter', 'Jones', 1), (4, 'Lisa', 'Ray', 0);")

	mustExec("INSERT INTO products (id, description, price) VALUES (1, 'Laptop', 1200.50), (2, 'Mouse', 25.00), (3, 'Keyboard', 75.75), (4, 'Monitor', 300.00), (5, 'Webcam', 99.99);")

	now := time.Now().Format("2006-01-02 15:04:05")
	mustExec("INSERT INTO invoices (id, `datetime`, customer_id, total) VALUES (1, ?, 1, 1225.50), (2, ?, 2, 375.75), (3, ?, 3, 125.74), (4, ?, 1, 300.00);", now, now, now, now)

	mustExec("INSERT INTO sales (id, quantity, invoice_id, product_id) VALUES (1, 1, 1, 1), (2, 1, 1, 2), (3, 5, 2, 3), (4, 1, 3, 2), (5, 1, 3, 5), (6, 1, 4, 4);")
}

func TestProductsMySQL_GetTopSellingProducts(t *testing.T) {
	db, err := sql.Open(txdbDriverName, t.Name())
	assert.NoError(t, err)
	defer db.Close()

	seedData(t, db)

	repo := NewProductsMySQL(db)

	t.Run("should return top 5 selling products", func(t *testing.T) {
		products, err := repo.GetTopSellingProducts()

		assert.NoError(t, err)
		assert.NotNil(t, products)
		assert.Len(t, products, 5)

		assert.Equal(t, "Keyboard", products[0].Description)
		assert.Equal(t, float64(5), products[0].TotalQuantity)

		assert.Equal(t, "Mouse", products[1].Description)
		assert.Equal(t, float64(2), products[1].TotalQuantity)

		remainingProducts := products[2:]
		expectedRemaining := map[string]bool{
			"Laptop":  false,
			"Monitor": false,
			"Webcam":  false,
		}

		for _, p := range remainingProducts {
			if _, ok := expectedRemaining[p.Description]; ok {
				expectedRemaining[p.Description] = true
				assert.Equal(t, float64(1), p.TotalQuantity, "product %s should have quantity 1", p.Description)
			}
		}

		for desc, found := range expectedRemaining {
			assert.True(t, found, "Expected to find product '%s' in the results", desc)
		}
	})

	t.Run("should return empty slice when no sales exist", func(t *testing.T) {
		_, err := db.Exec("TRUNCATE TABLE sales;")
		assert.NoError(t, err)

		products, _ := repo.GetTopSellingProducts()

		assert.NotNil(t, products)
		assert.Len(t, products, 0)
	})
}

func TestCustomersMySQL_GetTotalSpentByCondition(t *testing.T) {
	db, err := sql.Open(txdbDriverName, t.Name())
	assert.NoError(t, err)
	defer db.Close()

	seedData(t, db)
	repo := NewCustomersMySQL(db)

	t.Run("should return total spent grouped by condition", func(t *testing.T) {
		totals, err := repo.GetTotalSpentByCondition()

		assert.NoError(t, err)
		assert.NotNil(t, totals)
		assert.Len(t, totals, 2)

		expected := []internal.ConditionTotal{
			{Condition: 0, TotalSpent: 378.75},
			{Condition: 1, TotalSpent: 1650.49},
		}

		assert.Equal(t, expected, totals)
	})

	t.Run("should return empty slice when no sales exist", func(t *testing.T) {
		_, err := db.Exec("TRUNCATE TABLE sales;")
		assert.NoError(t, err)

		totals, err := repo.GetTotalSpentByCondition()

		assert.NoError(t, err)
		assert.NotNil(t, totals)
		assert.Len(t, totals, 0)
	})
}

func TestCustomersMySQL_GetTopSpenders(t *testing.T) {
	db, err := sql.Open(txdbDriverName, t.Name())
	assert.NoError(t, err)
	defer db.Close()

	seedData(t, db)
	repo := NewCustomersMySQL(db)

	t.Run("should return top 5 spenders", func(t *testing.T) {
		spenders, err := repo.GetTopSpenders()

		assert.NoError(t, err)
		assert.NotNil(t, spenders)
		assert.Len(t, spenders, 3)

		expected := []internal.CustomerSpender{
			{FirstName: "John", LastName: "Doe", TotalSpent: 1525.50},
			{FirstName: "Jane", LastName: "Smith", TotalSpent: 378.75},
			{FirstName: "Peter", LastName: "Jones", TotalSpent: 124.99},
		}

		assert.Equal(t, expected, spenders)
	})

	t.Run("should return empty slice when no sales exist", func(t *testing.T) {
		_, err := db.Exec("TRUNCATE TABLE sales;")
		assert.NoError(t, err)

		spenders, err := repo.GetTopSpenders()

		assert.NoError(t, err)
		assert.NotNil(t, spenders)
		assert.Len(t, spenders, 0)
	})
}
