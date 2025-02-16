package stock_repository

import (
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/service/loms"
)

var _ loms.StockProvider = (*Repository)(nil)

type Storage = map[model.Sku]model.Stock

type Repository struct {
	storage Storage
}

func NewRepository() *Repository {
	return &Repository{
		storage: make(Storage),
	}
}
