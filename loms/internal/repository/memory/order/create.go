package order_repository

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/google/uuid"
)

func (r *Repository) Create(_ context.Context, order model.Order) (model.OrderID, error) {
	orderID := uuid.New()
	r.storage[orderID] = order
	return orderID, nil
}
