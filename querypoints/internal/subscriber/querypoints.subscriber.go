package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	protousers "github.com/lucasalmeron/microc3/users/pkg/users/proto"
)

type QueryPoints struct{}

func (e *QueryPoints) Handle(ctx context.Context, msg *protousers.ResponseUser) error {
	log.Info("Handler Received message: ", msg)
	return nil
}
