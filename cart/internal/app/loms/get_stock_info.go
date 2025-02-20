package loms_app

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/pkg/api/loms/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *App) GetStockInfo(ctx context.Context, sku model.Sku) (uint64, error) {

	req := &loms.GetStockInfoRequest{Sku: uint32(sku)}

	response, err := a.client.GetStockInfo(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.NotFound {
			return 0, fmt.Errorf("%w: %s", ErrStockNotFound, st.Message())
		}
		return 0, fmt.Errorf("failed to get stock info: %w", err)
	}

	return response.GetCount(), nil
}
