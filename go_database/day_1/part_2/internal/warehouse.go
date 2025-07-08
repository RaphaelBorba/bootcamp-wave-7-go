package internal

type Warehouse struct {
	ID        int
	Name      string
	Adress    string
	Telephone string
	Capacity  int
}

type WarehouseReport struct {
	Name         string
	ProductCount int
}
