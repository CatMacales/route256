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

	application := app.New(cfg.Host, cfg.Port)

	application.MustInitStocks(context.Background(), stocksInitPath)

	application.GRPCServer.MustRun()
}
