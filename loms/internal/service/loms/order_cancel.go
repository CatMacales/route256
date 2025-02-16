package loms

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/service"
)

func (s *Service) CancelOrder(ctx context.Context, orderID model.OrderID) error {
	order, err := s.orderProvider.GetByOrderID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to get order by ID: %w", err)
	}

	if order.Status != model.StatusAwaitingPayment {
		return fmt.Errorf("%w, need %s, get %s", service.ErrBadStatus, model.StatusAwaitingPayment.String(), order.Status.String())
	}

	err = s.stockProvider.ReserveCancel(ctx, order.Items)
	if err != nil {
		return fmt.Errorf("failed to cancel reserved stock: %w", err)
	}

	err = s.orderProvider.SetStatus(ctx, orderID, model.StatusCancelled)
	if err != nil {
		return fmt.Errorf("failed to set order status to %s: %w", model.StatusCancelled.String(), err)
	}

	return nil
}
