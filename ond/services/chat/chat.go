package chat

import (
	"context"
	"github.com/dinson/ond-api-client-go/ond/client"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"github.com/dinson/ond-api-client-go/ond/params"
)

type Session interface {
	CreateSession(ctx context.Context, req *params.CreateChatSessionParams) (*CreateSessionResponse, *errors.ErrResponse)
	ListSessions(ctx context.Context, req *params.ListSessionParams) (*ListSessionsResponse, *errors.ErrResponse)
	GetSession(ctx context.Context, sessionID string) (*GetSessionResponse, *errors.ErrResponse)
}

type Chat interface {
	Session

	// Query can be used to get response for a question via sync or webhook.
	Query(ctx context.Context, req *params.QueryParams) (*SubmitQueryResponse, *errors.ErrResponse)
	// OpenStream initiates an SSE connection with OnDemand servers.
	// The response of this method should be passed to "Consume()" to start receiving the response.
	OpenStream(ctx context.Context, req *params.QueryParams) (*StreamConsumer, *errors.ErrResponse)
	GetMessage(ctx context.Context, sessionID, messageID string) (*GetMessageResponse, *errors.ErrResponse)
	ListMessages(ctx context.Context, req *params.ListMessageParams) (*ListMessagesResponse, *errors.ErrResponse)
}

type impl struct {
	opts   *client.Options
	client client.Client
}

const (
	resourceURL = "/chat/v1/sessions%s"
)

func New(opts *client.Options) Chat {
	return &impl{
		opts:   opts,
		client: client.New(),
	}
}
