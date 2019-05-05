// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"errors"
	"strings"
)

// BadRequestError holds list of errors from http response that represent
// invalid data submitted by the user.
type BadRequestError struct {
	errors []string
}

// NewBadRequestError constructs a new BadRequestError with provided errors.
func NewBadRequestError(errors ...string) (err *BadRequestError) {
	return &BadRequestError{
		errors: errors,
	}
}

func (e *BadRequestError) Error() (s string) {
	return strings.Join(e.errors, " ")
}

// Errors returns a list of error messages.
func (e *BadRequestError) Errors() (errs []string) {
	return e.errors
}

// Errors that are returned by the API.
var (
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrNotFound            = errors.New("not found")
	ErrMethodNotAllowed    = errors.New("method not allowed")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrInternalServerError = errors.New("internal server error")
	ErrMaintenance         = errors.New("maintenance")
)

var errInvalidPageNumber = errors.New("invalid page number")
