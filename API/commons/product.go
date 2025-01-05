package commons

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	ModifiedOn  string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// function to return all the products
func GetProducts() []*Product {
	return productList
}

func (p *Products) ToJSON(w io.Writer) error {
	jsonEncoder := json.NewEncoder(w)
	return jsonEncoder.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	jsonDecoder := json.NewDecoder(r)
	return jsonDecoder.Decode(p)
}

var productList []*Product = []*Product{
	&Product{
		ID:          001,
		Name:        "Espresso",
		Description: "Bitter Coffee",
		Price:       9.99,
		SKU:         "esp001",
		CreatedOn:   time.Now().UTC().String(),
		ModifiedOn:  time.Now().UTC().String(),
	},
	&Product{
		ID:          002,
		Name:        "Latte",
		Description: "Frothy Coffee",
		Price:       12.99,
		SKU:         "lat001",
		CreatedOn:   time.Now().UTC().String(),
		ModifiedOn:  time.Now().UTC().String(),
	},
}
