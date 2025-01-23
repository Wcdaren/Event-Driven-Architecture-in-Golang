// cmd/server/main.go
package main

import (
	"DDDSample/internal/ddd"
	"context"

	"github.com/google/uuid"
)

func main() {
	// Create components
	eventDispatcher := ddd.NewEventDispatcher()
	orderRepo := memory.NewOrderRepository()

	// Create handlers
	createOrderHandler := handlers.NewCreateOrderHandler(orderRepo, eventDispatcher)
	notificationHandler := handlers.NewNotificationHandler()

	// Subscribe to events
	eventDispatcher.Subscribe(
		&domain.OrderCreated{},
		notificationHandler.HandleOrderCreated,
	)

	// Create an order
	cmd := commands.CreateOrder{
		ID:         uuid.New().String(),
		CustomerID: "customer-123",
		Items: []commands.OrderItem{
			{
				ProductID: "product-1",
				Quantity:  2,
				Price:     99.99,
			},
		},
	}

	if err := createOrderHandler.Handle(context.Background(), cmd); err != nil {
		panic(err)
	}
}
