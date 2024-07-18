package client

import (
	"bytes"
	"context"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
	"runtime"
	"time"
)

const (
	Charset        string        = "UTF-8"
	DefaultTimeout time.Duration = 5 * time.Second
	baseURL        string        = "https://api.on-demand.io"
)

func Do(ctx context.Context, opts *Options, method string, path string, payload []byte) (*http.Response, error) {
	url := baseURL + path

	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	addHeaders(httpReq, opts.AuthKey)

	retryableClient := retryablehttp.NewClient()
	retryableClient.RetryMax = opts.Retries

	httpClient := retryableClient.StandardClient() // convert retryable http client to standard http client
	httpClient.Timeout = opts.HTTPTimeout

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode >= http.StatusBadRequest {
		// TODO: return appropriate error message
	}

	return resp, nil
}

func addHeaders(httpReq *http.Request, secret string) {
	httpReq.Header.Add("Accept-Charset", Charset)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("User-Agent", "OnDemand-Go-Client")
	httpReq.Header.Add("Authorization", "Bearer "+secret)
	httpReq.Header.Add("Lang-Version", runtime.Version())
	httpReq.Header.Add("OS-Version", runtime.GOOS+" "+runtime.GOARCH)
}
