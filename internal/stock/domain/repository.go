package domain

import (
	"context"
	"fmt"
	"strings"

	"github.com/peileiscott/gorder/common/genproto/orderpb"
)

type Repository interface {
	GetItems(ctx context.Context, itemIDs []string) ([]*orderpb.Item, error)
}

type NotFoundError struct {
	MissingItemIDs []string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("missing items: %s", strings.Join(e.MissingItemIDs, ", "))
}
