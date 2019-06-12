package main

import (
	"context"
	"fmt"

	"github.com/BiteBit/grpc-gateway-example/api"
	"github.com/BiteBit/grpc-gateway-example/config"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(config.GRPC_SERVER_ADDR, grpc.WithInsecure())
	defer conn.Close()

	client := api.NewEchoServiceClient(conn)
	resp, _ := client.Echo(context.Background(), &api.StringMessage{Value: "bitebit"})
	fmt.Println(resp)
}
