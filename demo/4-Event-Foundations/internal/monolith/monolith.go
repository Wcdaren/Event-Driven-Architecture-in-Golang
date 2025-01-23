package monolith

import (
	"context"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type Monolith interface {
	// Config() config.AppConfig
	// DB() *sql.DB
	// Logger() zerolog.Logger
	Mux() *chi.Mux
	RPC() *grpc.Server
	// Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, Monolith) error
}
