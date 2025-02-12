package grpc

import (
	"DDDSample/ordering/internal/application"
	"DDDSample/ordering/internal/application/queries"
	"DDDSample/ordering/internal/domain"
	orderingv1 "DDDSample/ordering/orderingpb/v1"
	"context"

	"google.golang.org/grpc"
)

type server struct {
	app application.App
	orderingv1.UnimplementedOrderingServiceServer
}

var _ orderingv1.OrderingServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	orderingv1.RegisterOrderingServiceServer(registrar, server{app: app})
	return nil
}

func (s server) GetOrder(ctx context.Context, request *orderingv1.GetOrderRequest) (*orderingv1.GetOrderResponse, error) {
	order, err := s.app.GetOrder(ctx, queries.GetOrder{ID: request.GetId()})
	if err != nil {
		return nil, err
	}
	return &orderingv1.GetOrderResponse{
		Order: s.orderFromDomain(order),
	}, nil
}

func (s server) orderFromDomain(order *domain.Order) *orderingv1.Order {
	items := make([]*orderingv1.Item, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, s.itemFromDomain(item))
	}

	return &orderingv1.Order{
		Id:         order.ID,
		CustomerId: order.CustomerID,
		PaymentId:  order.PaymentID,
		Items:      items,
		Status:     order.Status.String(),
	}
}

func (s server) itemFromDomain(item *domain.Item) *orderingv1.Item {
	return &orderingv1.Item{
		StoreId:     item.StoreID,
		ProductId:   item.ProductID,
		StoreName:   item.StoreName,
		ProductName: item.ProductName,
		Price:       item.Price,
		Quantity:    int32(item.Quantity),
	}
}
