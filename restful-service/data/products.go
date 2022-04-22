package data

import (
	"encoding/json"
	"fmt"
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

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	n := len(productList) - 1
	return productList[n].ID + 1
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

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}
