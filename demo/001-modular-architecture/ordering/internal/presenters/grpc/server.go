package grpc

import (
	"DDDSample/ordering/internal/application"
	"DDDSample/ordering/orderingpb"
)

type Server struct {
	app application.App
	orderingpb.UnimplementedOrderServiceServer
}

var _ orderingpb.OrderServiceClient
