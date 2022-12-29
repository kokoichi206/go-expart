package handlers

import (
	"net/http"

	"github.com/kokoichi206/go-expert/micro-intro/data"
)

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {

	p.l.Debug("Handle Post Product")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(prod)
	p.l.Debug("Prod: %#v", prod)
}
