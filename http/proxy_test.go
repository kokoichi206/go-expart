package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestProxy(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		w.Write([]byte("hello"))
	}))
	defer ts.Close()

	// Proxy settings (used in client > Transport > proxy?)
	os.Setenv("HTTP_PROXY", ts.URL)

	res, err := http.Get("http://my.domain.com:8930")
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Error("a response code is not 200")
	}
	if string(body) != "hello" {
		t.Error("a response is not hello")
	}
}
