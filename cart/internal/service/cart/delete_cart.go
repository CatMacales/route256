package cart

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

func (s *Service) DeleteCart(ctx context.Context, userID model.UserID) error {
	return s.cartRepository.DeleteCart(ctx, userID)
}
