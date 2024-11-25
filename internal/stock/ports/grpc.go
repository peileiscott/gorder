package ports

import (
	"context"

	"github.com/peileiscott/gorder/common/genproto/stockpb"
)

type GRPCServer struct{}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (G GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (G GRPCServer) CheckItemsInStock(ctx context.Context, request *stockpb.CheckItemsInStockRequest) (*stockpb.CheckItemsInStockResponse, error) {
	//TODO implement me
	panic("implement me")
}
