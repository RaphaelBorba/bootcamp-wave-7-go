package repository

import (
	"app/internal"
	"database/sql"
	"fmt"
	"log"
)

func NewWarehouseMySQL(db *sql.DB) *WarehouseMySQL {
	return &WarehouseMySQL{
		db: db,
	}
}

type WarehouseMySQL struct {
	db *sql.DB
}

func (r *WarehouseMySQL) GetAllWarehouses() (warehouses []internal.Warehouse, err error) {
	const query = `
        SELECT id, name, adress, telephone, capacity
        FROM warehouses
    `

	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("[GetAllWarehouses][MySQL] query error: %v", err)
		return nil, fmt.Errorf("querying all warehouses: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var w internal.Warehouse
		if err := rows.Scan(
			&w.ID,
			&w.Name,
			&w.Adress,
			&w.Telephone,
			&w.Capacity,
		); err != nil {
			log.Printf("[GetAllWarehouses][MySQL] scan error: %v", err)
			return nil, fmt.Errorf("scanning warehouse: %w", err)
		}
		warehouses = append(warehouses, w)
	}
	if err := rows.Err(); err != nil {
		log.Printf("[GetAllWarehouses][MySQL] rows iteration error: %v", err)
		return nil, fmt.Errorf("iterating warehouse rows: %w", err)
	}

	return warehouses, nil
}

func (r *WarehouseMySQL) GetOneWarehouse(id int) (w internal.Warehouse, err error) {
	row := r.db.QueryRow(
		"SELECT `id`, `name`, `adress`, `telephone`, `capacity`"+
			"FROM `warehouses` WHERE `id` = ?",
		id,
	)
	if err = row.Err(); err != nil {
		return
	}

	err = row.Scan(&w.ID, &w.Name, &w.Adress, &w.Telephone, &w.Capacity)
	if err != nil {
		if err == sql.ErrNoRows {
			err = internal.ErrProductNotFound
		}
		return
	}

	return
}

func (r *WarehouseMySQL) StoreWarehouse(w *internal.Warehouse) (err error) {

	result, err := r.db.Exec(
		"INSERT INTO `warehouses` (`name`, `adress`, `telephone`, `capacity`) "+
			"VALUES (?, ?, ?, ?)",
		w.Name, w.Adress, w.Telephone, w.Capacity,
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	w.ID = int(id)

	return
}

func (r *WarehouseMySQL) UpdateWarehouse(w *internal.Warehouse) (err error) {
	_, err = r.db.Exec(
		"UPDATE `warehouses` SET `name` = ?, `adress` = ?, `telephone` = ?, `capacity` = ? "+
			"WHERE `id` = ?",
		w.Name, w.Adress, w.Telephone, w.Capacity, w.ID,
	)
	if err != nil {
		return
	}

	return
}

func (r *WarehouseMySQL) DeleteWarehouse(id int) (err error) {
	_, err = r.db.Exec(
		"DELETE FROM `warehouses` WHERE `id` = ?",
		id,
	)
	if err != nil {
		return
	}

	return
}

func (r *WarehouseMySQL) GetWarehouseReport(warehouseID *int) (reports []internal.WarehouseReport, err error) {
	query := `
        SELECT w.name,
               COUNT(p.id) AS product_count
        FROM warehouses AS w
        LEFT JOIN products AS p
          ON p.id_warehouse = w.id
    `
	args := []any{}

	if warehouseID != nil {
		query += " WHERE w.id = ?"
		args = append(args, *warehouseID)
	}

	query += " GROUP BY w.name"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Printf("[GetWarehouseReport][MySQL] query error: %v, args=%v", err, args)
		return nil, fmt.Errorf("querying warehouse report: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rpt internal.WarehouseReport
		if err := rows.Scan(&rpt.Name, &rpt.ProductCount); err != nil {
			log.Printf("[GetWarehouseReport][MySQL] scan error: %v", err)
			return nil, fmt.Errorf("scanning warehouse report: %w", err)
		}
		reports = append(reports, rpt)
	}
	if err := rows.Err(); err != nil {
		log.Printf("[GetWarehouseReport][MySQL] rows iteration error: %v", err)
		return nil, fmt.Errorf("iterating rows: %w", err)
	}

	return reports, nil
}
