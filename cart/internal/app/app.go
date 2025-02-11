package app

import (
	"github.com/CatMacales/route256/cart/internal/app/product"
	"github.com/CatMacales/route256/cart/internal/app/server"
	"github.com/CatMacales/route256/cart/internal/http-server/handler/cart"
	"github.com/CatMacales/route256/cart/internal/http-server/middleware"
	"github.com/CatMacales/route256/cart/internal/lib/validation"
	"github.com/CatMacales/route256/cart/internal/repository/memory/cart"
	"github.com/CatMacales/route256/cart/internal/service/cart"
	"net/http"
)

type App struct {
	ProductService *product_app.App
	Server         *server.Server
}

func New(host string, port uint32, productURL, productToken string) *App {
	productApp := product_app.New(productURL, productToken, &http.Client{Transport: middleware.NewRetry(http.DefaultTransport)})

	cartRepository := cart_repository.NewRepository()

	cartService := cart.NewService(cartRepository, productApp)

	handler := cart_handler.New(cartService)

	validation.InitValidator() // init beauty validator

	srv := server.New(host, port, handler)

	return &App{
		ProductService: productApp,
		Server:         srv,
	}
}

// MustRun starts the server. Panic if there is an error.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run starts the server.
func (a *App) Run() error {
	return a.Server.Serve()
}
