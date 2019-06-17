package main

import (
	"github.com/gin-gonic/gin"
	"github.com/BiteBit/grpc-gateway-example/api"
	"github.com/BiteBit/grpc-gateway-example/services"
)

func main() {
	r := gin.Default()

	srv := &services.EchoServices{}
	s := api.NewEchoServiceGinServer(srv)
	s.RegisterEchoServiceHander(r)

	r.Run(":8080")
}
