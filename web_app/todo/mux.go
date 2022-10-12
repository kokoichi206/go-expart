package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/kokoichi206/go-expert/web/todo/clock"
	"github.com/kokoichi206/go-expert/web/todo/config"
	"github.com/kokoichi206/go-expert/web/todo/handler"
	"github.com/kokoichi206/go-expert/web/todo/service"
	"github.com/kokoichi206/go-expert/web/todo/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	// mux stands for HTTP request multiplexer
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	// validator とかを new するのって、NewMux(ルーティイングを初期化)の役割なんだっけ？
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{
		Service:   &service.AddTask{DB: db, Repo: &r},
		Validator: v,
	}

	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{
		Service: &service.ListTask{DB: db, Repo: &r},
	}
	mux.Get("/tasks", lt.ServeHTTP)

	return mux, cleanup, nil
}
