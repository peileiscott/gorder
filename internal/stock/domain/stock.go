package domain

import (
	"context"
	"fmt"
	"strings"

	"github.com/peileiscott/gorder/common/genproto/orderpb"
)

type StockRepository interface {
	GetItems(ctx context.Context, itemIDs []string) ([]*orderpb.Item, error)
}

type ItemNotFoundError struct {
	ItemIDs []string
}

func (e ItemNotFoundError) Error() string {
	return fmt.Sprintf("items not found in stock: %s", strings.Join(e.ItemIDs, ", "))
}
