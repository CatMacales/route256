package cart_repository

import (
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/service/cart"
)

var _ cart.CartRepository = (*Repository)(nil)

type Storage = map[model.UserID][]model.Item

type Repository struct {
	storage Storage
}

func NewRepository() *Repository {
	return &Repository{
		storage: make(Storage),
	}
}
