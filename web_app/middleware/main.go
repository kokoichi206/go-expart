package main

import (
	"log"
	"net/http"
	"time"
)

func TimeMeasurement(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		h.ServeHTTP(w, r)
		d := time.Now().Sub(s).Milliseconds()
		log.Printf("end %s(%d ms)\n", time.Now().Format(time.RFC3339), d)
	})
}

// 追加情報を利用したミドルウェアパターンの実装
//
// vmw := VersioinAdder("1.0.1")
// http.handle("/users", vmw(userHandler))
func VersioinAdder(v string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Add("App-Version", v)
			next.ServeHTTP(w, r)
		})
	}
}
