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
	inMemRepo := adapters.NewInMemoryRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{
			CreateOrder: command.NewCreateOrderHandler(inMemRepo, logger, metricsClient),
			UpdateOrder: command.NewUpdateOrderHandler(inMemRepo, logger, metricsClient),
		},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(inMemRepo, logger, metricsClient),
		},
	}
}
