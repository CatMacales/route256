package stock_repository

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/repository"
)

func (r *Repository) GetBySKU(_ context.Context, sku model.Sku) (*model.Stock, error) {
	if stock, ok := r.storage[sku]; ok {
		return &stock, nil
	}

	return nil, repository.ErrStockNotFound
}
