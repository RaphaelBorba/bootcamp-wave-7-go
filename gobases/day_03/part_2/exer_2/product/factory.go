package product

func NewProduct(productType string, price float64) Product {
	switch productType {
	case "small":
		return SmallProduct{DefaultPrice: price}
	case "medium":
		return MedProduct{DefaultPrice: price}
	case "large":
		return BigProduct{DefaultPrice: price}
	default:
		return nil // ou lançar erro, dependendo da estratégia
	}
}
