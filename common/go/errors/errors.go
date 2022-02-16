package errors

import "errors"

// ErrExpected is an error that is thrown when it is intended behaviour.
var ErrExpected = errors.New("this error is intended behaviour")
