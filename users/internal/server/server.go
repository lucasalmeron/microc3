package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/lucasalmeron/microc3/users/internal/handler"
	mongostorage "github.com/lucasalmeron/microc3/users/internal/storage"
	"github.com/lucasalmeron/microc3/users/internal/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	users "github.com/lucasalmeron/microc3/users/internal/proto"
)

var (
	mongoURI      = os.Getenv("MONGODB_URI")
	mongoDataBase = os.Getenv("MONGODB_DB")
)

type GRPCServer struct {
	MicroService micro.Service
}

func (srv *GRPCServer) Init() {

	// New Service
	srv.MicroService = micro.NewService(
		micro.Name("go.micro.service.users"),
		micro.Version("latest"),
	)

	// Initialise service
	srv.MicroService.Init()

	srv.registerHandlers()
	srv.registerEventSubscribers()

}

func (srv *GRPCServer) registerHandlers() error {
	log.Info("Registering Handlers")
	//INIT EVENTS
	handler.InitEvents(srv.MicroService.Client())
	// Register Handler
	err := users.RegisterUsersHandler(srv.MicroService.Server(), new(handler.UsersHandler))
	if err != nil {
		return err
	}

	return nil
}

func (srv *GRPCServer) registerEventSubscribers() error {
	log.Info("Registering Subscribers")
	// Register Struct as Subscriber
	err := micro.RegisterSubscriber("go.micro.auth.test", srv.MicroService.Server(), new(subscriber.Users))
	if err != nil {
		return err
	}
	return nil
}

func (srv *GRPCServer) ConnectMongoDB() error {
	return mongostorage.NewMongoDBConnection(mongoURI, mongoDataBase)
}

func (s *GRPCServer) StartAndListen() {
	go s.waitForShutdown()

	// Run service
	if err := s.MicroService.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s *GRPCServer) waitForShutdown() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	fmt.Println("\nShutdown started...")
	if err := s.MicroService.Server().Stop(); err != nil {
		log.Logf(log.InfoLevel, log.String(), "Shutdown Error: %v", err)
	}
}
