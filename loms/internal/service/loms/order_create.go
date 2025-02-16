package loms

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/loms/internal/domain/model"
)

func (s *Service) CreateOrder(ctx context.Context, order model.Order) (model.OrderID, error) {
	order.Status = model.StatusNew

	orderID, err := s.orderProvider.Create(ctx, order)
	if err != nil {
		return model.OrderID{}, fmt.Errorf("failed to create order: %w", err)
	}

	err = s.stockProvider.Reserve(ctx, order.Items)
	if err != nil {
		err = s.orderProvider.SetStatus(ctx, orderID, model.StatusFailed)
		if err != nil {
			return model.OrderID{}, fmt.Errorf("failed to set order status to %s: %w", model.StatusFailed.String(), err)
		}
		return model.OrderID{}, fmt.Errorf("failed to reserve stocks: %w", err)
	}

	err = s.orderProvider.SetStatus(ctx, orderID, model.StatusAwaitingPayment)
	if err != nil {
		return model.OrderID{}, fmt.Errorf("failed to set order status to %s: %w", model.StatusAwaitingPayment.String(), err)
	}

	return orderID, nil
}
