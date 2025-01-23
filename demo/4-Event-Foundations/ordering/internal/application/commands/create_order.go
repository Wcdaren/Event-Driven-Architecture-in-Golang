// internal/ordering/application/commands/create_order.go
package commands

import (
	"DDDSample/internal/ddd"
	"DDDSample/ordering/internal/domain"
	"context"
)

type CreateOrder struct {
	ID         string
	CustomerID string
	Items      []*domain.Item
}

type CreateOrderHandler struct {
	orders          domain.OrderRepository
	domainPublisher ddd.EventPublisher
}

func NewCreateOrderHandler(orders domain.OrderRepository, domainPublisher ddd.EventPublisher) CreateOrderHandler {
	return CreateOrderHandler{
		orders:          orders,
		domainPublisher: domainPublisher,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := domain.CreateOrder(cmd.ID, cmd.Items)
	if err != nil {
		return err
	}

	// 保存订单
	if err = h.orders.Save(ctx, order); err != nil {
		return err
	}

	// 发布领域事件
	if err = h.domainPublisher.Publish(ctx, order.GetEvents()...); err != nil {
		return err
	}

	return nil
}
