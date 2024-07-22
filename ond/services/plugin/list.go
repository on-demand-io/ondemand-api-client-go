package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dinson/ond-api-client-go/ond/client"
	"github.com/dinson/ond-api-client-go/ond/util"
	"io"
	"net/http"
)

func (i impl) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	endpoint := fmt.Sprintf(resourceURL, "list")

	if len(req.PluginIDs) != 0 {
		req.PluginIDs = convertToCommaSeperatedString(req.PluginIDs)
	}

	queryString, err := util.BuildQuery(req)
	if err != nil {
		return nil, err
	}

	if len(queryString) != 0 {
		endpoint = fmt.Sprintf("%s?%s", endpoint, queryString)
	}

	resp, respErr := client.Do(ctx, i.Opts, http.MethodGet, endpoint, nil)
	if respErr != nil {
		return nil, respErr.Error()
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
