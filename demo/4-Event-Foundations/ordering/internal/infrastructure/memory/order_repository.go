// internal/ordering/infrastructure/memory/order_repository.go
package memory

import (
	"DDDSample/ordering/internal/domain"
	"context"
	"sync"
)

type OrderRepository struct {
	orders map[string]*domain.Order
	mu     sync.RWMutex
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]*domain.Order),
	}
}

func (r *OrderRepository) Save(ctx context.Context, order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.ID] = order
	return nil
}

func (r *OrderRepository) FindByID(ctx context.Context, id string) (*domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if order, exists := r.orders[id]; exists {
		return order, nil
	}
	return nil, domain.ErrOrderHasNoItems
}
