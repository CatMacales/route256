package cart

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

func (s *Service) AddItem(ctx context.Context, userID model.UserID, item model.Item) error {
	_, err := s.productService.GetProduct(ctx, uint32(item.SKU))
	if err != nil {
		return err
	}

	return s.cartRepository.AddItem(ctx, userID, item)
}
