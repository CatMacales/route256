package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

func (r *Repository) DeleteCart(_ context.Context, userID model.UserID) error {
	r.storage[userID] = []model.Item{}
	return nil
}
