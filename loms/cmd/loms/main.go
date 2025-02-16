package main

import (
	"github.com/CatMacales/route256/loms/internal/app"
	"github.com/CatMacales/route256/loms/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	application := app.New(cfg.Host, cfg.Port)

	application.GRPCServer.MustRun()
}
