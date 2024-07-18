package client

import "time"

type Options struct {
	AuthKey     string
	HTTPTimeout time.Duration
	Retries     int
}
