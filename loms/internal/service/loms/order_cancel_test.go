package loms

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/service"
	"github.com/CatMacales/route256/loms/internal/service/loms/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCancelOrder(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)

	type fields struct {
		orderProviderMock *mock.OrderProviderMock
		stockProviderMock *mock.StockProviderMock
	}

	type args struct {
		orderID model.OrderID
	}

	tests := []struct {
		name    string
		setup   func(*fields, *args)
		args    args
		wantErr error
	}{
		{
			name: "success cancel order",
			setup: func(fields *fields, args *args) {
				order := &model.Order{UserID: 1, Items: []model.Item{{SKU: 123, Count: 12}}, Status: model.StatusAwaitingPayment}
				fields.orderProviderMock.
					GetByOrderIDMock.Expect(ctx, args.orderID).Return(order, nil).
					SetStatusMock.Expect(ctx, args.orderID, model.StatusCancelled).Return(nil)
				fields.stockProviderMock.ReserveCancelMock.Expect(ctx, order.Items).Return(nil)
			},
			args: args{
				orderID: uuid.New(),
			},
			wantErr: nil,
		},
		{
			name: "bad order status",
			setup: func(fields *fields, args *args) {
				order := &model.Order{UserID: 1, Items: []model.Item{{SKU: 123, Count: 12}}, Status: model.StatusNew}
				fields.orderProviderMock.GetByOrderIDMock.Expect(ctx, args.orderID).Return(order, nil)
			},
			args: args{
				orderID: uuid.New(),
			},
			wantErr: service.ErrBadStatus,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				orderProviderMock: mock.NewOrderProviderMock(ctrl),
				stockProviderMock: mock.NewStockProviderMock(ctrl),
			}
			tt.setup(&f, &tt.args)

			lomsService := NewService(f.orderProviderMock, f.stockProviderMock)

			err := lomsService.CancelOrder(ctx, tt.args.orderID)

			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
