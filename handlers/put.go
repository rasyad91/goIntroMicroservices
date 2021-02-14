package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rasyad91/introMicroservices/data"
)

// swagger:route PUT /products/{id} products updateProduct
// Updates a product from the database
// responses:
//	200: productsResponseWrapper

// PutProduct uses UpdateProduct to updates product to database
func (p *Products) PutProduct(rw http.ResponseWriter, r *http.Request) {
	// p.l.Println("Handle POST Products")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println(err)
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
	}
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("PUT Prod: %#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
