package cart

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

// DeleteCart removes all items from the cart by userID.
func (s *Service) DeleteCart(ctx context.Context, userID model.UserID) error {
	err := s.cartRepository.DeleteCart(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete cart: %w", err)
	}

	return nil
}
