package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/CatMacales/route256/loms/internal/app/grpc"
	"github.com/CatMacales/route256/loms/internal/app/http-gateway"
	"github.com/CatMacales/route256/loms/internal/repository/memory/order"
	"github.com/CatMacales/route256/loms/internal/repository/memory/stock"
	"github.com/CatMacales/route256/loms/internal/service/loms"
	"os"
)

type App struct {
	GRPCServer      *grpc_app.App
	HttpGateway     *http_gateway_app.App
	stockRepository *stock_repository.Repository
}

type stockData struct {
	SKU        uint32 `json:"sku"`
	TotalCount uint64 `json:"total_count"`
	Reserved   uint64 `json:"reserved"`
}

func New(grpcHost, httpHost string, grpcPort, HttpPort uint32) *App {
	orderRepository := order_repository.NewRepository()
	stockRepository := stock_repository.NewRepository()

	lomsService := loms.NewService(orderRepository, stockRepository)

	grpcApp := grpc_app.New(lomsService, grpcHost, grpcPort)
	httpGatewayApp := http_gateway_app.New(httpHost, HttpPort, fmt.Sprintf("%s:%d", grpcHost, grpcPort))

	return &App{
		GRPCServer:      grpcApp,
		HttpGateway:     httpGatewayApp,
		stockRepository: stockRepository,
	}
}

func (a *App) MustInitStocks(ctx context.Context, path string) {
	if err := a.InitStocks(ctx, path); err != nil {
		panic(err)
	}
}

func (a *App) InitStocks(ctx context.Context, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var stocks []stockData
	err = json.Unmarshal(data, &stocks)
	if err != nil {
		return err
	}

	for _, stock := range stocks {
		err = a.stockRepository.Add(ctx, stock.SKU, stock.TotalCount, stock.Reserved)
		if err != nil {
			return err
		}
	}

	return nil
}
