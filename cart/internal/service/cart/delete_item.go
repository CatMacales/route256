package cart

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

// DeleteItem removes an item by sku from the user's cart by userID.
func (s *Service) DeleteItem(ctx context.Context, userID model.UserID, sku model.Sku) error {
	err := s.cartRepository.DeleteItem(ctx, userID, sku)
	if err != nil {
		return fmt.Errorf("failed to delete item from cart: %w", err)
	}

	return nil
}
