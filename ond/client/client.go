package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

const (
	Charset        string        = "UTF-8"
	DefaultTimeout time.Duration = 5 * time.Second
	baseURL        string        = "https://api.on-demand.io"
)

type ErrResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
	Status    int
}

func (e ErrResponse) Error() error {
	r := e.ErrorCode
	if e.Status != 0 {
		statusString := strconv.Itoa(e.Status)
		r = statusString + " " + r
	}
	if len(e.Message) != 0 {
		r = r + " " + e.Message
	}

	return errors.New(r)
}

func Do(ctx context.Context, opts *Options, method string, path string, payload []byte) (*http.Response, *ErrResponse) {
	url := baseURL + path

	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, &ErrResponse{
			Message:   err.Error(),
			ErrorCode: "",
		}
	}

	addHeaders(httpReq, opts.AuthKey)

	retryableClient := retryablehttp.NewClient()
	retryableClient.RetryMax = opts.Retries

	httpClient := retryableClient.StandardClient() // convert retryable http client to standard http client
	httpClient.Timeout = opts.HTTPTimeout

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, &ErrResponse{
			Message:   err.Error(),
			ErrorCode: "",
		}
	}

	if resp.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, &ErrResponse{
				Message:   err.Error(),
				ErrorCode: "",
				Status:    resp.StatusCode,
			}
		}
		var errResp ErrResponse
		if err = json.Unmarshal(body, &errResp); err != nil {
			return nil, &ErrResponse{
				Message:   err.Error(),
				ErrorCode: "",
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
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("User-Agent", "OnDemand-Go-Client")
	httpReq.Header.Add("Lang-Version", runtime.Version())
	httpReq.Header.Add("OS-Version", runtime.GOOS+" "+runtime.GOARCH)
	httpReq.Header.Add("apikey", secret)
}
