package http_gateway_app

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/loms/pkg/api/loms/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type App struct {
	server  *http.Server
	Router  *runtime.ServeMux
	grpcUrl string
}

func New(host string, port uint32, grpcUrl string) *App {
	gwmux := runtime.NewServeMux()

	server := &http.Server{
		Handler: gwmux,
		Addr:    fmt.Sprintf("%s:%d", host, port),
	}

	return &App{
		server:  server,
		Router:  gwmux,
		grpcUrl: grpcUrl,
	}
}

func (a *App) MustConnectToGRPC() {
	if err := a.ConnectToGRPC(); err != nil {
		panic(err)
	}
}

func (a *App) ConnectToGRPC() error {
	conn, err := grpc.NewClient(a.grpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	if err = loms.RegisterLOMSHandler(context.Background(), a.Router, conn); err != nil {
		log.Fatalln("Failed to register http-gateway:", err)
	}

	return nil
}

// Serve starts the server and blocks until it exits.
func (a *App) Serve() error {
	log.Printf("Starting http-gateway on %s\n", a.server.Addr)
	return a.server.ListenAndServe()
}
