package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
)

// Subscribe connects to the SSE endpoint and sends events through the channel
func (i impl) Subscribe(ctx context.Context, opts *Options, method string, path string, payload []byte) (*http.Response, *errors.ErrResponse) {
	url := baseURL + path

	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: "",
		}
	}

	addHeaders(httpReq, opts.AuthKey)
	httpReq.Header.Add("Content-Type", "text/event-stream")

	retryableClient := retryablehttp.NewClient()
	retryableClient.RetryMax = opts.Retries
	retryableClient.Logger = nil

	httpClient := retryableClient.StandardClient() // convert retryable http client to standard http client
	httpClient.Timeout = opts.HTTPTimeout

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: "",
		}
	}

	if resp.StatusCode >= http.StatusBadRequest {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, &errors.ErrResponse{
				Message:   err.Error(),
				ErrorCode: "",
				Status:    resp.StatusCode,
			}
		}
		var errResp errors.ErrResponse
		if err = json.Unmarshal(body, &errResp); err != nil {
			return nil, &errors.ErrResponse{
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
