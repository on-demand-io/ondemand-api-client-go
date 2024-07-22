package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dinson/ond-api-client-go/ond/errors"
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

func (i impl) ListMessages(ctx context.Context) (*ListMessagesResponse, *errors.ErrResponse) {
	//TODO implement me
	panic("implement me")
}
