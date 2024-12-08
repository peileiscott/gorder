package adapters

import (
	"context"
	"sync"

	"github.com/peileiscott/gorder/common/genproto/orderpb"
	"github.com/peileiscott/gorder/stock/domain"
	"github.com/sirupsen/logrus"
)

type MemoryStockRepository struct {
	lock  *sync.RWMutex
	store map[string]*orderpb.Item
}

func NewMemoryStockRepository() *MemoryStockRepository {
	return &MemoryStockRepository{
		lock:  &sync.RWMutex{},
		store: make(map[string]*orderpb.Item),
	}
}

func (m MemoryStockRepository) GetItems(ctx context.Context, itemIDs []string) ([]*orderpb.Item, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	var items []*orderpb.Item
	var missingItemIDs []string
	for _, id := range itemIDs {
		item, ok := m.store[id]
		if ok {
			items = append(items, item)
		} else {
			missingItemIDs = append(missingItemIDs, id)
		}
	}

	if len(missingItemIDs) > 0 {
		return nil, domain.ItemNotFoundError{ItemIDs: missingItemIDs}
	}

	logrus.WithFields(logrus.Fields{
		"items": items,
	}).Debug("MemoryStockRepository.GetItems")
	return items, nil
}
