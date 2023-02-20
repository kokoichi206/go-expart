package main

import (
	"log"
	"net/http"

	"video-chat/server"
)

func main() {

	server.AllRooms.Init()

	http.HandleFunc("/health", server.HealthHandler)
	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
