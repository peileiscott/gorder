package query

import (
	"context"

	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/common/genproto/stockpb"
)

type StockService interface {
	GetItems(ctx context.Context, itemIDs []string) ([]*orderpb.Item, error)
	CheckIfItemsInStock(ctx context.Context, items []*orderpb.ItemWithQuantity) (*stockpb.CheckIfItemsInStockResponse, error)
}
