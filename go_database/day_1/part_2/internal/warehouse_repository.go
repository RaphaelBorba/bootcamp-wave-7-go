package internal

type RepositoryWarehouse interface {
	GetOneWarehouse(id int) (p Warehouse, err error)
	StoreWarehouse(p *Warehouse) (err error)
	UpdateWarehouse(p *Warehouse) (err error)
	DeleteWarehouse(id int) (err error)
	GetWarehouseReport(warehouseID *int) (reports []WarehouseReport, err error)
	GetAllWarehouses() (warehouses []Warehouse, err error)
}
