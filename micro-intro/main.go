package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	// protos "kokoichi206/go-expart/currency/protos/currency"
	protos "kokoichi206/go-expart/currency/protos/currency"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/kokoichi206/go-expert/micro-intro/data"
	"github.com/kokoichi206/go-expert/micro-intro/handlers"
	"google.golang.org/grpc"
)

func main() {

	hl := hclog.Default()

	// By default, grpc uses http2
	serverAddr := "localhost:9092"
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// create client
	cc := protos.NewCurrencyClient(conn)

	db := data.NewProductDB(cc, hl)

	ph := handlers.NewProducts(hl, db)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)
	getRouter.HandleFunc("/products", ph.GetProducts).Queries("currency", "{[A-Z]{3}}")

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

	// hclog は standardlog に変更可能。
	s := &http.Server{
		Addr:         ":9091",
		Handler:      sm,
		ErrorLog:     hl.StandardLogger(&hclog.StandardLoggerOptions{}),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			hl.Error(err.Error())
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	hl.Debug("Recieved terminal, graceful shutdown", sig)
	// Timeout Context
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
