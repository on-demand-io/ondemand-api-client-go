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

type StreamConsumer struct {
	HTTPResp  *http.Response
	EventChan chan Event
}

// Event represents a single SSE event
type Event struct {
	Data  EventData
	Error error
	Done  bool
}

type EventData struct {
	SessionID  string
	MessageID  string
	Answer     string
	Status     string
	EventIndex int
	EventType  string
}

func (i impl) OpenStream(ctx context.Context, req *params.QueryParams) (*StreamConsumer, *errors.ErrResponse) {

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

	events := make(chan Event)

	return &StreamConsumer{
		HTTPResp:  resp,
		EventChan: events,
	}, nil
}

// Consume helps to receive question responses via SSE events.
// Invoke this method as a go routine to prevent blocking.
func (s StreamConsumer) Consume() {
	defer close(s.EventChan) // Close the channel when the connection ends

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(s.HTTPResp.Body)

	scanner := bufio.NewScanner(s.HTTPResp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "event:message" || line == "" {
			continue
		}

		var eventData EventData

		if line == "data:[DONE]" {
			s.EventChan <- Event{
				Data:  eventData,
				Error: nil,
				Done:  true,
			}
			break
		}

		parsedString := strings.Split(line, "data:")
		jsonString := parsedString[1]

		if err := json.Unmarshal([]byte(jsonString), &eventData); err != nil {
			s.EventChan <- Event{Error: err}
			break
		}

		s.EventChan <- Event{
			Data:  eventData,
			Error: nil,
			Done:  false,
		}
	}

	if err := scanner.Err(); err != nil {
		s.EventChan <- Event{Error: err}
	}
}
