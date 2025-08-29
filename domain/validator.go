package domain

import (
	"regexp"
)

type Validator struct {
	errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{
		errors: make(map[string]string),
	}
}

func (v *Validator) MustNotBeEmpty(field, value string) bool {
	// if error already exists for this field, skip further validation
	if _, ok := v.errors[field]; ok {
		return false
	}

	// skip validation if value is empty
	if value == "" || len(value) == 0 {
		v.errors[field] = ErrIsRequired{field: field}.Error()

		return false
	}

	return true
}

func (v *Validator) MustBeLongerThan(field, value string, constraint int) bool {
	// if error already exists for this field, skip further validation
	if _, ok := v.errors[field]; ok {
		return false
	}
	// skip validation if value is empty
	if value == "" {
		return true
	}

	if len(value) < constraint {
		v.errors[field] = ErrNotLongEnough{field: field, minLength: constraint}.Error()
		return false
	}

	return true
}

// Can also use custom struct instead of 4 parameters
//
//	type ElementMatcher struct {
//		field string,
//		val string
//	}
//
// func (v *Validator) MustMatch(o, c ElementMatcher) bool {
func (v *Validator) MustMatch(field1, val1, field2, val2 string) bool {
	if _, ok := v.errors[field1]; ok {
		return false
	}
	if val1 != val2 {
		v.errors[field2] = ErrShouldMatch{field1: field1, field2: field2}.Error()
		return false
	}
	return true
}

func (v *Validator) MustBeValidEmail(field, email string) bool {
	var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if _, ok := v.errors[field]; ok {
		return false
	}

	if len(email) < 3 || len(email) > 254 || !emailRegexp.MatchString(email) {
		v.errors[field] = ErrInvalidEmail{field: field}.Error()
		return false
	}

	return true
}

func (v *Validator) HasErrors() bool {
	return len(v.errors) == 0
}
