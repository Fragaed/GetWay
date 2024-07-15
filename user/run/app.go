package run

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"

	"google.golang.org/grpc"
	"net"
	controller "user/internal/modules/gRPC"
	"user/internal/modules/service"
	"user/internal/modules/storage"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewApp(log *slog.Logger, port int, db *sqlx.DB) *App {
	gRPCServer := grpc.NewServer()
	UStorage := storage.NewUserStorage(db)
	UService := service.NewUserService(UStorage)

	controller.Register(gRPCServer, UService)
	controller.NewServerAPI(UService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}

}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("operation", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("address", l.Addr().String()))
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"
	a.log.With(slog.String("operation", op)).
		Info("grpc server is stopping", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
