package ordering

import (
	"DDDSample/internal/ddd"
	"DDDSample/internal/monolith"
	"DDDSample/ordering/internal/application"
	"DDDSample/ordering/internal/infrastructure/persistence/memory"
	"DDDSample/ordering/internal/presenters/grpc"
	"DDDSample/ordering/internal/presenters/rest"
	"context"
	"fmt"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// 1. 注册领域事件处理器
	domainDispatcher := ddd.NewEventDispatcher()
	orders := memory.NewOrderRepository()

	var app application.App

	app = application.New(orders, domainDispatcher)
	fmt.Println("Ordering module started")

	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		fmt.Println(err)
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		fmt.Println(err)
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
