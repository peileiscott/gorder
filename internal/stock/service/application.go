package service

import (
	"context"

	"github.com/peileiscott/gorder/common/metrics"
	"github.com/peileiscott/gorder/stock/adapters"
	"github.com/peileiscott/gorder/stock/app"
	"github.com/peileiscott/gorder/stock/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	inMemRepo := adapters.NewInMemoryRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetItems:            query.NewGetItemsHandler(inMemRepo, logger, metricsClient),
			CheckIfItemsInStock: query.NewCheckIfItemsInStockHandler(inMemRepo, logger, metricsClient),
		},
	}
}
