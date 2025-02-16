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

func (s *server) GetOrderInfo(ctx context.Context, request *loms.GetOrderInfoRequest) (*loms.GetOrderInfoResponse, error) {
	order, err := s.lomsService.GetOrder(ctx, model.OrderID(request.GetOrderId()))
	if err != nil {
		if errors.Is(err, repository.ErrOrderNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return model.OrderToProto(order), nil
}
