package main

import (
	"log"

	"github.com/peileiscott/gorder/common/config"
	"github.com/peileiscott/gorder/common/genproto/stockpb"
	"github.com/peileiscott/gorder/common/server"
	"github.com/peileiscott/gorder/stock/ports"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")
	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			stockpb.RegisterStockServiceServer(server, ports.NewGRPCServer())
		})
	default:
		panic("unsupported server type")
	}
}
