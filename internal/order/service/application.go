package service

import (
	"context"

	"github.com/peileiscott/gorder/common/metrics"
	"github.com/peileiscott/gorder/order/adapters"
	"github.com/peileiscott/gorder/order/app"
	"github.com/peileiscott/gorder/order/app/command"
	"github.com/peileiscott/gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	client := metrics.TodoMetrics{}

	return app.Application{
		Commands: app.Commands{
			CreateOrder: command.NewCreateOrderHandler(orderRepo, logger, client),
			UpdateOrder: command.NewUpdateOrderHandler(orderRepo, logger, client),
		},
		Queries: app.Queries{
			GetOrder: query.NewGetOrderHandler(orderRepo, logger, client),
		},
	}
}
