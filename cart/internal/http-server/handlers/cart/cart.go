package cart_http

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
)

type CartService interface {
	AddItem(ctx context.Context, userID model.UserID, item model.Item) error
	DeleteItem(ctx context.Context, userID model.UserID, sku model.Sku) error
	DeleteCart(ctx context.Context, userID model.UserID) error
	GetCart(ctx context.Context, userID model.UserID) (*model.Cart, error)
}

type Server struct {
	cartService CartService
}

func New(cartService CartService) *Server {
	return &Server{cartService: cartService}
}
