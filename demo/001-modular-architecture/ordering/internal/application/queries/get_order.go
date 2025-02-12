package queries

import (
	"DDDSample/ordering/internal/domain"
	"context"
)

type GetOrder struct {
	ID string
}

type GetOrderHandler struct {
	repo domain.OrderRepository
}

func NewGetOrderHandler(orders domain.OrderRepository) GetOrderHandler {
	return GetOrderHandler{
		repo: orders,
	}
}

func (h GetOrderHandler) GetOrder(ctx context.Context, query GetOrder) (*domain.Order, error) {
	order, err := h.repo.Find(ctx, query.ID)
	return order, err
}
