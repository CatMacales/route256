package cart

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/google/uuid"
)

func (s *Service) Checkout(ctx context.Context, userID model.UserID) (uuid.UUID, error) {
	items, err := s.cartRepository.GetCart(ctx, userID)
	if err != nil {
		return [16]byte{}, fmt.Errorf("failed to get cart: %w", err)
	}

	orderID, err := s.lomsService.CreateOrder(ctx, userID, items)
	if err != nil {
		return [16]byte{}, fmt.Errorf("failed to create order: %w", err)
	}

	err = s.cartRepository.DeleteCart(ctx, userID)
	if err != nil {
		return [16]byte{}, fmt.Errorf("failed to delete cart: %w", err)
	}

	return orderID, nil
}
