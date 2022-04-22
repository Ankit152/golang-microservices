package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"desc"`
	Price       float32 `json:"price"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Coffe",
		Description: "Simple coffe",
		Price:       50,
	},
	&Product{
		ID:          2,
		Name:        "Tea",
		Description: "Simple Tea",
		Price:       40,
	},
	&Product{
		ID:          3,
		Name:        "Green Tea",
		Description: "Good for health",
		Price:       50,
	},
	&Product{
		ID:          4,
		Name:        "Espresso",
		Description: "High caffene coffe",
		Price:       70,
	},
}
