package rpc

import (
	"context"

	"github.com/crafted-systems/smartcollect-pro/user/proto"
)

// user is a struct which contains methods for the gRPC handlers.
type user struct {
}

func (s *user) AMethod(_ context.Context, ok *proto.OK) (*proto.OK, error) {
	return &proto.OK{Ok: ok.GetOk()}, nil
}

func New() (user, error) {
	return user{}, nil
}
