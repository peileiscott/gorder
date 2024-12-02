package ports

import (
	"context"

	"github.com/peileiscott/gorder/common/genproto/stockpb"
	"github.com/peileiscott/gorder/stock/app"
	"github.com/peileiscott/gorder/stock/app/query"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (s GRPCServer) GetItems(ctx context.Context, request *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	items, err := s.app.Queries.GetItems.Handle(ctx, query.GetItems{ItemIDs: request.ItemIDs})
	if err != nil {
		return nil, err
	}
	return &stockpb.GetItemsResponse{Items: items}, nil
}

func (s GRPCServer) CheckIfItemsInStock(ctx context.Context, request *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	items, err := s.app.Queries.CheckIfItemsInStock.Handle(ctx, query.CheckIfItemsInStock{Items: request.Items})
	if err != nil {
		return nil, err
	}
	return &stockpb.CheckIfItemsInStockResponse{
		InStock: true,
		Items:   items,
	}, nil
}
