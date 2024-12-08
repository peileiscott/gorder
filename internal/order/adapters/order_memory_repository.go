package adapters

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/peileiscott/gorder/order/domain"
	"github.com/sirupsen/logrus"
)

type MemoryOrderRepository struct {
	lock  *sync.RWMutex
	store []*domain.Order
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		lock:  &sync.RWMutex{},
		store: make([]*domain.Order, 0),
	}
}

func (m *MemoryOrderRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	newOrder := &domain.Order{
		ID:          strconv.FormatInt(time.Now().Unix(), 10),
		CustomerID:  order.CustomerID,
		Status:      order.Status,
		Items:       order.Items,
		PaymentLink: order.PaymentLink,
	}
	m.store = append(m.store, newOrder)
	logrus.WithFields(logrus.Fields{
		"input_order":        *order,
		"store_after_create": m.store,
	}).Debug("MemoryOrderRepository.CreateOrder")
	return nil
}

func (m *MemoryOrderRepository) GetOrder(ctx context.Context, orderID, customerID string) (*domain.Order, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	order, _, err := m.getOrder(orderID, customerID)
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"order": *order,
	}).Debug("MemoryOrderRepository.GetOrder")
	return order, nil
}

func (m *MemoryOrderRepository) UpdateOrder(ctx context.Context, order *domain.Order, updateFn func(context.Context, *domain.Order) (*domain.Order, error)) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	currentOrder, idx, err := m.getOrder(order.ID, order.CustomerID)
	if err != nil {
		return err
	}

	updatedOrder, err := updateFn(ctx, currentOrder)
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"order_before_update": *currentOrder,
		"order_after_update":  *updatedOrder,
	}).Debug("MemoryOrderRepository.UpdateOrder")
	m.store[idx] = updatedOrder
	return nil
}

func (m *MemoryOrderRepository) getOrder(orderID, customerID string) (*domain.Order, int, error) {
	for i, order := range m.store {
		if order.ID == orderID && order.CustomerID == customerID {
			return order, i, nil
		}
	}

	return nil, 0, domain.OrderNotFoundError{OrderID: orderID}
}
