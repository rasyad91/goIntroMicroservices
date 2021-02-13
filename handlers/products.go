package handlers

import (
	"log"
	"net/http"

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

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			p.getProducts(rw, r)
			return
		}
	default:
		{
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	// Get products from data products => encode to JSON
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
