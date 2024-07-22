package chat

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

func (i impl) GetMessage(ctx context.Context, sessionID, messageID string) (*GetMessageResponse, *errors.ErrResponse) {
	endpoint := fmt.Sprintf(resourceURL, "/"+sessionID+"/messages/"+messageID)

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

	var result GetMessageResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
			Status:    0,
		}
	}

	return &result, nil
}

func (i impl) ListMessages(ctx context.Context, req *params.ListMessageParams) (*ListMessagesResponse, *errors.ErrResponse) {
	endpoint := fmt.Sprintf(resourceURL, "/"+req.SessionID+"/messages")

	queryString, err := util.BuildQueryParamsString(req)
	if err != nil {
		fmt.Println(err)
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

	var result ListMessagesResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
			Status:    0,
		}
	}

	return &result, nil
}
