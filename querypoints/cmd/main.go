package main

import "github.com/lucasalmeron/microc3/querypoints/internal/server"

func main() {
	srv := new(server.GRPCServer)

	srv.ConnectMongoDB()
	srv.Init()
	srv.StartAndListen()
}
