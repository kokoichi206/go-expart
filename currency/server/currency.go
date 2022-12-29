package server

import (
	"context"
	"io"
	"kokoichi206/go-expart/currency/data"
	protos "kokoichi206/go-expart/currency/protos/currency"
	"time"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
	protos.UnimplementedCurrencyServer
	rates *data.ExchangeRates
}

func NewCurrency(l hclog.Logger, e *data.ExchangeRates) *Currency {
	return &Currency{log: l, rates: e}
}

func (c *Currency) GetRate(ctx context.Context, in *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", in.GetBase(), "destination", in.GetDestination())

	rate, err := c.rates.GetRate(in.GetBase().String(), in.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}

func (c *Currency) mustEmbedUnimplementedCurrencyServer() {
	c.log.Info("hogee")
}

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {

	go func() {
		for {
			// Blocking method
			rr, err := src.Recv()
			if err == io.EOF {
				c.log.Info("Client closed connection!")
				break
			}
			if err != nil {
				c.log.Error("Unable to read from client", "error", err)
				break
			}

			c.log.Info("Handle client request", "request", rr)
		}	
	}()

	for {
		err := src.Send(&protos.RateResponse{Rate: 12.1})
		if err != nil {
			return err
		}

		// dummy
		time.Sleep(5 * time.Second)
	}
}
