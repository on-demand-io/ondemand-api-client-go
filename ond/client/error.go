package client

import (
	"errors"
	"strconv"
)

type ErrResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
	Status    int
}

func (e ErrResponse) Error() error {
	r := e.ErrorCode
	if e.Status != 0 {
		statusString := strconv.Itoa(e.Status)
		r = statusString + " " + r
	}
	if len(e.Message) != 0 {
		r = r + " " + e.Message
	}

	return errors.New(r)
}
