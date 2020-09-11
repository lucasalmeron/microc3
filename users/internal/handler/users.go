package handler

import (
	"context"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	users "github.com/lucasalmeron/microc3/users/internal/proto"
)

var (
	pub micro.Publisher
)

func NewPublisher(c client.Client) {
	pub = micro.NewEvent("go.micro.users.created", c)
}

type Users struct{}

func (e *Users) CreateUser(ctx context.Context, req *users.RequestUser, res *users.ResponseUser) error {
	log.Info("Received Users.CreateUser request")
	res.FirstName = "TOTO"
	res.LastName = req.LastName

	return status.Error(codes.Internal, "testing error")
}
