package application

import (
	"DDDSample/internal/ddd"
	"DDDSample/ordering/internal/application/commands"
	"DDDSample/ordering/internal/application/queries"
	"DDDSample/ordering/internal/domain"
	"context"
)

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateOrder(ctx context.Context, cmd commands.CreateOrder) error
	}
	Queries interface {
		GetOrder(ctx context.Context, query queries.GetOrder) (*domain.Order, error)
	}

	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateOrderHandler
	}
	appQueries struct {
		queries.GetOrderHandler
	}
)

var _ App = (*Application)(nil)

func New(orders domain.OrderRepository, domainPublisher ddd.EventPublisher,
) *Application {
	return &Application{
		appCommands: appCommands{
			CreateOrderHandler: commands.NewCreateOrderHandler(orders, domainPublisher),
		},
		appQueries: appQueries{
			GetOrderHandler: queries.NewGetOrderHandler(orders),
		},
	}
}
