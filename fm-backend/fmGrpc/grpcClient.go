package fmGrpc

import (
	"context"
	"github.com/FusionMate/fm-backend/common"
	"github.com/FusionMate/fm-backend/common/log"
	"github.com/FusionMate/fm-backend/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func GetGrpcConn() (*grpc.ClientConn, error) {
	var retryPolicy = `{
            "methodConfig": [{
                "name": [{"service": "grpc.examples.echo.Echo"}],
                "waitForReady": true,
                "retryPolicy": {
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]
        }`

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(retryPolicy),
	}
	conn, err := grpc.Dial(conf.GConfig.GetString("server.fmBotGrpc"), opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PoeBotCreation(collectionId int64, makerAddress string, prompt string) error {
	conn, err := GetGrpcConn()
	if err != nil {
		log.Error("[GRPC - Client - PoeBotCreation] init grpc connection err: %s", err)
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Error("[GRPC - Client - PoeBotCreation] close grpc connection error: %v", err)
		}
	}(conn)
	c := NewBotServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	reply, err := c.BotCreate(ctx, &BotCreateRequest{
		CollectionId: collectionId,
		MakerAddress: makerAddress,
		Prompt:       prompt,
	})
	if err != nil {
		log.Error("[GRPC - Client - PoeBotCreation] create bot from grpc server error: %s", err)
		return err
	}
	if reply.Status == 0 {
		err = common.ErrDuplicatedBotCreation
		log.Error("[GRPC - Client - PoeBotCreation] create bot error: %s", err)
		return err
	} else if reply.Status == 2 {
		err = common.ErrPoeRejectBotCreation
		log.Error("[GRPC - Client - PoeBotCreation] create bot error: %s", err)
		return err
	} else if reply.Status == 3 {
		err = common.ErrBotLimitReaches
		log.Error("[GRPC - Client - PoeBotCreation] create bot error: %s", err)
		return err
	}
	log.Info("[GRPC - Client - PoeBotCreation] create bot by collection id %d successfully", collectionId)
	return nil
}
