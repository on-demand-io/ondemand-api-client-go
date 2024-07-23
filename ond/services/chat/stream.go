package chat

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"github.com/dinson/ond-api-client-go/ond/params"
	"io"
	"net/http"
	"strings"
)

type OpenStreamResponse struct {
	HTTPResp *http.Response
}

// Event represents a single SSE event
type Event struct {
	Data  EventData
	Error error
	Done  bool
}

//type Event string

type EventData struct {
	SessionID  string
	MessageID  string
	Answer     string
	Status     string
	EventIndex int
	EventType  string
}

func (i impl) OpenStream(ctx context.Context, req *params.QueryParams) (*OpenStreamResponse, *errors.ErrResponse) {

	req.ResponseMode = params.ResponseModeStream

	endpoint := fmt.Sprintf(resourceURL, "/"+req.SessionID+"/query")

	payloadBytes, err := json.Marshal(req)
	if err != nil {
		return nil, &errors.ErrResponse{
			Message:   err.Error(),
			ErrorCode: errors.ErrAPIClientError.String(),
		}
	}

	resp, respErr := i.client.Subscribe(ctx, i.opts, http.MethodPost, endpoint, payloadBytes)
	if respErr != nil {
		return nil, respErr
	}

	return &OpenStreamResponse{
		HTTPResp: resp,
	}, nil
}

func (i impl) ConsumeStream(resp *OpenStreamResponse, events chan<- Event) {
	defer close(events) // Close the channel when the connection ends

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.HTTPResp.Body)

	scanner := bufio.NewScanner(resp.HTTPResp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "event:message" || line == "" {
			continue
		}

		var eventData EventData

		if line == "data:[DONE]" {
			events <- Event{
				Data:  eventData,
				Error: nil,
				Done:  true,
			}
			break
		}

		parsedString := strings.Split(line, "data:")
		jsonString := parsedString[1]

		if err := json.Unmarshal([]byte(jsonString), &eventData); err != nil {
			events <- Event{Error: err}
			break
		}

		events <- Event{
			Data:  eventData,
			Error: nil,
			Done:  false,
		}
	}

	if err := scanner.Err(); err != nil {
		events <- Event{Error: err}
	}
}
