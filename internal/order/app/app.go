package app

import (
	"github.com/peileiscott/gorder/order/app/command"
	"github.com/peileiscott/gorder/order/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateOrder command.CreateOrderHandler
	UpdateOrder command.UpdateOrderHandler
}

type Queries struct {
	GetOrder query.GetOrderHandler
}
