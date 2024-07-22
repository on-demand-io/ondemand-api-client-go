package chat

import (
	"context"
	"github.com/dinson/ond-api-client-go/ond/client"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"github.com/dinson/ond-api-client-go/ond/params"
)

type Session interface {
	CreateSession(ctx context.Context, req *params.CreateChatSessionParams) (*CreateSessionResponse, *errors.ErrResponse)
	ListSessions(ctx context.Context, req *params.ListSessionsParams) (*ListSessionsResponse, *errors.ErrResponse)
}

type Chat interface {
	Session

	Query(ctx context.Context, req *params.QueryParams) (*SubmitQueryResponse, *errors.ErrResponse)
	GetSession(ctx context.Context, sessionID string) (*GetSessionResponse, *errors.ErrResponse)
	GetMessage(ctx context.Context, sessionID, messageID string) (*GetMessageResponse, *errors.ErrResponse)
	ListMessages(ctx context.Context) (*ListMessagesResponse, *errors.ErrResponse)
}

type impl struct {
	opts   *client.Options
	client client.Client
}

const (
	resourceURL = "/chat/v1/sessions%s"
)

func New(opts *client.Options) Chat {
	// session must be created here
	return &impl{
		opts:   opts,
		client: client.New(),
	}
}

func (i impl) Query(ctx context.Context, req *params.QueryParams) (*SubmitQueryResponse, *errors.ErrResponse) {
	//TODO implement me
	panic("implement me")
}

func (i impl) GetSession(ctx context.Context, sessionID string) (*GetSessionResponse, *errors.ErrResponse) {
	//TODO implement me
	panic("implement me")
}

func (i impl) GetMessage(ctx context.Context, sessionID, messageID string) (*GetMessageResponse, *errors.ErrResponse) {
	//TODO implement me
	panic("implement me")
}

func (i impl) ListMessages(ctx context.Context) (*ListMessagesResponse, *errors.ErrResponse) {
	//TODO implement me
	panic("implement me")
}
