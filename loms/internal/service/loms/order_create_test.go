package loms

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/repository"
	"github.com/CatMacales/route256/loms/internal/service/loms/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)

	type fields struct {
		orderProviderMock *mock.OrderProviderMock
		stockProviderMock *mock.StockProviderMock
	}

	type args struct {
		order model.Order
	}

	tests := []struct {
		name    string
		setup   func(*fields, *args)
		args    args
		wantErr error
	}{
		{
			name: "success create order",
			setup: func(fields *fields, args *args) {
				orderID := uuid.New()
				fields.orderProviderMock.
					CreateMock.Expect(ctx, args.order).Return(orderID, nil).
					SetStatusMock.When(ctx, orderID, model.StatusAwaitingPayment).Then(nil)
				fields.stockProviderMock.ReserveMock.Expect(ctx, args.order.Items).Return(nil)
			},
			args: args{
				order: model.Order{UserID: 1, Items: []model.Item{{SKU: 123, Count: 12}}, Status: model.StatusNew},
			},
			wantErr: nil,
		},
		{
			name: "not enough stock",
			setup: func(fields *fields, args *args) {
				orderID := uuid.New()
				fields.orderProviderMock.
					CreateMock.Expect(ctx, args.order).Return(orderID, nil).
					SetStatusMock.When(ctx, orderID, model.StatusFailed).Then(nil)
				fields.stockProviderMock.ReserveMock.Expect(ctx, args.order.Items).Return(repository.ErrNotEnoughStock)
			},
			args: args{
				order: model.Order{UserID: 1, Items: []model.Item{{SKU: 123, Count: 12}}, Status: model.StatusNew},
			},
			wantErr: repository.ErrNotEnoughStock,
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

			orderID, err := lomsService.CreateOrder(ctx, tt.args.order)

			require.ErrorIs(t, err, tt.wantErr)
			require.NoError(t, uuid.Validate(orderID.String()))
		})
	}
}
