package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

// DeleteCart removes the cart by userID.
func (r *Repository) DeleteCart(_ context.Context, userID model.UserID) error {
	r.storage[userID] = []model.Item{}
	return nil
}
