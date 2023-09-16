package fmGrpc

import (
	"context"
	"errors"
	"github.com/FusionMate/fm-backend/service"
)

type BotServiceServerImpl struct {
	UnimplementedBotServiceServer
}

func (b *BotServiceServerImpl) TokenVerify(ctx context.Context, request *TokenVerifyRequest) (*TokenVerifyReply, error) {
	address, ok := service.TokenCheck(request.Token)
	if !ok {
		err := errors.New("token verify failed")
		return nil, err
	}
	return &TokenVerifyReply{
		UserAddress: address,
	}, nil
}
