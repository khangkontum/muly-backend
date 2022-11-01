package appError

import "errors"

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)
