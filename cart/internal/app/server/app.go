package server

import (
	"fmt"
	"github.com/CatMacales/route256/cart/internal/http-server/handler/cart"
	"github.com/CatMacales/route256/cart/internal/http-server/middleware"
	"log"
	"net/http"
)

type App struct {
	server *http.Server
	Router http.Handler
}

func New(host string, port uint32, h *cart_handler.Handler) *App {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", h.AddItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", h.DeleteItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart", h.DeleteCart)
	mux.HandleFunc("GET /user/{user_id}/cart", h.GetCart)
	mux.HandleFunc("POST /cart/checkout", h.Checkout)

	logMux := middleware.NewLogger(mux)

	server := &http.Server{
		Handler: logMux,
		Addr:    fmt.Sprintf("%s:%d", host, port),
	}

	return &App{
		server: server,
		Router: logMux,
	}
}

// Serve starts the server and blocks until it exits.
func (a *App) Serve() error {
	log.Printf("Starting server on %s\n", a.server.Addr)
	return a.server.ListenAndServe()
}
