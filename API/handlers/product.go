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

	if r.Method == http.MethodPut {
		p.editProduct(rw, r)
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

	/*
		The method written below was written because the AddProductsToProductList was bound to the Product type so
		only objects/variables of this type can access this method.
	*/
	// newProduct.AddProductToProductList(newProduct)

	/*
		Over here i modified the method to be bound to the module and not to the object being used.
		This is better coding standard as we dont need to glue the method to the object and pass it again to the module to process.
	*/
	commons.AddProductToProductList(newProduct)

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
		http.Error(rw, "Unable to marshal the Product Struct -> JSON", http.StatusInternalServerError)
	}

	// We dont need this as we will be writing directly to the ResponseWriter using the NewEncoder
	// rw.Write(jsonData)
}

func (p *Product) editProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling PUT Request")

	editedProduct := &commons.Product{}

	err := editedProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Un-Marshal Request Body JSON", http.StatusBadRequest)
	}

	var requiredId int = editedProduct.ID

	// Check whether the ID in the URL is present in the ProductsList
	var productIndex int = commons.IsIdPresent(requiredId)

	if productIndex == -1 {
		http.Error(rw, "ID not found in productsList. Please use another ID.", http.StatusBadRequest)
		return
	}

	commons.EditProductInProductList(editedProduct, productIndex)

	p.l.Printf("Products List edited with %#v", editedProduct)
}
