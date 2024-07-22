package speech

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"github.com/dinson/ond-api-client-go/ond/params"
	"io"
	"net/http"
)

func (i impl) ToText(ctx context.Context, req *params.SpeechToTextParams) (*ToTextResponse, *errors.ErrResponse) {
	endpoint := fmt.Sprintf(resourceURL, "/speech_to_text")

	payloadBytes, err := json.Marshal(req)
	if err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
		}
	}

	resp, respErr := i.client.Do(ctx, i.opts, http.MethodPost, endpoint, payloadBytes)
	if respErr != nil {
		return nil, respErr
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
			Status:    0,
		}
	}

	var result ToTextResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
			Status:    0,
		}
	}

	return &result, nil
}
