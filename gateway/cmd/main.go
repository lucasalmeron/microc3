package main

import (
	server "github.com/lucasalmeron/microc3/gateway/internal/server"
)

func main() {
	server := new(server.Server)

	server.Init()
}
