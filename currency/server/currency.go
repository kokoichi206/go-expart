package server

import (
	"context"
	"io"
	"kokoichi206/go-expart/currency/data"
	protos "kokoichi206/go-expart/currency/protos/currency"
	"time"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Currency struct {
	log hclog.Logger
	protos.UnimplementedCurrencyServer
	rates         *data.ExchangeRates
	subscriptions map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest
}

func NewCurrency(l hclog.Logger, e *data.ExchangeRates) *Currency {
	c := &Currency{
		log:           l,
		rates:         e,
		subscriptions: make(map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest),
	}
	go c.handleUpdates()

	return c
}

func (c *Currency) GetRate(ctx context.Context, in *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", in.GetBase(), "destination", in.GetDestination())

	if in.Base == in.Destination {
		// grpc/status
		// status.Errorf, status.Newf ...etc
		err := status.Newf(
			codes.InvalidArgument,
			"Base currency %s cannot be the same as the destination currency %s",
			in.Base.String(), in.Destination.String(),
		)
		err, wde := err.WithDetails(in)
		if wde != nil {
			return nil, wde
		}

		return nil, err.Err()
	}

	rate, err := c.rates.GetRate(in.GetBase().String(), in.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)
	for range ru {
		c.log.Info("Got updated rates")

		// loop over subscribed clients
		for k, v := range c.subscriptions {
			// loop over subscribed rates
			for _, rr := range v {
				r, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
				if err != nil {
					c.log.Error("Unable to get update rate", "base", rr.GetBase().String())
				}

				err = k.Send(
					&protos.StreamingRateResponse{
						Message: &protos.StreamingRateResponse_RateResponse{
							RateResponse: &protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: r},
						},
					},
				)
				if err != nil {
					c.log.Error("Unable to send updated rate")
				}
			}
		}
	}
}

func (c *Currency) mustEmbedUnimplementedCurrencyServer() {
	c.log.Info("hogee")
}

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {

	for {
		// Blocking method
		rr, err := src.Recv()
		if err == io.EOF {
			c.log.Info("Client closed connection!")
			break
		}
		if err != nil {
			c.log.Error("Unable to read from client", "error", err)
			return err
		}

		c.log.Info("Handle client request", "request", rr)

		rrs, ok := c.subscriptions[src]
		if !ok {
			rrs = []*protos.RateRequest{}
		}

		var validationError *status.Status
		for _, v := range rrs {
			if v.Base == rr.Base && v.Destination == rr.Destination {
				validationError = status.Newf(
					codes.AlreadyExists,
					"Unable to subscribe for currency as subscription already exists",
				)
				validationError, err = validationError.WithDetails(rr)
				if err != nil {
					c.log.Error("Unable to add metadata to error", err)
					break

				}

				if validationError != nil {
					src.Send(
						&protos.StreamingRateResponse{
							Message: &protos.StreamingRateResponse_Error{
								Error: validationError.Proto(),
							},
						},
					)
					continue
				}
			}
		}

		// all ok
		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}

	return nil
}
