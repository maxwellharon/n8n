package rpc

import (
	"context"
	"testing"
	"time"

	user_proto "github.com/crafted-systems/smartcollect-pro/user/proto"
	mock_user_proto "github.com/crafted-systems/smartcollect-pro/user/proto/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserServer(t *testing.T) {
	mockedController := gomock.NewController(t)
	defer mockedController.Finish()
	mockedServer := mock_user_proto.NewMockUserServer(mockedController)
	rpcMessage := &user_proto.OK{Ok: true}

	mockedServer.EXPECT().AMethod(
		gomock.Any(),
		rpcMessage,
	).Return(rpcMessage, nil).Times(1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := mockedServer.AMethod(ctx, rpcMessage)
	assert.Equal(t, resp, rpcMessage)
	assert.NoError(t, err)
}
