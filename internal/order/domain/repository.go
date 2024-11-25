package domain

import (
	"context"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, order *Order) (*Order, error)
	Get(ctx context.Context, orderID, customerID string) (*Order, error)
	Update(ctx context.Context, order *Order, updateFn func(context.Context, *Order) (*Order, error)) error
}

type NotFoundError struct {
	OrderID string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("order %s not found", e.OrderID)
}
