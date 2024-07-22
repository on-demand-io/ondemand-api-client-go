package plugin

import (
	"context"
	"github.com/dinson/ond-api-client-go/ond/client"
	"github.com/dinson/ond-api-client-go/ond/errors"
)

type Plugin interface {
	List(ctx context.Context, req *ListRequest) (*ListResponse, *errors.ErrResponse)
}

type impl struct {
	opts   *client.Options
	client client.Client
}

const (
	resourceURL = "/plugin/v1/%s"
)

func New(opts *client.Options) Plugin {
	return &impl{
		opts:   opts,
		client: client.New(),
	}
}
