package app

import (
	"context"
	"encoding/json"
	"github.com/CatMacales/route256/loms/internal/app/grpc"
	"github.com/CatMacales/route256/loms/internal/repository/memory/order"
	"github.com/CatMacales/route256/loms/internal/repository/memory/stock"
	"github.com/CatMacales/route256/loms/internal/service/loms"
	"os"
)

type App struct {
	GRPCServer      *grpc_app.App
	stockRepository *stock_repository.Repository
}

type stockData struct {
	SKU        uint32 `json:"sku"`
	TotalCount uint64 `json:"total_count"`
	Reserved   uint64 `json:"reserved"`
}

func New(host string, port uint32) *App {
	orderRepository := order_repository.NewRepository()
	stockRepository := stock_repository.NewRepository()

	lomsService := loms.NewService(orderRepository, stockRepository)

	grpcApp := grpc_app.New(lomsService, host, port)

	return &App{
		GRPCServer:      grpcApp,
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

	var stockData []stockData
	err = json.Unmarshal(data, &stockData)
	if err != nil {
		return err
	}

	for _, stock := range stockData {
		err = a.stockRepository.Add(ctx, stock.SKU, stock.TotalCount, stock.Reserved)
		if err != nil {
			return err
		}
	}

	return nil
}
