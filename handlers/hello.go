package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello struct
type Hello struct {
	l *log.Logger
}

// NewHello : Idiomatic principles of Go code - to create newHello and defining it
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	//replaced log.Println("Hello World")
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Opps", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s\n", d)
}
