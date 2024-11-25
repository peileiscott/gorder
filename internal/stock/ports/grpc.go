package ports

import (
	"context"

	"github.com/peileiscott/gorder/common/genproto/stockpb"
	"github.com/peileiscott/gorder/stock/app"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (G GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) CheckItemsInStock(ctx context.Context, request *stockpb.CheckItemsInStockRequest) (*stockpb.CheckItemsInStockResponse, error) {
	//TODO implement me
	panic("implement me")
}
