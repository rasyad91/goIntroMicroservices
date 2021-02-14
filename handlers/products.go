package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rasyad91/introMicroservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{
		l: l,
	}
}

type KeyProduct struct{}

// REFACTORED USING GORILLA
// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		{
// 			p.getProducts(rw, r)
// 			return
// 		}
// 	case http.MethodPost:
// 		{
// 			p.addProduct(rw, r)
// 			return
// 		}
// 	case http.MethodPut:
// 		{
// 			rg := regexp.MustCompile(`/([0-9]+)`)
// 			g := rg.FindAllStringSubmatch(r.URL.Path, -1)
// 			fmt.Printf("Request URL Path: %v\n", r.URL.Path)
// 			fmt.Printf("Regex findallStringSubmatch, g: %v\n", g)
// 			if len(g) != 1 {
// 				fmt.Println(g)
// 				http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 				return
// 			}
// 			if len(g[0]) != 2 {
// 				fmt.Println(g)
// 				http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 				return
// 			}
// 			idString := g[0][1]
// 			id, err := strconv.Atoi(idString)
// 			if err != nil {
// 				http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 				return
// 			}
// 			p.l.Println("got id", id)
// 			p.updateProduct(id, rw, r)
// 			return
// 		}
// 	default:
// 		{
// 			rw.WriteHeader(http.StatusMethodNotAllowed)
// 		}
// 	}
// }

// getProducts returns the products from the data stored
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// Get products from data products => encode to JSON
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// PostProduct add a new product into the data stored
func (p *Products) PostProduct(rw http.ResponseWriter, r *http.Request) {
	// p.l.Println("Handle POST Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("POST Prod: %#v", prod)
	data.AddProduct(&prod)

}

// PutProduct uses UpdateProduct to updates product to database
func (p *Products) PutProduct(rw http.ResponseWriter, r *http.Request) {
	// p.l.Println("Handle POST Products")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println(err)
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
	}
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("PUT Prod: %#v", prod)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

// MiddlewareFromJSON middleware for validates if
func (p *Products) MiddlewareFromJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println(err)
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
