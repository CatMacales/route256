package loms

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/grpc/loms"
)

var _ loms_grpc.LOMSService = (*Service)(nil)

type OrderProvider interface {
	Create(context.Context, model.Order) (model.OrderID, error)
	GetByOrderID(context.Context, model.OrderID) (*model.Order, error)
	SetStatus(context.Context, model.OrderID, model.OrderStatus) error
}

type StockProvider interface {
	GetBySKU(context.Context, model.Sku) (*model.Stock, error)
	Reserve(context.Context, []model.Item) error
	ReserveCancel(context.Context, []model.Item) error
	ReserveRemove(context.Context, []model.Item) error
}

type Service struct {
	orderProvider OrderProvider
	stockProvider StockProvider
}

func NewService(orderProvider OrderProvider, stockProvider StockProvider) *Service {
	return &Service{
		orderProvider: orderProvider,
		stockProvider: stockProvider,
	}
}
