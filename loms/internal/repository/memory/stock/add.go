package stock_repository

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
)

func (r *Repository) Add(_ context.Context, sku model.Sku, quantity uint64, reserved uint64) error {
	stock := r.storage[sku]
	stock.Quantity += quantity
	stock.Reserved += reserved
	r.storage[sku] = stock
	return nil
}
