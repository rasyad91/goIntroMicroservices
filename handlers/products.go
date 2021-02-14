package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/rasyad91/introMicroservices/data"
)

// A list of products returns in the response

// Products struct
type Products struct {
	l *log.Logger
}

// NewProducts creates new list of Product for handler
func NewProducts(l *log.Logger) *Products {
	return &Products{
		l: l,
	}
}

// KeyProduct for reflection
type KeyProduct struct{}

// MiddlewareValidateProduct middleware for validates if
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println(err)
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] invalid product")
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %v\n", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

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
