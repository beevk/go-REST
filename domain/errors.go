package domain

import (
	"errors"
	"fmt"
)

var (
	ErrUserWithUsernameAlreadyExists = errors.New("user with username already exists")
	ErrUserWithEmailAlreadyExists    = errors.New("user with email already exists")
	ErrNoResult                      = errors.New("no result found")
	ErrUnauthorized                  = errors.New("unauthorized")
	ErrForbidden                     = errors.New("forbidden")
	ErrInternalServerError           = errors.New("internal server error")
	ErrBadRequest                    = errors.New("bad request")
	ErrInvalidCredentials            = errors.New("invalid credentials. Please check your email and password")
)

type ErrNotLongEnough struct {
	field     string
	minLength int
}

type ErrInvalidEmail struct {
	field string
}

type ErrShouldMatch struct {
	field1 string
	field2 string
}

type ErrIsRequired struct {
	field string
}

func (e ErrNotLongEnough) Error() string {
	return fmt.Sprintf("%v length must be at least %v", e.field, e.minLength)
}

func (e ErrIsRequired) Error() string {
	return fmt.Sprintf("%v is required", e.field)
}

func (e ErrShouldMatch) Error() string {
	return fmt.Sprintf("%v must match %v", e.field2, e.field1)
}

func (e ErrInvalidEmail) Error() string {
	return fmt.Sprintf("%v must be a valid email address", e.field)
}
