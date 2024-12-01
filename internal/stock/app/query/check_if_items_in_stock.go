package query

import (
	"context"

	"github.com/peileiscott/gorder/common/decorator"
	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/stock/domain"
	"github.com/sirupsen/logrus"
)

type CheckIfItemsInStock struct {
	Items []*orderpb.ItemWithQuantity
}

type CheckIfItemsInStockHandler decorator.QueryHandler[CheckIfItemsInStock, []*orderpb.Item]

type checkIfItemsInStockHandler struct {
	repo domain.Repository
}

func NewCheckIfItemsInStockHandler(
	repo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CheckIfItemsInStockHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators(
		checkIfItemsInStockHandler{repo: repo},
		logger,
		metricsClient,
	)
}

func (h checkIfItemsInStockHandler) Handle(ctx context.Context, query CheckIfItemsInStock) ([]*orderpb.Item, error) {
	var res []*orderpb.Item
	for _, item := range query.Items {
		res = append(res, &orderpb.Item{
			ID:       item.ID,
			Quantity: item.Quantity,
		})
	}
	return res, nil
}
