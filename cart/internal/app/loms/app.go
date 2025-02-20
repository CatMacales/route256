package loms_app

import (
	"github.com/CatMacales/route256/cart/pkg/api/loms/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	url    string
	client loms.LOMSClient
}

func New(url string) *App {
	return &App{
		url: url,
	}
}

func (a *App) MustConnect() {
	if err := a.Connect(); err != nil {
		panic(err)
	}
}

func (a *App) Connect() error {
	conn, err := grpc.NewClient(a.url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	a.client = loms.NewLOMSClient(conn)

	return nil
}
