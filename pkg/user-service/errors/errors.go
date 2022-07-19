package errors

import (
	"errors"
)

var (
	ErrEmptyUsername      = errors.New("empty username")
	ErrEmptyEmail         = errors.New("empty email")
	ErrBadUsernameLength  = errors.New("bad username length")
	ErrBadEmailLength     = errors.New("bad email length")
	ErrDifferentPasswords = errors.New("different passwords")
)
