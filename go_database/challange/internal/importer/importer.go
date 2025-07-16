package importer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"app/internal"
	"app/internal/repository"
	"app/internal/service"

	_ "github.com/go-sql-driver/mysql"
)

// Run reads JSON files from jsonDir, opens a DB connection, and imports their data into the database
func Run(jsonDir string) error {
	// Read DSN from environment

	dsn := "root:rootpass@tcp(localhost:3306)/fantasy_products?parseTime=true"

	// Open the database connection once
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer db.Close()

	// Initialize repositories
	rpCustomer := repository.NewCustomersMySQL(db)
	rpProduct := repository.NewProductsMySQL(db)
	rpInvoice := repository.NewInvoicesMySQL(db)
	rpSale := repository.NewSalesMySQL(db)

	// Initialize services
	svCustomer := service.NewCustomersDefault(rpCustomer)
	svProduct := service.NewProductsDefault(rpProduct)
	svInvoice := service.NewInvoicesDefault(rpInvoice)
	svSale := service.NewSalesDefault(rpSale)

	// Define JSON files and their import handlers
	files := []struct {
		Name    string
		Handler func([]byte) error
	}{
		{"customers.json", func(data []byte) error {
			var items []internal.CustomerDTO
			if err := json.Unmarshal(data, &items); err != nil {
				return fmt.Errorf("unmarshal customers: %w", err)
			}
			return svCustomer.Import(items)
		}},

		{"products.json", func(data []byte) error {
			var items []internal.ProductDTO
			if err := json.Unmarshal(data, &items); err != nil {
				return fmt.Errorf("unmarshal products: %w", err)
			}
			return svProduct.Import(items)
		}},

		{"invoices.json", func(data []byte) error {
			var items []internal.InvoiceDTO
			if err := json.Unmarshal(data, &items); err != nil {
				return fmt.Errorf("unmarshal invoices: %w", err)
			}
			return svInvoice.Import(items)
		}},

		{"sales.json", func(data []byte) error {
			var items []internal.SaleDTO
			if err := json.Unmarshal(data, &items); err != nil {
				return fmt.Errorf("unmarshal sales: %w", err)
			}
			return svSale.Import(items)
		}},
	}

	// Loop through each file, read, parse, and import
	for _, f := range files {
		path := filepath.Join(jsonDir, f.Name)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", f.Name, err)
		}

		if err := f.Handler(data); err != nil {
			return fmt.Errorf("importing %s: %w", f.Name, err)
		}

		fmt.Printf("Imported %s successfully\n", f.Name)
	}

	return nil
}
