package handlers

import (
	"DDDSample/internal/ddd"
	"DDDSample/ordering/internal/domain"
	"context"
	"fmt"
)

type NotificationHandlers struct {
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*NotificationHandlers)(nil)

func NewNotificationHandlers() *NotificationHandlers {
	return &NotificationHandlers{}
}

func (h NotificationHandlers) OnOrderCreated(ctx context.Context, event ddd.Event) error {
	orderCreated := event.(*domain.OrderCreated)
	// In real application, this would send an actual notification
	fmt.Printf("Order %s created with status %s and items %v\n",
		orderCreated.Order.ID,
		orderCreated.Order.Status,
		orderCreated.Order.Items,
	)
	return nil
}

func RegisterNotificationHandlers(notificationHandlers DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderCreated{}, notificationHandlers.OnOrderCreated)
}
