package domain

import (
	"errors"
	"fmt"
)

var (
	ErrUserWithUsernameAlreadyExists = errors.New("user with username already exists")
	ErrUserWithEmailAlreadyExists    = errors.New("user with email already exists")
	ErrNoResult                      = errors.New("no result found")
)

type ErrNotLongEnough struct {
	field     string
	minLength int
}

func (e ErrNotLongEnough) Error() string {
	return fmt.Sprintf("%v length must be at least %v", e.field, e.minLength)
}

type ErrIsRequired struct {
	field string
}

func (e ErrIsRequired) Error() string {
	return fmt.Sprintf("%v is required", e.field)
}

type ErrShouldMatch struct {
	field1 string
	field2 string
}

func (e ErrShouldMatch) Error() string {
	return fmt.Sprintf("%v must match %v", e.field2, e.field1)
}

type ErrInvalidEmail struct {
	field string
}

func (e ErrInvalidEmail) Error() string {
	return fmt.Sprintf("%v must be a valid email address", e.field)
}
