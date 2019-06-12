package main

import (
	"fmt"
	"net"

	"github.com/BiteBit/grpc-gateway-example/api"
	"github.com/BiteBit/grpc-gateway-example/config"
	"github.com/BiteBit/grpc-gateway-example/services"
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
