package cart

import (
	"context"
	"errors"
	"fmt"
	"github.com/CatMacales/route256/cart/internal/app/product"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/repository"
	"github.com/CatMacales/route256/cart/internal/service"
	"sort"
)

// GetCart return the cart by userID. Items sorts them by SKU.
// If the user is not found or the cart is empty, it returns an ErrEmptyCart error.
// If a product is not found, it returns an ErrProductNotFound error.
func (s *Service) GetCart(ctx context.Context, userID model.UserID) (*model.Cart, error) {
	items, err := s.cartRepository.GetCart(ctx, userID)
	if err != nil || len(items) == 0 {
		if errors.Is(err, repository.ErrUserNotFound) || len(items) == 0 {
			return nil, service.ErrEmptyCart
		}
		return nil, err
	}

	sort.Slice(items, func(i, j int) bool {
		return (items)[i].SKU < (items)[j].SKU
	})

	var cart model.Cart

	for _, item := range items {
		product, err := s.productService.GetProduct(ctx, uint32(item.SKU))
		if err != nil {
			if errors.Is(err, product_app.ErrProductNotFound) {
				return nil, fmt.Errorf("SKU: %d, err: %w", item.SKU, service.ErrProductNotFound)
			}

			return nil, err
		}

		cart.Items = append(cart.Items, model.CartItem{Item: item, Product: *product})
		cart.TotalPrice += uint32(item.Count) * product.Price
	}

	return &cart, nil
}
