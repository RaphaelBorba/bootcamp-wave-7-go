// internal/repository/product_store_mysql.go
package repository

import (
	"app/internal"
	"context"
	"database/sql"
	"errors"
)

// RepositoryProductStore is a repository for products backed by MySQL.
type RepositoryProductStore struct {
	db *sql.DB
}

// NewRepositoryProductStore creates a new repository for products.
func NewRepositoryProductMySQL(db *sql.DB) *RepositoryProductStore {
	return &RepositoryProductStore{db: db}
}

// FindById finds a product by id.
func (r *RepositoryProductStore) FindById(id int) (p internal.Product, err error) {
	row := r.db.QueryRowContext(context.Background(), `
        SELECT id, name, quantity, code_value, is_published, expiration, price
          FROM products
         WHERE id = ?
    `, id)
	if err = row.Scan(
		&p.Id,
		&p.Name,
		&p.Quantity,
		&p.CodeValue,
		&p.IsPublished,
		&p.Expiration,
		&p.Price,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = internal.ErrRepositoryProductNotFound
		}
	}
	return
}

// Save inserts a new product and sets its Id.
func (r *RepositoryProductStore) Save(p *internal.Product) (err error) {
	res, err := r.db.ExecContext(context.Background(), `
        INSERT INTO products
            (name, quantity, code_value, is_published, expiration, price)
        VALUES (?, ?, ?, ?, ?, ?)
    `,
		p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price,
	)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	p.Id = int(lastID)
	return
}

// UpdateOrSave updates if p.Id exists, otherwise inserts.
func (r *RepositoryProductStore) UpdateOrSave(p *internal.Product) (err error) {
	if p.Id > 0 {
		// try update first
		res, err := r.db.ExecContext(context.Background(), `
            UPDATE products
               SET name=?, quantity=?, code_value=?, is_published=?, expiration=?, price=?
             WHERE id=?
        `,
			p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.Id,
		)
		if err != nil {
			return err
		}
		if rows, _ := res.RowsAffected(); rows > 0 {
			return nil
		}
		// if no rows affected, fall through to Save
	}
	return r.Save(p)
}

// Update updates an existing product or returns ErrRepositoryProductNotFound.
func (r *RepositoryProductStore) Update(p *internal.Product) (err error) {
	res, err := r.db.ExecContext(context.Background(), `
        UPDATE products
           SET name=?, quantity=?, code_value=?, is_published=?, expiration=?, price=?
         WHERE id=?
    `,
		p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.Id,
	)
	if err != nil {
		return
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		err = internal.ErrRepositoryProductNotFound
	}
	return
}

// Delete deletes a product or returns ErrRepositoryProductNotFound.
func (r *RepositoryProductStore) Delete(id int) (err error) {
	res, err := r.db.ExecContext(context.Background(), `
        DELETE FROM products
         WHERE id=?
    `, id)
	if err != nil {
		return
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		err = internal.ErrRepositoryProductNotFound
	}
	return
}
