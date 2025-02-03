package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/repository"
)

func (r *Repository) GetCart(_ context.Context, userID model.UserID) (*[]model.Item, error) {
	items, ok := r.storage[userID]
	if !ok {
		return nil, repository.ErrUserNotFound
	}

	return &items, nil
}
