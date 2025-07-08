package repository

import (
	"app/internal"
	"database/sql"
	"fmt"
	"log"
)

// NewProductsMySQL returns a new instance of ProductsMySQL
func NewProductsMySQL(db *sql.DB) *ProductsMySQL {
	return &ProductsMySQL{
		db: db,
	}
}

// ProductsMySQL is a struct that represents a product repository
type ProductsMySQL struct {
	// db is the database connection
	db *sql.DB
}

func (r *ProductsMySQL) GetAll() (products []internal.Product, err error) {
	const query = `
        SELECT id, name, quantity, code_value, is_published, expiration, price
        FROM products
    `

	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("[GetAll][MySQL] query error: %v", err)
		return nil, fmt.Errorf("querying all products: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p internal.Product
		if err := rows.Scan(
			&p.ID, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price,
		); err != nil {
			log.Printf("[GetAll][MySQL] scan error: %v", err)
			return nil, fmt.Errorf("scanning products: %w", err)
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		log.Printf("[GetAll][MySQL] rows iteration error: %v", err)
		return nil, fmt.Errorf("iterating products rows: %w", err)
	}

	return products, nil
}

// GetOne returns a product by id
func (r *ProductsMySQL) GetOne(id int) (p internal.Product, err error) {
	// execute the query
	row := r.db.QueryRow(
		"SELECT `id`, `name`, `quantity`, `code_value`, `is_published`, `expiration`, `price` "+
			"FROM `products` WHERE `id` = ?",
		id,
	)
	if err = row.Err(); err != nil {
		return
	}

	// scan the row into the product
	err = row.Scan(&p.ID, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrProductNotFound
		}
		return
	}

	return
}

// Store stores a product
func (r *ProductsMySQL) Store(p *internal.Product) (err error) {

	// execute the query
	result, err := r.db.Exec(
		"INSERT INTO `products` (`name`, `quantity`, `code_value`, `is_published`, `expiration`, `price`, `id_warehouse`) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?)",
		p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.IdWarehouse,
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	// get the last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	p.ID = int(id)

	return
}

// Update updates a product
func (r *ProductsMySQL) Update(p *internal.Product) (err error) {
	// execute the query
	_, err = r.db.Exec(
		"UPDATE `products` SET `name` = ?, `quantity` = ?, `code_value` = ?, `is_published` = ?, `expiration` = ?, `price` = ? "+
			"WHERE `id` = ?",
		p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.ID,
	)
	if err != nil {
		return
	}

	return
}

// Delete deletes a product by id
func (r *ProductsMySQL) Delete(id int) (err error) {
	// execute the query
	_, err = r.db.Exec(
		"DELETE FROM `products` WHERE `id` = ?",
		id,
	)
	if err != nil {
		return
	}

	return
}
