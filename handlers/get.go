package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rasyad91/introMicroservices/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponseWrapper

// GetProducts returns the products from the data stored
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// Get products from data products => encode to JSON
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// GetProductByID returns the product from the database corresponding to the id
func (p *Products) GetProductByID(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println(err)
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
	}

	p.l.Printf("Handle GET Product by ID %v\n", id)

	gp, err := data.GetProductByID(id)
	if err == data.ErrProductNotFound {
		p.l.Println(err)
		http.Error(rw, "Product not found", http.StatusNoContent)
		return
	}
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = gp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}

}
