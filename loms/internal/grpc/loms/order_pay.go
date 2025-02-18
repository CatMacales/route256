package loms_grpc

import (
	"context"
	"errors"
	"github.com/CatMacales/route256/loms/internal/repository"
	"github.com/CatMacales/route256/loms/internal/service"
	"github.com/CatMacales/route256/loms/pkg/api/loms/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) PayOrder(ctx context.Context, request *loms.PayOrderRequest) (*loms.PayOrderResponse, error) {
	err := s.lomsService.PayOrder(ctx, uuid.MustParse(request.GetOrderId()))
	if err != nil {
		if errors.Is(err, repository.ErrOrderNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		if errors.Is(err, service.ErrBadStatus) {
			return nil, status.Errorf(codes.FailedPrecondition, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &loms.PayOrderResponse{}, nil
}
