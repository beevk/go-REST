package domain

import "errors"

var (
	ErrUserWithUsernameAlreadyExists = errors.New("user with username already exists")
	ErrUserWithEmailAlreadyExists    = errors.New("user with email already exists")
	ErrNoResult                      = errors.New("no result found")
)
