package cart

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

func (s *Service) DeleteItem(ctx context.Context, userID model.UserID, sku model.Sku) error {
	return s.cartRepository.DeleteItem(ctx, userID, sku)
}
