package loms_grpc

import (
	"context"
	"errors"
	"github.com/CatMacales/route256/loms/internal/repository"
	"github.com/CatMacales/route256/loms/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetStockInfo(ctx context.Context, request *loms.GetStockInfoRequest) (*loms.GetStockInfoResponse, error) {
	count, err := s.lomsService.GetStockInfo(ctx, request.GetSku())
	if err != nil {
		if errors.Is(err, repository.ErrStockNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &loms.GetStockInfoResponse{Count: count}, nil
}
