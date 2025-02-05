package cart_http

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"net/http"
	"strconv"
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

func parseIntPathValue(r *http.Request, key string) (int64, error) {
	rawValue := r.PathValue(key)
	return strconv.ParseInt(rawValue, 10, 64)
}
