package queries

import (
	"DDDSample/ordering/internal/domain"
	"context"
)

type GetOrder struct {
	ID string
}

type GetOrderHandler struct {
	orders domain.OrderRepository
}

func NewGetOrderHandler(orders domain.OrderRepository) GetOrderHandler {
	return GetOrderHandler{
		orders: orders,
	}
}

func (h GetOrderHandler) GetOrder(ctx context.Context, query GetOrder) (*domain.Order, error) {
	order, err := h.orders.FindByID(ctx, query.ID)
	return order, err
}
