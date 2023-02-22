package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWSOrderbook(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client to orderbook feed: ", ws.RemoteAddr())

	maxAge := 10
	cnt := 0
	for {
		payload := fmt.Sprintf("orderbook data -> %d\n", time.Now().UnixNano())

		ws.Write([]byte(payload))
		time.Sleep(2 * time.Second)

		cnt++
		if cnt > maxAge {
			ws.Write([]byte("Finished!!"))
			break
		}
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client: ", ws.RemoteAddr())

	// map is not concurrent safe, so usually should use mutex
	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		// 反対から閉じられた時。
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read error: ", err)
			continue
		}

		msg := buf[:n]

		s.broadcast(msg)
		// fmt.Println(string(msg))

		// ws.Write([]byte("Thank you for the message!!!"))
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println(err)
			}
		}(ws)
	}
}

func main() {
	server := NewServer()

	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/orderbook-feed", websocket.Handler(server.handleWSOrderbook))

	http.ListenAndServe(":3333", nil)
}
