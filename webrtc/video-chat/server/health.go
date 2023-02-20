package server

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := json.Marshal(struct {
		Health string `json:"health"`
	}{
		Health: "ok",
	})
	w.Write(body)
}
