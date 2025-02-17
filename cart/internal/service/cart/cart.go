package cart

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/http-server/handler/cart"
	"github.com/google/uuid"
)

var _ cart_handler.CartService = (*Service)(nil)

type CartProvider interface {
	AddItem(context.Context, model.UserID, model.Item) error
	DeleteItem(context.Context, model.UserID, model.Sku) error
	DeleteCart(context.Context, model.UserID) error
	GetCart(context.Context, model.UserID) ([]model.Item, error)
}

type ProductService interface {
	GetProduct(context.Context, uint32) (*model.Product, error)
}

type LOMSService interface {
	CreateOrder(context.Context, model.UserID, []model.Item) (uuid.UUID, error)
}

type Service struct {
	cartRepository CartProvider
	productService ProductService
	lomsService    LOMSService
}

func NewService(cartRepository CartProvider, productService ProductService, lomsService LOMSService) *Service {
	return &Service{
		cartRepository: cartRepository,
		productService: productService,
		lomsService:    lomsService,
	}
}
