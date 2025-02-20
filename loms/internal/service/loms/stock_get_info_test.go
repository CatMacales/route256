package loms

import (
	"context"
	"github.com/CatMacales/route256/loms/internal/domain/model"
	"github.com/CatMacales/route256/loms/internal/repository"
	"github.com/CatMacales/route256/loms/internal/service/loms/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetStockInfo(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)

	type fields struct {
		stockProviderMock *mock.StockProviderMock
	}

	type args struct {
		sku model.Sku
	}

	tests := []struct {
		name      string
		setup     func(*fields, *args)
		args      args
		wantErr   error
		wantCount uint64
	}{
		{
			name: "success get stock info",
			setup: func(fields *fields, args *args) {
				fields.stockProviderMock.GetBySKUMock.Expect(ctx, args.sku).Return(&model.Stock{
					Quantity: 150,
					Reserved: 50,
				}, nil)
			},
			args: args{
				sku: 12,
			},
			wantErr:   nil,
			wantCount: 100,
		},
		{
			name: "stock not found",
			setup: func(fields *fields, args *args) {
				fields.stockProviderMock.GetBySKUMock.Expect(ctx, args.sku).Return(nil, repository.ErrStockNotFound)
			},
			args: args{
				sku: 12,
			},
			wantErr:   repository.ErrStockNotFound,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				stockProviderMock: mock.NewStockProviderMock(ctrl),
			}
			tt.setup(&f, &tt.args)

			lomsService := NewService(nil, f.stockProviderMock)

			count, err := lomsService.GetStockInfo(ctx, tt.args.sku)

			require.ErrorIs(t, err, tt.wantErr)
			require.Equal(t, tt.wantCount, count)
		})
	}
}
