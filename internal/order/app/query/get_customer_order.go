package query

import (
	"context"

	"github.com/peileiscott/gorder/common/decorator"
	"github.com/peileiscott/gorder/order/domain"
	"github.com/sirupsen/logrus"
)

type GetCustomerOrder struct {
	OrderID    string
	CustomerID string
}

type GetCustomerOrderHandler decorator.QueryHandler[GetCustomerOrder, *domain.Order]

type getCustomerOrderHandler struct {
	repo domain.Repository
}

func NewGetCustomerOrderHandler(
	repo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) GetCustomerOrderHandler {
	if repo == nil {
		panic("nil repo")
	}
	return decorator.ApplyQueryDecorators(
		getCustomerOrderHandler{repo: repo},
		logger,
		metricsClient,
	)
}

func (h getCustomerOrderHandler) Handle(ctx context.Context, query GetCustomerOrder) (*domain.Order, error) {
	order, err := h.repo.Get(ctx, query.OrderID, query.CustomerID)
	if err != nil {
		return nil, err
	}
	return order, nil
}
