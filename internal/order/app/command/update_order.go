package command

import (
	"context"

	"github.com/peileiscott/gorder/common/decorator"
	"github.com/peileiscott/gorder/order/domain"
	"github.com/sirupsen/logrus"
)

type UpdateOrderRequest struct {
	Order    *domain.Order
	UpdateFn func(context.Context, *domain.Order) (*domain.Order, error)
}

type UpdateOrderHandler decorator.CommandHandler[UpdateOrderRequest]

func NewUpdateOrderHandler(
	orderRepo domain.OrderRepository,
	logger *logrus.Entry,
	client decorator.MetricsClient,
) UpdateOrderHandler {
	if orderRepo == nil {
		logrus.Panic("orderRepo is nil")
	}

	return decorator.ApplyCommandDecorators(
		updateOrderHandler{orderRepo: orderRepo},
		logger,
		client,
	)
}

type updateOrderHandler struct {
	orderRepo domain.OrderRepository
}

func (h updateOrderHandler) Handle(ctx context.Context, command UpdateOrderRequest) error {
	if command.UpdateFn == nil {
		logrus.WithFields(logrus.Fields{
			"order": *command.Order,
		}).Warnf("UpdateOrderHandler.Handle: received nil UpdateFn")

		return nil
	}

	return h.orderRepo.UpdateOrder(ctx, command.Order, command.UpdateFn)
}
