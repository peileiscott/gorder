package main

import (
	"context"

	"github.com/peileiscott/gorder/common/config"
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
		logrus.Panic(err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service_name")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)

	server.RunGRPCServer(serviceName, func(s *grpc.Server) {
		stockpb.RegisterStockServiceServer(s, ports.NewGRPCServer(application))
	})
}
