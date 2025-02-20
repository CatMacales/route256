package loms_grpc

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/pkg/api/loms/v1"
	"google.golang.org/grpc"
)

var _ loms.LOMSServer = (*server)(nil)

type LOMSService interface {
	CancelOrder(context.Context, model.OrderID) error
	CreateOrder(context.Context, model.Order) (model.OrderID, error)
	GetOrder(context.Context, model.OrderID) (*model.Order, error)
	PayOrder(context.Context, model.OrderID) error
	GetStockInfo(context.Context, model.Sku) (uint64, error)
}

type server struct {
	loms.UnsafeLOMSServer
	lomsService LOMSService
}

func RegisterServer(gRPCServer *grpc.Server, lomsService LOMSService) {
	loms.RegisterLOMSServer(gRPCServer, &server{lomsService: lomsService})
}
