package integration

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/pkg/api/loms/v1"
	"github.com/stretchr/testify/require"
)

func (ls *LOMSSuite) TestGetOrder() {
	ctx := context.Background()

	err := ls.stockRepo.Add(ctx, 773297411, 1002, 12)
	require.NoError(ls.T(), err)
	err = ls.stockRepo.Add(ctx, 1002, 1002, 12)
	require.NoError(ls.T(), err)

	userID := model.UserID(123)

	item1 := model.Item{SKU: 773297411, Count: 2}
	item2 := model.Item{SKU: 1002, Count: 1}

	orderID, err := ls.orderRepo.Create(ctx, model.Order{
		UserID: userID,
		Items:  []model.Item{item1, item2},
		Status: model.StatusAwaitingPayment,
	})
	require.NoError(ls.T(), err)

	resp, err := ls.client.GetOrderInfo(ctx,
		&loms.GetOrderInfoRequest{
			OrderId: orderID.String(),
		})

	require.NoError(ls.T(), err)
	require.Equal(ls.T(), loms.OrderStatus_ORDER_STATUS_AWAITING_PAYMENT, resp.Status)
	wantItems := make([]model.Item, 0, len(resp.Items))
	for _, item := range resp.Items {
		wantItems = append(wantItems, model.Item{SKU: item.GetSku(), Count: uint16(item.GetCount())})
	}
	require.ElementsMatch(ls.T(), wantItems, []model.Item{item1, item2})
	require.Equal(ls.T(), userID, resp.UserId)
}
