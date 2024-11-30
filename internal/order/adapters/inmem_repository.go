package adapters

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/peileiscott/gorder/order/domain"
	"github.com/sirupsen/logrus"
)

type InMemoryRepository struct {
	lock  *sync.RWMutex
	store []*domain.Order
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		lock:  &sync.RWMutex{},
		store: make([]*domain.Order, 0),
	}
}

func (r *InMemoryRepository) Create(_ context.Context, order *domain.Order) (*domain.Order, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	newOrder := &domain.Order{
		ID:          strconv.FormatInt(time.Now().Unix(), 10),
		CustomerID:  order.CustomerID,
		Status:      order.Status,
		Items:       order.Items,
		PaymentLink: order.PaymentLink,
	}
	r.store = append(r.store, newOrder)
	logrus.WithFields(logrus.Fields{
		"input_order": order,
	}).Info("InMemoryRepository.Create")
	return newOrder, nil
}

func (r *InMemoryRepository) Get(_ context.Context, orderID, customerID string) (*domain.Order, error) {
	for i, order := range r.store {
		logrus.Infof("InMemoryRepository.Get r.store[%d] = %+v\n", i, *order)
	}
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, o := range r.store {
		if o.ID == orderID && o.CustomerID == customerID {
			logrus.Infof("InMemoryRepository.Get || Found order %+v\n", *o)
			return o, nil
		}
	}
	return nil, domain.NotFoundError{OrderID: orderID}
}

func (r *InMemoryRepository) Update(ctx context.Context, order *domain.Order, updateFn func(context.Context, *domain.Order) (*domain.Order, error)) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	foundOrder := false
	for i, o := range r.store {
		if o.ID == order.ID && o.CustomerID == order.CustomerID {
			foundOrder = true
			updateOrder, err := updateFn(ctx, o)
			if err != nil {
				return err
			}
			r.store[i] = updateOrder
		}
	}

	if !foundOrder {
		return domain.NotFoundError{OrderID: order.ID}
	}
	return nil
}
