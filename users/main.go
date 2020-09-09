package main

import (
	"github.com/lucasalmeron/microc3/users/handler"
	"github.com/lucasalmeron/microc3/users/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	users "github.com/lucasalmeron/microc3/users/proto/users"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.users"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	users.RegisterUsersHandler(service.Server(), new(handler.Users))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.users", service.Server(), new(subscriber.Users))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
