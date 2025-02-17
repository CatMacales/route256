package loms_app

import (
	"github.com/CatMacales/route256/cart/pkg/api/loms/v1"
	"google.golang.org/grpc"
)

type App struct {
	url    string
	client loms.LOMSClient
}

func New(url string, conn *grpc.ClientConn) *App {
	client := loms.NewLOMSClient(conn)
	return &App{
		url:    url,
		client: client,
	}
}
