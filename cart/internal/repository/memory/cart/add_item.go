package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

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
