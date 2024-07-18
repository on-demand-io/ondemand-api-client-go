// Package ond - official golang client package for OnDemand API.
package ond

import (
	"airevai/ondemand-api/ond/client"
	"airevai/ondemand-api/ond/services/chat"
	"airevai/ondemand-api/ond/services/media"
	"airevai/ondemand-api/ond/services/plugin"
	"time"
)

type Ond interface{}

type impl struct {
	Chat   chat.Chat
	Media  media.Media
	Plugin plugin.Plugin
}

type Options struct {
	HTTPTimeout time.Duration
	Retries     int
}

// Init returns a new OnDemand API client
func Init(secretKey string, opts ...*Options) Ond {
	clientOpts := defaultClientOpts(secretKey)

	if len(opts) != 0 {
		clientOpts.HTTPTimeout = opts[0].HTTPTimeout
		clientOpts.Retries = opts[0].Retries
	}

	return &impl{
		Chat:   chat.New(),
		Media:  media.New(),
		Plugin: plugin.New(clientOpts),
	}
}

func defaultClientOpts(secret string) *client.Options {
	return &client.Options{
		AuthKey:     secret,
		HTTPTimeout: client.DefaultTimeout,
		Retries:     0,
	}
}
