package handlers

import (
	"log"
	"module/new/directory/API/commons"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var allProducts commons.Products = commons.GetProducts()
	/*
		One way to convert product Struct to JSON is by using json.Marshal. But with this, we need to write it to a responseWriter (or any io.Writer).
		To get over this, we use json.NewEncoder(io.Writer) [registering a writer to the newEncoder] and json.Encode(type struct) to encode and write it to the writer directly.
	*/
	// jsonData, err := json.Marshal(allProducts)
	var err error = allProducts.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal the Product Strut -> JSON", http.StatusInternalServerError)
	}

	// We dont need this as we will be writing directly to the ResponseWriter using the NewEncoder
	// rw.Write(jsonData)
}