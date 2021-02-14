package handlers

import (
	"net/http"

	"github.com/rasyad91/introMicroservices/data"
)

// swagger:route POST / products createProduct
// Creates a product from the database
// responses:
//	200: productsResponseWrapper

// PostProduct add a new product into the data stored
func (p *Products) PostProduct(rw http.ResponseWriter, r *http.Request) {
	// p.l.Println("Handle POST Products")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("POST Prod: %#v", prod)
	data.AddProduct(prod)
}
