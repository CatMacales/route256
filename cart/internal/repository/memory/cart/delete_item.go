package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

func (r *Repository) DeleteItem(_ context.Context, userID model.UserID, sku model.Sku) error {
	for i, item := range r.storage[userID] {
		if item.SKU == sku {
			r.storage[userID] = append(r.storage[userID][:i], r.storage[userID][i+1:]...)
			return nil
		}
	}

	return nil
}
