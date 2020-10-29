package main

import "github.com/lucasalmeron/microc3/auth/internal/server"

func main() {
	srv := new(server.GRPCServer)

	srv.ConnectMongoDB()
	srv.Init()
	srv.LoadAuthRoutes()
	srv.StartAndListen()
}
