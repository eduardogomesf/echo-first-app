package entities

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}
