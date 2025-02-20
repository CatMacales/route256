package cart

import (
	"context"
	"errors"
	"fmt"
	"github.com/CatMacales/route256/cart/internal/app/loms"
	"github.com/CatMacales/route256/cart/internal/app/product"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/service"
)

// AddItem adds an item to the user's cart.
//
// It adds product to cart only if it exists in the ProductService.
// If the product is not found, it returns a product_app.ErrProductNotFound.
func (s *Service) AddItem(ctx context.Context, userID model.UserID, item model.Item) error {
	_, err := s.productService.GetProduct(ctx, uint32(item.SKU))
	if err != nil {
		if errors.Is(err, product_app.ErrProductNotFound) {
			return service.ErrProductNotFound
		}

		return err
	}

	count, err := s.lomsService.GetStockInfo(ctx, item.SKU)
	if err != nil || uint64(item.Count) > count {
		if errors.Is(err, loms_app.ErrStockNotFound) || uint64(item.Count) > count {
			return service.ErrNotEnoughStock
		}

		return err
	}

	err = s.cartRepository.AddItem(ctx, userID, item)
	if err != nil {
		return fmt.Errorf("failed to add item to cart: %w", err)
	}

	return nil
}
