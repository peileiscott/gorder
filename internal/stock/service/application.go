package service

import (
	"context"

	"github.com/peileiscott/gorder/stock/app"
)

func NewApplication(ctx context.Context) app.Application {
	return app.Application{
		Commands: app.Commands{},
		Queries:  app.Queries{},
	}
}
