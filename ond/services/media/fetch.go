package media

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"github.com/dinson/ond-api-client-go/ond/params"
	"github.com/dinson/ond-api-client-go/ond/util"
	"io"
	"net/http"
)

func (i impl) Fetch(ctx context.Context, req *params.FetchMediaParams) (*FetchMediaResponse, *errors.ErrResponse) {
	endpoint := fmt.Sprintf(resourceURL, "")

	queryString, err := util.BuildQueryParamsString(req)
	if err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
			Status:    0,
		}
	}

	if len(queryString) != 0 {
		endpoint = fmt.Sprintf("%s?%s", endpoint, queryString)
	}

	resp, respErr := i.client.Do(ctx, i.opts, http.MethodGet, endpoint, nil)
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

	var result FetchMediaResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
			Status:    0,
		}
	}

	return &result, nil
}
