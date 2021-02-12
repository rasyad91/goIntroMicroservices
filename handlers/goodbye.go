package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GoodBye struct
type GoodBye struct {
	l *log.Logger
}

// NewGoodbye : Idiomatic principles of Go code - to create NewGoodbye and defining it
func NewGoodbye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	g.l.Println("Goodbye World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Opps", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Goodbye %s\n", d)
}
