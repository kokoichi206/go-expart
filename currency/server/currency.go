package server

import (
	"context"
	"kokoichi206/go-expart/currency/data"
	protos "kokoichi206/go-expart/currency/protos/currency"

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
