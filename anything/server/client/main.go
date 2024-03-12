package main

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:50051")

	req, _ := http.NewRequest("GET", "http://localhost:50051", nil)

	req.Write(conn)
	res, _ := http.ReadResponse(bufio.NewReader(conn), req)

	dump, _ := httputil.DumpResponse(res, true)
	println(string(dump))

}
