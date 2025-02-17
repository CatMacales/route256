package loms_app

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/pkg/api/loms/v1"
)

func (a *App) CreateOrder(ctx context.Context, userID model.UserID, items []model.Item) ([16]byte, error) {

	protoItems := make([]*loms.Item, 0, len(items))

	for _, item := range items {
		protoItems = append(protoItems, &loms.Item{
			Sku:   uint32(item.SKU),
			Count: uint32(item.Count),
		})
	}

	req := &loms.CreateOrderRequest{
		UserId: userID,
		Items:  protoItems,
	}

	response, err := a.client.CreateOrder(ctx, req)
	if err != nil {
		return [16]byte{}, fmt.Errorf("failed to create order: %w", err)
	}

	return [16]byte(response.OrderId), nil
}
