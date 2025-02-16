package stock_repository

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/repository"
)

func (r *Repository) Reserve(ctx context.Context, items []model.Item) error {
	reservedItems := make([]model.Item, 0, len(items))
	for _, item := range items {
		stock := r.storage[item.SKU]
		if uint64(item.Count) > stock.Quantity {
			err := r.ReserveCancel(ctx, reservedItems)
			if err != nil {
				return fmt.Errorf("failed to cancel reserve stocks %w", err)
			}
			return fmt.Errorf("%w, sku: %d", repository.ErrNotEnoughStock, item.SKU)
		}

		stock.Reserved += uint64(item.Count)
		r.storage[item.SKU] = stock
		reservedItems = append(reservedItems, item)
	}

	return nil
}
