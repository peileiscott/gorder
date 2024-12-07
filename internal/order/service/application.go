package service

import (
	"context"

	"github.com/peileiscott/gorder/order/app"
)

func NewApplication(ctx context.Context) app.Application {
	return app.Application{
		Commands: app.Commands{},
		Queries:  app.Queries{},
	}
}
