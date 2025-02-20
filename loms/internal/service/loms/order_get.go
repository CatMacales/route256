package loms

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/loms/internal/domain/model"
)

func (s *Service) GetOrder(ctx context.Context, orderID model.OrderID) (*model.Order, error) {
	order, err := s.orderProvider.GetByOrderID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order by ID: %w", err)
	}

	return order, nil
}
