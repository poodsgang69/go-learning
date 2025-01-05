package handlers

import (
	"API/commons"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	// catch any other type of http.method
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling the POST Request")

	newProduct := &commons.Product{}

	err := newProduct.FromJSON(r.Body)
	// p.l.Println(err)
	if err != nil {
		http.Error(rw, "Unable to Un-Marshal Request Body JSON", http.StatusBadRequest)
	}

	p.l.Printf("New Product Added is: %#v", newProduct)

}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling the GET Request")
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
