package cart

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/http-server/handler/cart"
)

var _ cart_handler.CartService = (*Service)(nil)

type CartRepository interface {
	AddItem(context.Context, model.UserID, model.Item) error
	DeleteItem(context.Context, model.UserID, model.Sku) error
	DeleteCart(context.Context, model.UserID) error
	GetCart(context.Context, model.UserID) ([]model.Item, error)
}

type ProductService interface {
	GetProduct(_ context.Context, sku uint32) (*model.Product, error)
}

type Service struct {
	cartRepository CartRepository
	productService ProductService
}

func NewService(cartRepository CartRepository, productService ProductService) *Service {
	return &Service{
		cartRepository: cartRepository,
		productService: productService,
	}
}
