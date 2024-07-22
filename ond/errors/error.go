package errors

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

type ErrorCode string

var (
	ErrAPIClientError ErrorCode = "api_client_error"
)

func (e ErrorCode) String() string {
	return string(e)
}
