package repository

import (
	"app/internal"
	"database/sql"
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
