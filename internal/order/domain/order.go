package domain

import (
	"context"
	"fmt"

	"github.com/peileiscott/gorder/common/genproto/orderpb"
)

type Order struct {
	ID          string
	CustomerID  string
	Status      string
	Items       []*orderpb.Item
	PaymentLink string
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) error
	GetOrder(ctx context.Context, orderID, customerID string) (*Order, error)
	UpdateOrder(ctx context.Context, order *Order, updateFn func(context.Context, *Order) (*Order, error)) error
}

type OrderNotFoundError struct {
	OrderID string
}

func (e OrderNotFoundError) Error() string {
	return fmt.Sprintf("order %s not found", e.OrderID)
}
