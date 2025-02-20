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

func TestGetOrder(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)

	type fields struct {
		orderProviderMock *mock.OrderProviderMock
	}

	type args struct {
		orderID model.OrderID
	}

	tests := []struct {
		name      string
		setup     func(*fields, *args)
		args      args
		wantErr   error
		wantOrder *model.Order
	}{
		{
			name: "success get order",
			setup: func(fields *fields, args *args) {
				order := &model.Order{UserID: 1, Items: []model.Item{{SKU: 123, Count: 12}}, Status: model.StatusAwaitingPayment}
				fields.orderProviderMock.GetByOrderIDMock.Expect(ctx, args.orderID).Return(order, nil)
			},
			args: args{
				orderID: uuid.New(),
			},
			wantErr:   nil,
			wantOrder: &model.Order{UserID: 1, Items: []model.Item{{SKU: 123, Count: 12}}, Status: model.StatusAwaitingPayment},
		},
		{
			name: "order not found",
			setup: func(fields *fields, args *args) {
				fields.orderProviderMock.GetByOrderIDMock.Expect(ctx, args.orderID).Return(nil, repository.ErrOrderNotFound)
			},
			args: args{
				orderID: uuid.New(),
			},
			wantErr:   repository.ErrOrderNotFound,
			wantOrder: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				orderProviderMock: mock.NewOrderProviderMock(ctrl),
			}
			tt.setup(&f, &tt.args)

			lomsService := NewService(f.orderProviderMock, nil)

			order, err := lomsService.GetOrder(ctx, tt.args.orderID)

			require.ErrorIs(t, err, tt.wantErr)
			require.Equal(t, tt.wantOrder, order)
		})
	}
}
