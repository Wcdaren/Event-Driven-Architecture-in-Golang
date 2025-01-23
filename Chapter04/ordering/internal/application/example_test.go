package application

import (
	"context"
	"fmt"

	"eda-in-golang/internal/ddd"
	"eda-in-golang/ordering/internal/domain"
)

func ExampleEventHandling() {
	// 创建事件分发器
	dispatcher := ddd.NewEventDispatcher()
	
	// 注册事件处理器
	dispatcher.Subscribe(&domain.OrderCreated{}, func(ctx context.Context, event ddd.Event) error {
		fmt.Println("Handler 1: Order created event received")
		return nil
	})
	
	dispatcher.Subscribe(&domain.OrderCreated{}, func(ctx context.Context, event ddd.Event) error {
		fmt.Println("Handler 2: Order created event received")
		return nil
	})

	// 创建订单
	items := []*domain.Item{{ProductID: "prod-1", Quantity: 1, Price: 10.0}}
	order, _ := domain.CreateOrder("order-1", "cust-1", "pay-1", items)
	
	// 发布事件
	// 这会按顺序执行所有注册的handlers
	_ = dispatcher.Publish(context.Background(), order.GetEvents()...)

	// Output:
	// Handler 1: Order created event received
	// Handler 2: Order created event received
}
