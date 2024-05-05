package model

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrNotValid = errors.New("not valid")
)
