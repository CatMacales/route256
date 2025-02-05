package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

// DeleteItem removes an item by SKU from the user's cart.
// If the item is found and removed or not found, it returns nil.
func (r *Repository) DeleteItem(_ context.Context, userID model.UserID, sku model.Sku) error {
	for i, item := range r.storage[userID] {
		if item.SKU == sku {
			r.storage[userID] = append(r.storage[userID][:i], r.storage[userID][i+1:]...)
			return nil
		}
	}

	return nil
}
