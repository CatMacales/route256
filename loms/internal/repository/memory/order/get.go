package order_repository

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/repository"
)

func (r *Repository) GetByOrderID(_ context.Context, orderID model.OrderID) (*model.Order, error) {
	if order, ok := r.storage[orderID]; ok {
		return &order, nil
	}

	return nil, repository.ErrOrderNotFound
}
