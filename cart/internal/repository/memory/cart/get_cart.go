package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/repository"
)

// GetCart return the list of items in the cart for a given user ID.
// It returns an error if the user is not found in the storage.
// The returned slice is a copy of the stored items.
func (r *Repository) GetCart(_ context.Context, userID model.UserID) ([]model.Item, error) {
	items, ok := r.storage[userID]
	if !ok {
		return nil, repository.ErrUserNotFound
	}

	itemsCopy := make([]model.Item, len(items))
	copy(itemsCopy, items)
	return itemsCopy, nil
}
