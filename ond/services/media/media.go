package media

import (
	"context"
	"github.com/dinson/ond-api-client-go/ond/client"
	"github.com/dinson/ond-api-client-go/ond/errors"
	"github.com/dinson/ond-api-client-go/ond/params"
)

type Media interface {
	Create(ctx context.Context, req *params.CreateMediaParams) (*CreateMediaResponse, *errors.ErrResponse)
	Fetch(ctx context.Context, req *params.FetchMediaParams) (*FetchMediaResponse, *errors.ErrResponse)
	Delete(ctx context.Context, fileID string) *errors.ErrResponse
}

type impl struct {
	opts   *client.Options
	client client.Client
}

const (
	resourceURL = "/media/v1/public/file%s"
)

func New(opts *client.Options) Media {
	return &impl{
		opts:   opts,
		client: client.New(),
	}
}
