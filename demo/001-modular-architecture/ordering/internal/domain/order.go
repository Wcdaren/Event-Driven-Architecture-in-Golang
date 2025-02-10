package domain

import (
	"DDDSample/internal/ddd"

	"github.com/stackus/errors"
)

var (
	ErrOrderHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the order has no items")
	ErrOrderCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the order cannot be cancelled")
	ErrCustomerIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	ErrPaymentIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
)

type Order struct {
	ddd.AggregateBase
	Items  []*Item
	Status OrderStatus
}

type Item struct {
	ProductID string
	Price     float64
	Quantity  int
}

type OrderStatus string

const (
	OrderUnknown     OrderStatus = ""
	OrderIsPending   OrderStatus = "pending"
	OrderIsInProcess OrderStatus = "in-progress"
	OrderIsReady     OrderStatus = "ready"
	OrderIsCompleted OrderStatus = "completed"
	OrderIsCancelled OrderStatus = "cancelled"
)

func CreateOrder(id string, items []*Item) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	order := &Order{
		AggregateBase: ddd.AggregateBase{
			ID: id,
		},
		Items:  items,
		Status: OrderIsPending,
	}

	order.AddEvent(&OrderCreated{
		Order: order,
	})

	return order, nil
}

func (o *Order) Cancel() error {
	if o.Status != OrderIsPending {
		return ErrOrderCannotBeCancelled
	}

	o.Status = OrderIsCancelled

	o.AddEvent(&OrderCanceled{
		Order: o,
	})
	return nil
}
