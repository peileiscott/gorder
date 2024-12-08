package command

import (
	"context"

	"github.com/peileiscott/gorder/common/decorator"
	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/order/domain"
	"github.com/sirupsen/logrus"
)

type CreateOrderRequest struct {
	CustomerID string
	Items      []*orderpb.ItemWithQuantity
}

type CreateOrderHandler decorator.CommandHandler[CreateOrderRequest]

func NewCreateOrderHandler(
	orderRepo domain.OrderRepository,
	logger *logrus.Entry,
	client decorator.MetricsClient,
) CreateOrderHandler {
	if orderRepo == nil {
		logrus.Panic("orderRepo is nil")
	}

	return decorator.ApplyCommandDecorators(
		createOrderHandler{orderRepo: orderRepo},
		logger,
		client,
	)
}

type createOrderHandler struct {
	orderRepo domain.OrderRepository
	// TODO: Stock GRPC
}

func (h createOrderHandler) Handle(ctx context.Context, command CreateOrderRequest) error {
	// TODO: Call Stock GRPC to get items
	var stockResponse []*orderpb.Item
	for _, item := range command.Items {
		stockResponse = append(stockResponse, &orderpb.Item{
			ID:       item.ID,
			Quantity: item.Quantity,
		})
	}

	return h.orderRepo.CreateOrder(ctx, &domain.Order{
		CustomerID: command.CustomerID,
		Items:      stockResponse,
	})
}
