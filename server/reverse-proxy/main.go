package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

const (
	cookie = "xxx"
)

func main() {
	director := func(request *http.Request) {
		url := *request.URL
		url.Scheme = "https"
		url.Host = "text.com"

		fmt.Printf("url.String(): %v\n", url.String())

		req, err := http.NewRequest(request.Method, url.String(), request.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		header := request.Header
		header.Set("Cookie", cookie)
		// これつけると 401 から 403 になった！
		header.Add("authority", "test.com")
		header.Del("origin")
		header.Del("referer")

		req.Header = header
		*request = *req
		fmt.Printf("req: %v\n", req)
		fmt.Printf("req.URL.Path: %v\n", req.URL.Path)

		fmt.Println("director called")
	}
	modifyResponse := func(r *http.Response) error {
		r.Header.Set("Access-Control-Allow-Origin", "http://localhost:8080")
		r.Header.Set("Access-Control-Expose", "*")
		r.Header.Set("Access-Control-Allow-Headers", "*")
		r.Header.Set("Access-Control-Allow-Credentials", "true")

		fmt.Printf("r.Status: %v\n", r.Status)
		fmt.Printf("r: %v\n", r)

		// teeReader := io.TeeReader(r.Body, os.Stdout)
		// io.Copy(os.Stdout, r.Body)
		return nil
	}

	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifyResponse,
	}
	server := http.Server{
		Addr: ":7777",
		// Addr:    ":9000",
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
