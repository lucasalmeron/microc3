package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	users "github.com/lucasalmeron/microc3/users/pkg/users/proto"
)

type Users struct{}

func (e *Users) Handle(ctx context.Context, msg *users.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *users.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
