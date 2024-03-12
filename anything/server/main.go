package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Server is running on port 50051...")

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("Failed to accept: %v", err)
		}

		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			req, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				log.Fatalf("Failed to read request: %v", err)
			}

			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				log.Fatalf("Failed to dump request: %v", err)
			}
			fmt.Println(string(dump))

			res := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       io.NopCloser(strings.NewReader("Hello, World!\n")),
			}
			res.Write(conn)
			conn.Close()
		}()
	}
}
