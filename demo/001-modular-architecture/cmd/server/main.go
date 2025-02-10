package main

import (
	"DDDSample/internal/config"
	"DDDSample/internal/logger"
	"DDDSample/internal/monolith"
	"DDDSample/internal/rpc"
	"DDDSample/internal/waiter"
	"DDDSample/internal/web"
	"DDDSample/ordering"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-chi/chi"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	m := app{cfg: cfg}

	m.db, err = sql.Open("pgx", cfg.PG.Conn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}

	}(m.db)

	m.logger = logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	}).With().Timestamp().Logger()

	m.rpc = initRpc(cfg.Rpc)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	// init modules
	m.modules = []monolith.Module{
		&ordering.Module{},
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	fmt.Println("server started")
	defer fmt.Println("server stopped")

	m.waiter.Add(
		m.waitForWeb,
		m.waitForRPC,
	)

	return m.waiter.Wait()
}

func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}
