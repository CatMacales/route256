package main

import (
	"github.com/CatMacales/route256/cart/internal/app"
	"github.com/CatMacales/route256/cart/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	application := app.New(cfg.Host, cfg.Port, cfg.ProductService.URL, cfg.ProductService.Token)

	err := application.Server.Serve()
	if err != nil {
		panic(err)
	}
}
