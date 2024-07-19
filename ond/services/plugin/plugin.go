package plugin

import (
	"context"
	"github.com/dinson/ond-api-client-go/ond/client"
)

type Plugin interface {
	List(ctx context.Context, req *ListRequest) (*ListResponse, error)
}

type impl struct {
	Opts *client.Options
}

const (
	resourceURL = "/plugin/v1/%s"
)

func New(opts *client.Options) Plugin {
	return &impl{
		Opts: opts,
	}
}
