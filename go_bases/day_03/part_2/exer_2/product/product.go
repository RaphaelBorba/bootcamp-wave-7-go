package product

type Product interface {
	Price() float64
}

type SmallProduct struct {
	Name         string
	DefaultPrice float64
	Amount       int
}

func (sp SmallProduct) Price() float64 {
	return sp.DefaultPrice
}

type MedProduct struct {
	Name         string
	DefaultPrice float64
	Amount       int
}

func (mp MedProduct) Price() float64 {
	return mp.DefaultPrice * 1.06
}

type BigProduct struct {
	Name         string
	DefaultPrice float64
	Amount       int
}

func (bp BigProduct) Price() float64 {
	return (bp.DefaultPrice * 1.06) + 2500
}
