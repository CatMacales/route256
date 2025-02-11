package server

import (
	"fmt"
	"github.com/CatMacales/route256/cart/internal/http-server/handler/cart"
	"github.com/CatMacales/route256/cart/internal/http-server/middleware"
	"log"
	"net/http"
)

type Server struct {
	server *http.Server
	Router *http.Handler
}

func New(host string, port uint32, h *cart_handler.Handler) *Server {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", h.AddItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", h.DeleteItem)
	mux.HandleFunc("DELETE /user/{user_id}/cart", h.DeleteCart)
	mux.HandleFunc("GET /user/{user_id}/cart", h.GetCart)

	logMux := middleware.NewLogger(mux)

	server := &http.Server{
		Handler: logMux,
		Addr:    fmt.Sprintf("%s:%d", host, port),
	}

	return &Server{
		server: server,
		Router: &logMux,
	}
}

// Serve starts the server and blocks until it exits.
func (s *Server) Serve() error {
	log.Printf("Starting server on %s\n", s.server.Addr)
	return s.server.ListenAndServe()
}
