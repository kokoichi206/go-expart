package main_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/magiconair/properties/assert"
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

func TestProxyWS(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Fatal(err)

			return
		}

		wsutil.WriteServerText(conn, []byte("hello"))

		conn.Close()
	}))
	defer ts.Close()

	// Proxy settings (used in client > Transport > proxy?)
	os.Setenv("HTTP_PROXY", ts.URL)

	conn, _, _, err := ws.Dial(context.Background(), "ws://my.domain.com:8930")
	if err != nil {
		log.Fatal(err)

		return
	}

	b, err := wsutil.ReadServerText(conn)
	assert.Equal(t, string(b), "hello")
}
