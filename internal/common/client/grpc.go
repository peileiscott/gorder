package client

import (
	"context"

	"github.com/peileiscott/gorder/common/genproto/stockpb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewStockServiceClient(ctx context.Context) (stockpb.StockServiceClient, func() error, error) {
	grpcAddr := viper.GetString("stock.grpc-addr")

	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	return stockpb.NewStockServiceClient(conn), conn.Close, nil
}
