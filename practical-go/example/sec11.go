package example

import (
	"context"
	"fmt"
	"net/http"
)

type myRoundTripper struct {
	base http.RoundTripper
}

// http.RoundTripper interface を満たすようにする
func (m myRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// <<リクエストの前処理>>
	resp, err := m.base.RoundTrip(req)
	// <<リクエストの後処理>>
	return resp, err
}

func roundTripTest(ctx context.Context) {

	client := http.Client{
		Transport: &myRoundTripper{
			base: http.DefaultTransport,
		},
	}
	req, err := http.NewRequestWithContext(ctx, "GET", "https://xxx", nil)
	if err != nil {
		// TODO: エラーハンドリング
	}
	resp, err := client.Do(req)
	if err != nil {
		// TODO: エラーハンドリング
	}
	fmt.Println(resp)
}
