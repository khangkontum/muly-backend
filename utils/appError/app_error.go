package appError

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrEditConflict   = errors.New("edit conflict")
)
