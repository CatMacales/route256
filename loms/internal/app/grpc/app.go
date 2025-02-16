package grpc_app

import (
	"fmt"
	"github.com/CatMacales/route256/loms/internal/grpc/loms"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	host       string
	port       uint32
}

// New creates new gRPC server app.
func New(lomsService loms_grpc.LOMSService, host string, port uint32) *App {
	grpcServer := grpc.NewServer()

	loms_grpc.RegisterServer(grpcServer, lomsService)

	return &App{
		gRPCServer: grpcServer,
		host:       host,
		port:       port,
	}
}

// MustRun runs gRPC server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run runs gRPC server.
func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.host, a.port))
	if err != nil {
		return err
	}

	log.Printf("Starting grpc server on %s\n", l.Addr().String())

	if err = a.gRPCServer.Serve(l); err != nil {
		return err
	}

	return nil
}

// Stop stops gRPC server.
func (a *App) Stop() {
	log.Printf("Stopping grpc server on %s:%d\n", a.host, a.port)

	a.gRPCServer.GracefulStop()
}
