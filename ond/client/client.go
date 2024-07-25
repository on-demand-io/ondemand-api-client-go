package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
	"runtime"
	"time"
)

type Client interface {
	Do(ctx context.Context, opts *Options, method string, path string, payload []byte) (*http.Response, *errors.ErrResponse)
	Subscribe(ctx context.Context, opts *Options, method string, path string, payload []byte) (*http.Response, *errors.ErrResponse)
}

type impl struct{}

func New() Client {
	return &impl{}
}

const (
	Charset        string        = "UTF-8"
	DefaultTimeout time.Duration = 10 * time.Second
	baseURL        string        = "https://api.on-demand.io"
)

func (i impl) Do(ctx context.Context, opts *Options, method string, path string, payload []byte) (*http.Response, *errors.ErrResponse) {
	url := baseURL + path

	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: "",
		}
	}

	addHeaders(httpReq, opts.AuthKey)
	httpReq.Header.Add("Content-Type", "application/json")

	retryableClient := retryablehttp.NewClient()
	retryableClient.RetryMax = opts.Retries

	httpClient := retryableClient.StandardClient() // convert retryable http client to standard http client
	httpClient.Timeout = opts.HTTPTimeout

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
		}
	}

	if resp.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, &errors.ErrResponse{
				Message:   err.Error(),
				ErrorCode: errors.ErrAPIClientError.String(),
				Status:    resp.StatusCode,
			}
		}
		var errResp errors.ErrResponse
		if err = json.Unmarshal(body, &errResp); err != nil {
			return nil, &errors.ErrResponse{
				Message:   err.Error(),
				ErrorCode: errors.ErrAPIClientError.String(),
				Status:    resp.StatusCode,
			}
		}

		errResp.Status = resp.StatusCode

		return nil, &errResp
	}

	return resp, nil
}

func addHeaders(httpReq *http.Request, secret string) {
	httpReq.Header.Add("Accept-Charset", Charset)
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("User-Agent", "OnDemand-Go-Client")
	httpReq.Header.Add("Lang-Version", runtime.Version())
	httpReq.Header.Add("OS-Version", runtime.GOOS+" "+runtime.GOARCH)
	httpReq.Header.Add("apikey", secret)
}
