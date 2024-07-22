package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"github.com/dinson/ond-api-client-go/ond/util"
	"io"
	"net/http"
)

func (i impl) List(ctx context.Context, req *ListRequest) (*ListResponse, *errors.ErrResponse) {
	endpoint := fmt.Sprintf(resourceURL, "list")

	if len(req.PluginIDs) != 0 {
		req.PluginIDs = convertToCommaSeperatedString(req.PluginIDs)
	}

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
			Status:    0,
		}
	}

	var result ListResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
			Status:    0,
		}
	}

	return &result, nil
}

func convertToCommaSeperatedString(ds []string) []string {
	s := ""

	for i, v := range ds {
		if i > 0 {
			s = s + ","
		}
		s = s + v
	}

	return []string{s}
}
