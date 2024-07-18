package plugin

import (
	"airevai/ondemand-api/ond/client"
	"context"
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
