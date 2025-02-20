package order_repository

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/repository"
)

func (r *Repository) SetStatus(_ context.Context, orderID model.OrderID, status model.OrderStatus) error {
	if order, ok := r.storage[orderID]; ok {
		order.Status = status
		r.storage[orderID] = order
		return nil
	}

	return repository.ErrOrderNotFound
}
