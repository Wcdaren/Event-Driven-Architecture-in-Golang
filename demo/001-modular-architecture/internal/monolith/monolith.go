package monolith

import (
	"DDDSample/internal/config"
	"DDDSample/internal/waiter"
	"context"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
	Logger() zerolog.Logger
	Mux() *chi.Mux
	RPC() *grpc.Server
	Waiter() waiter.Waiter
}

type Module interface {
	Startup(context.Context, Monolith) error
}
