package main

import "errors"

type ErrorWithStatus struct {
	error
	code int
}

func (error ErrorWithStatus) StatusCode() int { return error.code}


var ErrEmpty = ErrorWithStatus{ errors.New("Empty body"), 406}
var ErrWrongMethod = ErrorWithStatus{ errors.New("Request has wrong method") , 405}
