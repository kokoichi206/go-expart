package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/kokoichi206/go-expert/web/todo/auth"
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
	clocker := clock.RealClocker{}
	r := store.Repository{Clocker: clocker}
	rcli, err := store.NewKVS(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	jwter, err := auth.NewJWTer(rcli, clocker)
	if err != nil {
		return nil, cleanup, err
	}

	l := &handler.Login{
		Service: &service.Login{
			DB:             db,
			Repo:           &r,
			TokenGenerator: jwter,
		},
		Validator: v,
	}
	mux.Post("/login", l.ServeHTTP)

	at := &handler.AddTask{
		Service:   &service.AddTask{DB: db, Repo: &r},
		Validator: v,
	}

	lt := &handler.ListTask{
		Service: &service.ListTask{DB: db, Repo: &r},
	}

	mux.Route("/tasks", func(r chi.Router) {
		r.Use(handler.AuthMiddleware(jwter))
		r.Post("/", at.ServeHTTP)
		r.Get("/", lt.ServeHTTP)
	})

	ru := &handler.RegisterUser{
		Service: &service.RegisterUser{
			DB:   db,
			Repo: &r,
		},
		Validator: v,
	}
	mux.Post("/register", ru.ServeHTTP)

	return mux, cleanup, nil
}
