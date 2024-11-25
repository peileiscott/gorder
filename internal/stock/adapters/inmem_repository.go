package adapters

import (
	"context"
	"sync"

	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/stock/domain"
)

type InMemoryRepository struct {
	lock  *sync.RWMutex
	store map[string]*orderpb.Item
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		lock:  &sync.RWMutex{},
		store: make(map[string]*orderpb.Item),
	}
}

func (r InMemoryRepository) GetItems(ctx context.Context, itemIDs []string) ([]*orderpb.Item, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	var (
		items          []*orderpb.Item
		missingItemIDs []string
	)
	for _, id := range itemIDs {
		if item, exist := r.store[id]; exist {
			items = append(items, item)
		} else {
			missingItemIDs = append(missingItemIDs, id)
		}
	}

	if len(missingItemIDs) == 0 {
		return items, nil
	}
	return items, domain.NotFoundError{MissingItemIDs: missingItemIDs}
}
