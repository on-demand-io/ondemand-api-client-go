package plugin

import (
	"airevai/ondemand-api/ond/client"
	"airevai/ondemand-api/ond/util"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (i impl) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	endpoint := fmt.Sprintf(resourceURL, "list")

	queryString, err := util.BuildQuery(req)
	if err != nil {
		return nil, err
	}

	if len(queryString) != 0 {
		endpoint = fmt.Sprintf("%s?%s", endpoint, queryString)
	}

	resp, err := client.Do(ctx, i.Opts, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ListResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
