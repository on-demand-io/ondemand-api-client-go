package speech

import (
	"context"
	"github.com/dinson/ond-api-client-go/ond/client"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"github.com/dinson/ond-api-client-go/ond/params"
)

type Speech interface {
	ToText(ctx context.Context, req *params.SpeechToTextParams) (*ToTextResponse, *errors.ErrResponse)
	FromText(ctx context.Context, req *params.TextToSpeechParams) (*ToSpeechResponse, *errors.ErrResponse)
}

type impl struct {
	opts   *client.Options
	client client.Client
}

const (
	resourceURL = "/services/v1/public/service/execute%s"
)

func New(opts *client.Options) Speech {
	return &impl{
		opts:   opts,
		client: client.New(),
	}
}
