package ordering

import (
	"DDDSample/internal/ddd"
	"DDDSample/internal/monolith"
	"DDDSample/ordering/internal/application"
	"DDDSample/ordering/internal/application/handlers"
	"DDDSample/ordering/internal/infrastructure/memory"
	"context"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Module) {
	// 1. 注册领域事件处理器
	domainDispatcher := ddd.NewEventDispatcher()
	orders := memory.NewOrderRepository()

	var app application.App

	app = application.New(orders, domainDispatcher)

	handlers.RegisterNotificationHandlers(handlers.NewNotificationHandlers(), domainDispatcher)

	return nil

	// 2. 注册命令处理器
	// 3. 注册查询处理器
	// 4. 注册HTTP处理器
	// 5. 注册gRPC处理器

}
