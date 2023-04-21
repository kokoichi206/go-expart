package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
	cl := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(jsonToReturn)),
			Header:     make(http.Header),
		}
	})
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
