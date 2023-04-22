package main

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	testApp.App = test.NewApp()
	testApp.HTTPClient = cl

	os.Exit(m.Run())
}

var jsonToReturn = `
{"ts":1682092314357,"tsj":1682092305951,"date":"Apr 21st 2023, 11:51:45 am NY","items":[{"curr":"USD","xauPrice":1975.3225,"xagPrice":24.9656,"chgXau":-29.9425,"chgXag":-0.2889,"pcXau":-1.4932,"pcXag":-1.144,"xauClose":2005.265,"xagClose":25.2545}]}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var cl = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(jsonToReturn)),
		Header:     make(http.Header),
	}
})
