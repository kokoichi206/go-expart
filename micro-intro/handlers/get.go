package handlers

import (
	"context"
	protos "kokoichi206/go-expart/currency/protos/currency"
	"net/http"

	"github.com/kokoichi206/go-expert/micro-intro/data"
)

// swagger:route GET /products listProducts
// Returns list of products
// responses:
//  200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Get Products")

	lp := data.GetProducts()

	// get exchange rate
	rr := &protos.RateRequest{
		Base:        protos.Currencies_JPY,
		Destination: protos.Currencies_USD,
	}
	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Println("[Error] error getting new rate", err)
		return
	}
	// apply rate
	for _, prod := range lp {
		prod.Price = prod.Price * resp.Rate
	}

	// 失敗した時に書き込まれる可能性はない？
	err = lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
