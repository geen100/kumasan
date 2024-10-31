package models

import (
	"errors"
)

var (
	ErrAccountNotFound      = errors.New("account is not found")
	ErrAccountAlreadyExists = errors.New("account ID already exists")
	ErrInternalServer       = errors.New("internal server error")
)
