package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

// AddItem adds a new item to the user's cart.
// If the user does not have an existing cart, a new cart is created with the item.
// If the item already exists in the user's cart, its count is incremented by the new item's count.
func (r *Repository) AddItem(_ context.Context, userID model.UserID, newItem model.Item) error {
	_, ok := r.storage[userID]
	if !ok {
		r.storage[userID] = []model.Item{newItem}
		return nil
	}

	for i, item := range r.storage[userID] {
		if item.SKU == newItem.SKU {
			item.Count += newItem.Count
			r.storage[userID][i] = item
			return nil
		}
	}

	r.storage[userID] = append(r.storage[userID], newItem)

	return nil
}
