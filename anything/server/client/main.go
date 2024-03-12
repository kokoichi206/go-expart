package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	msgs := []string{
		"PIEN",
		"PIEN2",
		"PIEN3",
	}

	cur := 0
	var conn net.Conn = nil

	for {
		var err error

		if conn == nil {
			conn, _ = net.Dial("tcp", "localhost:50051")
		}

		req, _ := http.NewRequest("POST", "http://localhost:50051", strings.NewReader(msgs[cur]))
		req.Write(conn)

		res, err := http.ReadResponse(bufio.NewReader(conn), req)
		if err != nil {
			// timeout はここでエラーになる。
			fmt.Println("Retry!")
			conn = nil
			continue
		}

		dump, _ := httputil.DumpResponse(res, true)
		fmt.Println(string(dump))

		cur++
		if cur >= len(msgs) {
			break
		}
	}

	conn.Close()
}
