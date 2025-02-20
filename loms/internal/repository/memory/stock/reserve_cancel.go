package stock_repository

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
)

func (r *Repository) ReserveCancel(_ context.Context, items []model.Item) error {
	for _, item := range items {
		stock := r.storage[item.SKU]

		stock.Reserved -= uint64(item.Count)
		r.storage[item.SKU] = stock
	}

	return nil
}
