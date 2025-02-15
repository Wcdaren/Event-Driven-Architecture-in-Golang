// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: orderingpb/v1/api.proto

package orderingv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderingServiceClient is the client API for OrderingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderingServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
}

type orderingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderingServiceClient(cc grpc.ClientConnInterface) OrderingServiceClient {
	return &orderingServiceClient{cc}
}

func (c *orderingServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/orderingpb.v1.OrderingService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderingServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/orderingpb.v1.OrderingService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderingServiceServer is the server API for OrderingService service.
// All implementations must embed UnimplementedOrderingServiceServer
// for forward compatibility
type OrderingServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	mustEmbedUnimplementedOrderingServiceServer()
}

// UnimplementedOrderingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderingServiceServer struct {
}

func (UnimplementedOrderingServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderingServiceServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderingServiceServer) mustEmbedUnimplementedOrderingServiceServer() {}

// UnsafeOrderingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderingServiceServer will
// result in compilation errors.
type UnsafeOrderingServiceServer interface {
	mustEmbedUnimplementedOrderingServiceServer()
}

func RegisterOrderingServiceServer(s grpc.ServiceRegistrar, srv OrderingServiceServer) {
	s.RegisterService(&OrderingService_ServiceDesc, srv)
}

func _OrderingService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderingpb.v1.OrderingService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderingService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderingpb.v1.OrderingService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderingService_ServiceDesc is the grpc.ServiceDesc for OrderingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orderingpb.v1.OrderingService",
	HandlerType: (*OrderingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderingService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderingService_GetOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orderingpb/v1/api.proto",
}
