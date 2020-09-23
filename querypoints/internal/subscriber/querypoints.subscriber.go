package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	querypoints "github.com/lucasalmeron/microc3/querypoints/internal/proto"
)

type QueryPoints struct{}

func (e *QueryPoints) Handle(ctx context.Context, msg *querypoints.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *querypoints.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
