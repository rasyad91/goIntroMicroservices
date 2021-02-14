package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rasyad91/introMicroservices/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the database
// responses:
//	201: noContent

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	// p.l.Println("Handle POST Products")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println(err)
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
	}

	p.l.Printf("Handle DELETE Product %v\n", id)

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
