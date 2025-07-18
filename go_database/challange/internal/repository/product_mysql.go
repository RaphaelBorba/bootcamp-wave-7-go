package repository

import (
	"database/sql"
	"fmt"

	"app/internal"
)

// NewProductsMySQL creates new mysql repository for product entity.
func NewProductsMySQL(db *sql.DB) *ProductsMySQL {
	return &ProductsMySQL{db}
}

// ProductsMySQL is the MySQL repository implementation for product entity.
type ProductsMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all products from the database.
func (r *ProductsMySQL) FindAll() (p []internal.Product, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `description`, `price` FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var pr internal.Product
		// scan the row into the product
		err := rows.Scan(&pr.Id, &pr.Description, &pr.Price)
		if err != nil {
			return nil, err
		}
		// append the product to the slice
		p = append(p, pr)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the product into the database.
func (r *ProductsMySQL) Save(p *internal.Product) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO products (`description`, `price`) VALUES (?, ?)",
		(*p).Description, (*p).Price,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*p).Id = int(id)

	return
}

func (r ProductsMySQL) CreateFromImport(p internal.ProductDTO) error {
	_, err := r.db.Exec(`
        INSERT INTO products (id, description, price)
        VALUES (?, ?, ?)
        ON DUPLICATE KEY UPDATE
          description = VALUES(description),
          price  = VALUES(price)
    `, p.ID, p.Description, p.Price)
	return err
}

func (r *ProductsMySQL) GetTopSellingProducts() ([]internal.ProductQuantity, error) {
	query := `
    SELECT
      p.description,
      SUM(s.quantity) AS total_quantity
    FROM
      products AS p
      JOIN sales AS s ON s.product_id = p.id
    GROUP BY
      p.id
    ORDER BY
      total_quantity DESC
    LIMIT 5;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query top selling products: %w", err)
	}
	defer rows.Close()

	var results []internal.ProductQuantity
	for rows.Next() {
		var pq internal.ProductQuantity
		if err := rows.Scan(&pq.Description, &pq.TotalQuantity); err != nil {
			return nil, fmt.Errorf("scan product quantity: %w", err)
		}
		results = append(results, pq)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate top selling rows: %w", err)
	}

	return results, nil
}
