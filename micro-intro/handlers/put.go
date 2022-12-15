package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kokoichi206/go-expert/micro-intro/data"
)

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product NOT FOUND", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "xxx", http.StatusInternalServerError)
		return
	}
}
