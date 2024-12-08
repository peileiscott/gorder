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

func (s GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s GRPCServer) CheckIfItemsAvailable(ctx context.Context, request *stockpb.CheckIfItemsAvailableRequest) (*stockpb.CheckIfItemsAvailableResponse, error) {
	//TODO implement me
	panic("implement me")
}
