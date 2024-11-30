package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/peileiscott/gorder/common/config"
	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/common/server"
	"github.com/peileiscott/gorder/order/ports"
	"github.com/peileiscott/gorder/order/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := service.NewApplication(ctx)

	go server.RunGRPCServer(serviceName, func(s *grpc.Server) {
		orderpb.RegisterOrderServiceServer(s, ports.NewGRPCServer(app))
	})

	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, ports.NewHTTPServer(app), ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
