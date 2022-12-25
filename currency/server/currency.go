package server

import (
	"context"
	protos "kokoichi206/go-expart/currency/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
	protos.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, in *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", in.GetBase(), "destination", in.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}

func (c *Currency) mustEmbedUnimplementedCurrencyServer() {
	c.log.Info("hogee")
}
