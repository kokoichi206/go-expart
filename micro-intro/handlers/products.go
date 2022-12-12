package handlers

import (
	"log"
	"net/http"

	"github.com/kokoichi206/go-expert/micro-intro/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	// Catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Get Products")

	lp := data.GetProducts()

	// 失敗した時に書き込まれる可能性はない？
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Post Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	data.AddProduct(prod)
	p.l.Printf("Prod: %#v", prod)
}
