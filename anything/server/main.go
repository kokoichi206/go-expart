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
	"time"
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
			defer conn.Close()
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			for {
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))

				req, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					neterr, ok := err.(net.Error)
					if ok && neterr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						fmt.Println("EOF")
						break
					}
					log.Fatalf("Failed to read request: %v", err)
				}

				dump, err := httputil.DumpRequest(req, true)
				if err != nil {
					log.Fatalf("Failed to dump request: %v", err)
				}
				fmt.Println(string(dump))

				body := "Hello, World!\n"

				res := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(body)),
					Body:          io.NopCloser(strings.NewReader(body)),
				}
				res.Write(conn)
			}
		}()
	}
}
