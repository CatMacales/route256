//go:build integration

package integration

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/pkg/api/loms/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func (ls *LOMSSuite) TestCreateOrder() {
	ctx := context.Background()

	err := ls.stockRepo.Add(ctx, 773297411, 1002, 12)
	require.NoError(ls.T(), err)
	err = ls.stockRepo.Add(ctx, 1002, 1002, 12)
	require.NoError(ls.T(), err)

	userID := model.UserID(123)

	item1 := model.Item{SKU: 773297411, Count: 2}
	item2 := model.Item{SKU: 1002, Count: 1}

	resp, err := ls.client.CreateOrder(ctx,
		&loms.CreateOrderRequest{
			UserId: userID,
			Items:  model.ItemsToProto([]model.Item{item1, item2}),
		})

	require.NoError(ls.T(), err)
	require.NoError(ls.T(), uuid.Validate(resp.OrderId))

	order, err := ls.orderRepo.GetByOrderID(ctx, uuid.MustParse(resp.OrderId))

	require.NoError(ls.T(), err)
	require.Equal(ls.T(), model.StatusAwaitingPayment, order.Status)
	require.ElementsMatch(ls.T(), []model.Item{item1, item2}, order.Items)
	require.Equal(ls.T(), userID, order.UserID)
}
