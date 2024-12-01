package main

import (
	"context"

	"github.com/peileiscott/gorder/common/config"
	"github.com/peileiscott/gorder/common/discovery"
	"github.com/peileiscott/gorder/common/genproto/stockpb"
	"github.com/peileiscott/gorder/common/server"
	"github.com/peileiscott/gorder/stock/ports"
	"github.com/peileiscott/gorder/stock/service"
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
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := service.NewApplication(ctx)

	deregisterFunc, err := discovery.RegisterToConsul(serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		_ = deregisterFunc()
	}()

	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(s *grpc.Server) {
			stockpb.RegisterStockServiceServer(s, ports.NewGRPCServer(app))
		})
	default:
		panic("unsupported server type")
	}
}
