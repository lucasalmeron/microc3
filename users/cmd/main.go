package main

import "github.com/lucasalmeron/microc3/users/internal/server"

func main() {
	server.MicroServer.Init()
	server.MicroServer.ConnectMongoDB()
	server.MicroServer.StartAndListen()
}
