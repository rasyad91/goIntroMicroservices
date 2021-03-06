package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/rasyad91/introMicroservices/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create new handler
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handler
	//sm := http.NewServeMux()
	sm := mux.NewRouter()
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetProductByID)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.Use(ph.MiddlewareValidateProduct)
	putRouter.HandleFunc("/{id:[0-9]+}", ph.PutProduct)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.Use(ph.MiddlewareValidateProduct)
	postRouter.HandleFunc("/", ph.PostProduct)

	deleteRouter := sm.Methods("DELETE").Subrouter()
	deleteRouter.Use(ph.MiddlewareValidateProduct)
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct)

	// handler for documentation
	// uses redocs middleware
	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// why an address to the server
	// create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal()
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown ", sig)
	tc, c := context.WithTimeout(context.Background(), 30*time.Second)
	_ = c
	s.Shutdown(tc)
}
