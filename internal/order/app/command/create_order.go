package command

import (
	"context"
	"errors"

	"github.com/peileiscott/gorder/common/decorator"
	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/order/app/query"
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
	repo      domain.Repository
	stockGRPC query.StockService
}

func NewCreateOrderHandler(
	repo domain.Repository,
	stockGRPC query.StockService,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) CreateOrderHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators(
		createOrderHandler{repo: repo, stockGRPC: stockGRPC},
		logger,
		metricsClient,
	)
}

func (h createOrderHandler) Handle(ctx context.Context, command CreateOrder) (*CreateOrderResponse, error) {
	items, err := h.validateItems(ctx, command.Items)
	if err != nil {
		return nil, err
	}

	order, err := h.repo.Create(ctx, &domain.Order{
		CustomerID: command.CustomerID,
		Items:      items,
	})
	if err != nil {
		return nil, err
	}
	return &CreateOrderResponse{OrderID: order.ID}, nil
}

func (h createOrderHandler) validateItems(ctx context.Context, items []*orderpb.ItemWithQuantity) ([]*orderpb.Item, error) {
	if len(items) == 0 {
		return nil, errors.New("must have at least one item")
	}

	items = groupItems(items)
	res, err := h.stockGRPC.CheckIfItemsInStock(ctx, items)
	if err != nil {
		return nil, err
	}

	return res.Items, nil
}

func groupItems(items []*orderpb.ItemWithQuantity) []*orderpb.ItemWithQuantity {
	groups := make(map[string]int32)
	for _, item := range items {
		groups[item.ID] += item.Quantity
	}

	var result []*orderpb.ItemWithQuantity
	for id, quantity := range groups {
		result = append(result, &orderpb.ItemWithQuantity{ID: id, Quantity: quantity})
	}
	return result
}
