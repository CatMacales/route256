package cart_handler

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type CartService interface {
	AddItem(context.Context, model.UserID, model.Item) error
	DeleteItem(context.Context, model.UserID, model.Sku) error
	DeleteCart(context.Context, model.UserID) error
	GetCart(context.Context, model.UserID) (*model.Cart, error)
	Checkout(context.Context, model.UserID) (uuid.UUID, error)
}

type Handler struct {
	cartService CartService
}

func New(cartService CartService) *Handler {
	return &Handler{cartService: cartService}
}

func parseIntPathValue(r *http.Request, key string) (int64, error) {
	rawValue := r.PathValue(key)
	return strconv.ParseInt(rawValue, 10, 64)
}
