package services

import (
	"context"

	"github.com/BiteBit/grpc-gateway-example/api"
)

type (
	EchoServices struct{}
)

func (ctrl *EchoServices) Echo(ctx context.Context, msg *api.StringMessage) (*api.StringMessage, error) {
	return &api.StringMessage{
		Value: "hi " + msg.Value,
	}, nil
}
