package main

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/app"
	"github.com/CatMacales/route256/loms/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"

	stocksInitPath = "stock-data.json"
)

func main() {
	cfg := config.MustLoad()

	application := app.New(cfg.Grpc.Host, cfg.HTTP.Host, cfg.Grpc.Port, cfg.HTTP.Port)

	application.MustInitStocks(context.Background(), stocksInitPath)

	go func() {
		application.GRPCServer.MustRun()
	}()

	application.HttpGateway.MustConnect()
	err := application.HttpGateway.Serve()
	if err != nil {
		panic(err)
	}
}
