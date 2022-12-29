package main

import (
	"kokoichi206/go-expart/currency/data"
	protos "kokoichi206/go-expart/currency/protos/currency"
	"kokoichi206/go-expart/currency/server"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRates(log)
	if err != nil {
		log.Error("Unable to generate rates", "error: ", err)
		os.Exit(1)
	}

	gs := grpc.NewServer()
	cs := server.NewCurrency(log, rates)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
