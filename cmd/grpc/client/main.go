package main

import (
	"context"
	"fmt"

	"github.com/boxgo/grpc/api"
	"github.com/boxgo/grpc/config"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(config.GRPC_SERVER_ADDR, grpc.WithInsecure())
	defer conn.Close()

	client := api.NewEchoServiceClient(conn)
	resp, _ := client.Echo(context.Background(), &api.StringMessage{Value: "bitebit"})
	fmt.Println(resp)
}
