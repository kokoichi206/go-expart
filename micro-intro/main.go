package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	// protos "kokoichi206/go-expart/currency/protos/currency"
	protos "kokoichi206/go-expart/currency/protos/currency"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/kokoichi206/go-expert/micro-intro/handlers"
	"google.golang.org/grpc"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// By default, grpc uses http2
	serverAddr := "localhost:9092"
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// create client
	cc := protos.NewCurrencyClient(conn)

	ph := handlers.NewProducts(l, cc)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	s := &http.Server{
		Addr:         ":9091",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminal, graceful shutdown", sig)
	// Timeout Context
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
