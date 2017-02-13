package util

import "errors"

type ErrorWithStatus struct {
	error
	code int
}

func (error ErrorWithStatus) StatusCode() int { return error.code }

var ErrBadRoute = ErrorWithStatus{errors.New("Bad route"), 404}
var ErrEmpty = ErrorWithStatus{errors.New("Empty body"), 406}
var ErrWrongMethod = ErrorWithStatus{errors.New("Request has wrong method"), 405}
var ErrUnsupportedMediaType = ErrorWithStatus{errors.New(""), 415}
