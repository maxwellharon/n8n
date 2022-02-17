package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGRPCErrorString(t *testing.T) {
	err := GRPCError{
		Err:  assert.AnError,
		Code: codes.Unknown,
	}

	assert.Error(t, err)
	assert.Contains(t, err.Error(), assert.AnError.Error())
}

func TestGRPCErrorGRPC(t *testing.T) {
	err := GRPCError{
		Err:  assert.AnError,
		Code: codes.NotFound,
	}

	assert.Error(t, err)
	assert.Contains(t, err.Error(), assert.AnError.Error())
	assert.Equal(t, codes.NotFound, status.Convert(err).Code())
}

func TestGRPCErrorGRPCUnknown(t *testing.T) {
	err := GRPCError{
		Err: assert.AnError,
	}

	assert.Error(t, err)
	assert.Contains(t, err.Error(), assert.AnError.Error())
	assert.Equal(t, status.Convert(err).Code(), codes.Unknown)
}
