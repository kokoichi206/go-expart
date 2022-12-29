package handlers

import (
	"net/http"
)

// swagger:route GET /products listProducts
// Returns list of products
// responses:
//  200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Debug("Handle Get Products")

	cur := r.URL.Query().Get("currency")

	lp, err := p.productDB.GetProducts(cur)

	// 失敗した時に書き込まれる可能性はない？
	err = lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
