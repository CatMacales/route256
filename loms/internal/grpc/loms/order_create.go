package loms_grpc

import (
	"context"
	"errors"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/repository"
	"github.com/CatMacales/route256/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateOrder(ctx context.Context, request *loms.CreateOrderRequest) (*loms.CreateOrderResponse, error) {
	orderID, err := s.lomsService.CreateOrder(ctx, model.ProtoToOrder(request))
	if err != nil {
		if errors.Is(err, repository.ErrNotEnoughStock) {
			return nil, status.Errorf(codes.FailedPrecondition, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &loms.CreateOrderResponse{OrderId: orderID[:]}, nil
}
