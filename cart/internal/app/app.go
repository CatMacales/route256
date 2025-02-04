package app

import (
	"fmt"
	"github.com/CatMacales/route256/cart/internal/app/product"
	"github.com/CatMacales/route256/cart/internal/http-server/handlers/cart"
	"github.com/CatMacales/route256/cart/internal/http-server/middleware"
	"github.com/CatMacales/route256/cart/internal/lib/validation"
	"github.com/CatMacales/route256/cart/internal/repository/memory/cart"
	"github.com/CatMacales/route256/cart/internal/service/cart"
	"log"
	"net/http"
)

type App struct {
	ProductService *product_app.App
	mux            http.Handler
	host           string
	port           uint32
}

func New(host string, port uint32, productURL, productToken string) *App {
	productApp := product_app.New(productURL, productToken, &http.Client{Transport: middleware.NewRetry(http.DefaultTransport)})

	cartRepository := cart_repository.NewRepository()

	cartService := cart.NewService(cartRepository, productApp)

	server := cart_http.New(cartService)

	validation.NewValidator() // init beauty validator

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", server.AddItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", server.DeleteItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart", server.DeleteCart)
	mux.HandleFunc("GET /user/{user_id}/cart", server.GetCart)

	logMux := middleware.NewLogger(mux)

	return &App{
		ProductService: productApp,
		mux:            logMux,
		host:           host,
		port:           port,
	}
}

// MustRun starts the server and blocks until it exits. Panic if there is an error.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run starts the server and blocks until it exits. Panic if there is an error.
func (a *App) Run() error {
	log.Printf("Starting server on %s:%d\n", a.host, a.port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", a.host, a.port), a.mux)
}
