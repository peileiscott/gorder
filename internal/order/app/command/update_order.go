package command

import (
	"context"

	"github.com/peileiscott/gorder/common/decorator"
	"github.com/peileiscott/gorder/order/domain"
	"github.com/sirupsen/logrus"
)

type UpdateOrder struct {
	Order    *domain.Order
	UpdateFn func(context.Context, *domain.Order) (*domain.Order, error)
}

type UpdateOrderHandler decorator.CommandHandler[UpdateOrder, any]

type updateOrderHandler struct {
	repo domain.Repository
	// stock gRPC
}

func NewUpdateOrderHandler(
	repo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) UpdateOrderHandler {
	if repo == nil {
		panic("nil repo")
	}

	return decorator.ApplyCommandDecorators(
		updateOrderHandler{repo: repo},
		logger,
		metricsClient,
	)
}

func (h updateOrderHandler) Handle(ctx context.Context, command UpdateOrder) (any, error) {
	if command.UpdateFn == nil {
		logrus.Warnf("updateOrderHandler.Handle: UpdateFn is nil, order=%+v", command.Order)
		command.UpdateFn = func(ctx context.Context, order *domain.Order) (*domain.Order, error) {
			return order, nil
		}
	}

	if err := h.repo.Update(ctx, command.Order, command.UpdateFn); err != nil {
		return nil, err
	}
	return nil, nil
}
