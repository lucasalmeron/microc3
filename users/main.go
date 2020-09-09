package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"users/handler"
	"users/subscriber"

	users "users/proto/users"
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
