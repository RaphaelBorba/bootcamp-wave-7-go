package repository

import (
	"database/sql"
	"fmt"

	"app/internal"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
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
	(*c).Id = int(id)

	return
}

func (r CustomersMySQL) CreateFromImport(c internal.CustomerDTO) error {
	_, err := r.db.Exec(`
        INSERT INTO customers (id, first_name, last_name, `+"`condition`"+`)
        VALUES (?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE
          first_name = VALUES(first_name),
          last_name  = VALUES(last_name),
          `+"`condition`"+`  = VALUES(`+"`condition`"+`)
    `, c.ID, c.FirstName, c.LastName, c.Condition)
	return err
}

func (r *CustomersMySQL) GetTotalSpentByCondition() ([]internal.ConditionTotal, error) {
	query := "SELECT c.`condition`, ROUND(SUM(s.quantity * p.price), 2) AS total_spent FROM customers AS c JOIN invoices  AS i ON i.customer_id = c.id JOIN sales     AS s ON s.invoice_id  = i.id JOIN products  AS p ON p.id = s.product_id WHERE c.`condition` IN (0,1) GROUP BY c.`condition` ORDER BY c.`condition`"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query total spent by condition: %w", err)
	}
	defer rows.Close()

	results := make([]internal.ConditionTotal, 0)
	for rows.Next() {
		var ct internal.ConditionTotal
		if err := rows.Scan(&ct.Condition, &ct.TotalSpent); err != nil {
			return nil, fmt.Errorf("scan total spent row: %w", err)
		}
		results = append(results, ct)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate total spent rows: %w", err)
	}

	return results, nil
}

func (r *CustomersMySQL) GetTopSpenders() ([]internal.CustomerSpender, error) {
	query := `
    SELECT
      c.first_name,
      c.last_name,
      ROUND(SUM(s.quantity * p.price), 2) AS total_spent
    FROM customers AS c
      JOIN invoices  AS i ON i.customer_id = c.id
      JOIN sales     AS s ON s.invoice_id  = i.id
      JOIN products  AS p ON p.id         = s.product_id 
    GROUP BY
      c.first_name,
      c.last_name
    ORDER BY
      total_spent DESC
    LIMIT 5;
    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query top spenders: %w", err)
	}
	defer rows.Close()

	results := make([]internal.CustomerSpender, 0)
	for rows.Next() {
		var cs internal.CustomerSpender
		if err := rows.Scan(&cs.FirstName, &cs.LastName, &cs.TotalSpent); err != nil {
			return nil, fmt.Errorf("scan top spender row: %w", err)
		}
		results = append(results, cs)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate top spender rows: %w", err)
	}

	return results, nil
}
