package app

import (
	grpcapp "server/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(grpcPort int) *App {
	grpcApp := grpcapp.New(grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
