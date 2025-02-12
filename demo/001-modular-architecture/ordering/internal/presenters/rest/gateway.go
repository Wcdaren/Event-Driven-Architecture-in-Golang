package rest

import (
	orderingv1 "DDDSample/ordering/orderingpb/v1"
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGateway(ctx context.Context, mux *chi.Mux, grpcAddr string) error {
	const apiRoot = "/api/ordering"
	// https://github.com/grpc-ecosystem/grpc-gateway?tab=readme-ov-file#5-write-an-entrypoint-for-the-http-reverse-proxy-server
	gateway := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := orderingv1.RegisterOrderingServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, opts)
	if err != nil {
		return err
	}

	// mount the GRPC gateway
	mux.Mount(apiRoot, gateway)

	return nil

}
