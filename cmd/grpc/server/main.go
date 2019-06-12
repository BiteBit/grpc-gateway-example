package main

import (
	"fmt"
	"net"

	"github.com/boxgo/grpc/api"
	"github.com/boxgo/grpc/config"
	"github.com/boxgo/grpc/services"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	srv := services.EchoServices{}
	api.RegisterEchoServiceServer(server, &srv)

	fmt.Printf("grpc listen %s\n", config.GRPC_SERVER_ADDR)
	lis, err := net.Listen("tcp", config.GRPC_SERVER_ADDR)
	if err != nil {
		panic(err)
	}

	server.Serve(lis)
}
