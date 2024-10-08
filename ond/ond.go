// Package ond - official golang client package for OnDemand API.
package ond

import (
	"github.com/dinson/ond-api-client-go/ond/client"
	"github.com/dinson/ond-api-client-go/ond/services/chat"
	"github.com/dinson/ond-api-client-go/ond/services/media"
	"github.com/dinson/ond-api-client-go/ond/services/plugin"
	"github.com/dinson/ond-api-client-go/ond/services/speech"
	"time"
)

type Services struct {
	Chat   chat.Chat
	Media  media.Media
	Plugin plugin.Plugin
	Speech speech.Speech
}

type Options struct {
	HTTPTimeout time.Duration
	Retries     int
}

// Init returns a new OnDemand API client
func Init(secretKey string, opts ...*Options) *Services {
	clientOpts := defaultClientOpts(secretKey)

	if len(opts) != 0 {
		clientOpts.HTTPTimeout = opts[0].HTTPTimeout
		clientOpts.Retries = opts[0].Retries
	}

	return &Services{
		Chat:   chat.New(clientOpts),
		Media:  media.New(clientOpts),
		Plugin: plugin.New(clientOpts),
		Speech: speech.New(clientOpts),
	}
}

func defaultClientOpts(secret string) *client.Options {
	return &client.Options{
		AuthKey:     secret,
		HTTPTimeout: client.DefaultTimeout,
		Retries:     0,
	}
}
