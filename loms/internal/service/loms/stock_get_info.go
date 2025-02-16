package loms

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/loms/internal/domain/model"
)

func (s *Service) GetStockInfo(ctx context.Context, sku model.Sku) (uint64, error) {
	stock, err := s.stockProvider.GetBySKU(ctx, sku)
	if err != nil {
		return 0, fmt.Errorf("failed to get stock by SKU: %w", err)
	}

	return stock.Quantity - stock.Reserved, nil
}
