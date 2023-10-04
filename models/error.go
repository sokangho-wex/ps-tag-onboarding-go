package models

import "errors"

var (
	BadRequestError   = errors.New("bad request")
	UserNotFoundError = errors.New("user not found")
	UnexpectedError   = errors.New("unexpected error")
)
