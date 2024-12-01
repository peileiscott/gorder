package query

import (
	"context"

	"github.com/peileiscott/gorder/common/decorator"
	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/stock/domain"
	"github.com/sirupsen/logrus"
)

type GetItems struct {
	ItemIDs []string
}

type GetItemsHandler decorator.QueryHandler[GetItems, []*orderpb.Item]

type getItemsHandler struct {
	repo domain.Repository
}

func NewGetItemsHandler(
	repo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) GetItemsHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyQueryDecorators(
		getItemsHandler{repo: repo},
		logger,
		metricsClient,
	)
}

func (h getItemsHandler) Handle(ctx context.Context, query GetItems) ([]*orderpb.Item, error) {
	items, err := h.repo.GetItems(ctx, query.ItemIDs)
	if err != nil {
		return nil, err
	}

	return items, nil
}
