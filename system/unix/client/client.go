package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-server-sample")
	fmt.Printf("path: %v\n", path)

	conn, err := net.Dial("unix", path)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("get", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}

	req.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
