package main

import (
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
	g := Gold{
		Client: cl,
		Prices: nil,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	if p.Price != 1975.3225 {
		t.Errorf("expected %f, got %f", 1975.3225, p.Price)
	}
}
