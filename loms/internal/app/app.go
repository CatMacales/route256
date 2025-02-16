package app

import (
	"github.com/CatMacales/route256/loms/internal/app/grpc"
	"github.com/CatMacales/route256/loms/internal/repository/memory/order"
	"github.com/CatMacales/route256/loms/internal/repository/memory/stock"
	"github.com/CatMacales/route256/loms/internal/service/loms"
)

type App struct {
	GRPCServer *grpc_app.App
}

func New(host string, port uint32) *App {
	orderRepository := order_repository.NewRepository()
	stockRepository := stock_repository.NewRepository()

	lomsService := loms.NewService(orderRepository, stockRepository)

	grpcApp := grpc_app.New(lomsService, host, port)

	return &App{
		GRPCServer: grpcApp,
	}
}
