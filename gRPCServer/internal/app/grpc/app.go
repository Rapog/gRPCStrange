package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"server/internal/cache"
	"server/internal/domain/models"
	"server/internal/grpc/ex00"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(port int) *App {
	gRPCServer := grpc.NewServer()
	cacheMem := cache.New[*models.MeanStd](100)
	ex00.Register(gRPCServer, cacheMem)

	return &App{
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
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
