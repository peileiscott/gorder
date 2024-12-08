package query

import (
	"context"

	"github.com/peileiscott/gorder/common/decorator"
	"github.com/peileiscott/gorder/order/domain"
	"github.com/sirupsen/logrus"
)

type GetOrderRequest struct {
	OrderID    string
	CustomerID string
}

type GetOrderHandler decorator.QueryHandler[GetOrderRequest, *domain.Order]

func NewGetOrderHandler(
	orderRepo domain.OrderRepository,
	logger *logrus.Entry,
	client decorator.MetricsClient,
) GetOrderHandler {
	if orderRepo == nil {
		logrus.Panic("orderRepo is nil")
	}

	return decorator.ApplyQueryDecorators(
		getOrderHandler{orderRepo: orderRepo},
		logger,
		client,
	)
}

type getOrderHandler struct {
	orderRepo domain.OrderRepository
}

func (h getOrderHandler) Handle(ctx context.Context, query GetOrderRequest) (*domain.Order, error) {
	order, err := h.orderRepo.GetOrder(ctx, query.OrderID, query.CustomerID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
