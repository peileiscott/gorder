package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/peileiscott/gorder/common/config"
	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/common/server"
	"github.com/peileiscott/gorder/order/ports"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")

	go server.RunGRPCServer(serviceName, func(server *grpc.Server) {
		orderpb.RegisterOrderServiceServer(server, ports.NewGRPCServer())
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, ports.HTTPServer{}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
