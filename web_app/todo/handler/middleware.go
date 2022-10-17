package handler

import (
	"net/http"

	"github.com/kokoichi206/go-expert/web/todo/auth"
)

func AuthMiddleware(j *auth.JWTer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			req, err := j.FillContext(r)
			if err != nil {
				RespondJSON(r.Context(), w, ErrResponse{
					Message: "not find auth information",
					Details: []string{err.Error()},
				}, http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}

// context.Context にユーザー情報が含まれていることを前提とした構造
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !auth.IsAdmin(r.Context()) {
			RespondJSON(r.Context(), w, ErrResponse{
				Message: "not admin",
			}, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
