package example

import (
	"errors"
	"log"
	"net"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	transport http.RoundTripper
	logger    func(string, ...interface{})
}

func (lt *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if lt.logger == nil {
		lt.logger = log.Printf
	}

	start := time.Now()
	resp, err := lt.transport.RoundTrip(req)

	if resp != nil {
		lt.logger("%s %s %d %s, duration: %d", req.Method, req.URL.String(), resp.StatusCode, http.StatusText(resp.StatusCode), time.Since(start))
	}

	return resp, err
}

type basicAuthRoundTripper struct {
	username string
	password string
	base     http.RoundTripper
}

func (ba *basicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(ba.username, ba.password)
	return ba.base.RoundTrip(req)
}

type retryableRoundTripper struct {
	base     http.RoundTripper
	attempts int
	waitTime time.Duration
}

func (r *retryableRoundTripper) shouldRetry(resp *http.Response, err error) bool {
	// ネットワークエラーのリトライ
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return true
		}
	}

	// レスポンスコードによるリトライ（通信自体は成功してるため err は nil で返ってくる）
	if resp != nil {
		if resp.StatusCode == 429 || (500 <= resp.StatusCode && resp.StatusCode <= 504) {
			return true
		}
	}
	return false
}

func (r *retryableRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var (
		resp *http.Response
		err  error
	)
	for count := 0; count < r.attempts; count++ {
		resp, err = r.base.RoundTrip(req)

		if !r.shouldRetry(resp, err) {
			return resp, err
		}

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(r.waitTime): // リトライのために待機
		}
	}

	return resp, err
}
