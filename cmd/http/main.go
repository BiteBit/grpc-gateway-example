package main

import (
	"fmt"
	"net/http"

	"github.com/boxgo/grpc/api"
	"github.com/boxgo/grpc/config"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, config.GRPC_SERVER_ADDR, opts)
	if err != nil {
		return err
	}

	fmt.Printf("http listen %s\n", config.HTTP_SERVER_ADDR)
	return http.ListenAndServe(config.HTTP_SERVER_ADDR, mux)
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
