package server

import (
	"fmt"
	"os"
	"os/signal"

	mongostorage "github.com/lucasalmeron/backtabgo/pkg/storage/mongo"
	"github.com/lucasalmeron/microc3/users/internal/handler"
	"github.com/lucasalmeron/microc3/users/internal/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	users "github.com/lucasalmeron/microc3/users/internal/proto"
)

var (
	mongoURI      = os.Getenv("MONGODB_URI")
	mongoDataBase = os.Getenv("MONGODB_DB")
	MicroServer   = new(gRPCServer)
)

type gRPCServer struct {
	Addr         string
	MicroService micro.Service
}

func (srv *gRPCServer) Init() {

	srv.Addr = ":" + os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		srv.Addr = "127.0.0.1:3500"
	}

	// New Service
	srv.MicroService = micro.NewService(
		micro.Name("go.micro.service.users"),
		micro.Version("latest"),
	)

	// Initialise service
	srv.MicroService.Init()

	// Register Handler
	users.RegisterUsersHandler(srv.MicroService.Server(), new(handler.Users))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.auth.test", srv.MicroService.Server(), new(subscriber.Users))

	handler.NewPublisher(srv.MicroService.Client())

}

func (srv *gRPCServer) ConnectMongoDB() error {
	return mongostorage.NewMongoDBConnection(mongoURI, mongoDataBase)
}

func (s *gRPCServer) StartAndListen() {
	go s.waitForShutdown()

	// Run service
	if err := s.MicroService.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s *gRPCServer) waitForShutdown() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	fmt.Println("\nShutdown started...")
	if err := s.MicroService.Server().Stop(); err != nil {
		log.Logf(log.InfoLevel, log.String(), "Shutdown Error: %v", err)
	}
}
