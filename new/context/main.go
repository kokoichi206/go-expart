package main

import (
	"context"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	<-ctx.Done()
	fmt.Println("handler end")
	return
}

func handlerWithout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithoutCancel(ctx)
	<-ctx.Done()
	fmt.Println("handlerWithout end")
	return
}

func main() {
	http.HandleFunc("/handler", handler)
	http.HandleFunc("/handler1", handlerWithout)
	http.ListenAndServe(":8818", nil)
}
