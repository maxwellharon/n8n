package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCError is just an error that can be unwrapped and contains a GRPC code.
type GRPCError struct {
	Err  error
	Code codes.Code
}

func (m GRPCError) Unwrap() error {
	return m.Err
}

func (m GRPCError) Error() string {
	return m.Err.Error()
}

func (m GRPCError) GRPCStatus() *status.Status {
	if m.Code == codes.OK && m.Err != nil {
		return status.New(codes.Unknown, m.Error())
	}

	return status.New(m.Code, m.Error())
}

func New(err error, code codes.Code) *GRPCError {
	return &GRPCError{
		Err:  err,
		Code: code,
	}
}
