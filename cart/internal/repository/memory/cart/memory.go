package cart_repository

import "github.com/CatMacales/route256/cart/internal/domain/model"

type Storage = map[model.UserID][]model.Item

type Repository struct {
	storage Storage
}

func NewRepository() *Repository {
	return &Repository{
		storage: make(Storage),
	}
}
