package inmemory

import (
	"DDDSample/internal/ddd"
	"DDDSample/ordering/internal/domain"
	"context"
	"sync"

	"github.com/stackus/errors"
)

type OrderRepository struct {
	orders map[string]*domain.Order
	sync.RWMutex
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]*domain.Order),
	}
}

func (r *OrderRepository) Find(ctx context.Context, orderID string) (*domain.Order, error) {
	r.RLock()
	defer r.RUnlock()

	if order, exists := r.orders[orderID]; exists {
		// Return a deep copy to prevent external modifications
		items := make([]*domain.Item, len(order.Items))
		copy(items, order.Items)

		return &domain.Order{
			AggregateBase: ddd.AggregateBase{
				ID: order.ID,
			},
			// CustomerID: order.CustomerID,
			// PaymentID:  order.PaymentID,
			// ShoppingID: order.ShoppingID,
			// InvoiceID:  order.InvoiceID,
			Items:  items,
			Status: order.Status,
		}, nil
	}

	return nil, errors.Wrap(errors.ErrNotFound, "order not found")
}

func (r *OrderRepository) Save(ctx context.Context, order *domain.Order) error {
	r.Lock()
	defer r.Unlock()

	// Create a deep copy of the order to store
	items := make([]*domain.Item, len(order.Items))
	copy(items, order.Items)

	stored := &domain.Order{
		AggregateBase: ddd.AggregateBase{
			ID: order.ID,
		},
		// CustomerID: order.CustomerID,
		// PaymentID:  order.PaymentID,
		// ShoppingID: order.ShoppingID,
		// InvoiceID:  order.InvoiceID,
		Items:  items,
		Status: order.Status,
	}

	r.orders[order.ID] = stored
	return nil
}

func (r *OrderRepository) Update(ctx context.Context, order *domain.Order) error {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.orders[order.ID]; !exists {
		return errors.Wrap(errors.ErrNotFound, "order not found")
	}

	// Create a deep copy of the order to store
	items := make([]*domain.Item, len(order.Items))
	copy(items, order.Items)

	stored := &domain.Order{
		AggregateBase: ddd.AggregateBase{
			ID: order.ID,
		},
		// CustomerID: order.CustomerID,
		// PaymentID:  order.PaymentID,
		// ShoppingID: order.ShoppingID,
		// InvoiceID:  order.InvoiceID,
		// Items:  items,
		Status: order.Status,
	}

	r.orders[order.ID] = stored
	return nil
}
