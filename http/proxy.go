package main

import (
	"fmt"
	"net/http"
)

func proxy() {
	tp := http.Transport{
		// Default setting?
		Proxy: http.ProxyFromEnvironment,
	}

	client := http.Client{
		Transport: &tp,
	}

	fmt.Printf("client: %v\n", client)

	// DefaultTransport is the default implementation of Transport and is
	// used by DefaultClient. It establishes network connections as needed
	// and caches them for reuse by subsequent calls. It uses HTTP proxies
	// as directed by the environment variables HTTP_PROXY, HTTPS_PROXY
	// and NO_PROXY (or the lowercase versions thereof).
	fmt.Printf("http.DefaultTransport: %v\n", http.DefaultTransport)
}
