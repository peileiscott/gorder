package command

import (
	"context"

	"github.com/peileiscott/gorder/common/decorator"
	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/order/domain"
	"github.com/sirupsen/logrus"
)

type CreateOrder struct {
	CustomerID string
	Items      []*orderpb.ItemWithQuantity
}

type CreateOrderResponse struct {
	OrderID string
}

type CreateOrderHandler decorator.CommandHandler[CreateOrder, *CreateOrderResponse]

type createOrderHandler struct {
	repo domain.Repository
	// stock gRPC
}

func NewCreateOrderHandler(
	repo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CreateOrderHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators(
		createOrderHandler{repo: repo},
		logger,
		metricsClient,
	)
}

func (h createOrderHandler) Handle(ctx context.Context, command CreateOrder) (*CreateOrderResponse, error) {
	// TODO: call stock gRPC to get items
	var stockResponse []*orderpb.Item
	for _, item := range command.Items {
		stockResponse = append(stockResponse, &orderpb.Item{
			ID:       item.ID,
			Quantity: item.Quantity,
		})
	}

	order, err := h.repo.Create(ctx, &domain.Order{
		CustomerID: command.CustomerID,
		Items:      stockResponse,
	})
	if err != nil {
		return nil, err
	}
	return &CreateOrderResponse{OrderID: order.ID}, nil
}
