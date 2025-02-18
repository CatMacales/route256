package cart

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/app/loms"
	"github.com/CatMacales/route256/cart/internal/app/product"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/service"
	"github.com/CatMacales/route256/cart/internal/service/cart/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddItem(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)

	type fields struct {
		productMock      *mock.ProductServiceMock
		lomsMock         *mock.LOMSServiceMock
		cartProviderMock *mock.CartProviderMock
	}

	type args struct {
		userID model.UserID
		item   model.Item
	}

	tests := []struct {
		name    string
		setup   func(*fields, *args)
		args    args
		wantErr error
	}{
		{
			name: "success add item",
			setup: func(fields *fields, args *args) {
				fields.productMock.GetProductMock.Expect(ctx, uint32(args.item.SKU)).Return(&model.Product{}, nil)
				fields.cartProviderMock.AddItemMock.Expect(ctx, args.userID, args.item).Return(nil)
				fields.lomsMock.GetStockInfoMock.Expect(ctx, args.item.SKU).Return(100, nil)
			},
			args: args{
				userID: 123,
				item:   model.Item{SKU: 123, Count: 1},
			},
			wantErr: nil,
		},
		{
			name: "product not found",
			setup: func(fields *fields, args *args) {
				fields.productMock.GetProductMock.Return(&model.Product{}, product_app.ErrProductNotFound)
			},
			args: args{
				userID: 123,
				item:   model.Item{SKU: 123, Count: 1},
			},
			wantErr: service.ErrProductNotFound,
		},
		{
			name: "stock less than item count",
			setup: func(fields *fields, args *args) {
				fields.productMock.GetProductMock.Expect(ctx, uint32(args.item.SKU)).Return(&model.Product{}, nil)
				fields.lomsMock.GetStockInfoMock.Expect(ctx, args.item.SKU).Return(100, nil)
			},
			args: args{
				userID: 123,
				item:   model.Item{SKU: 123, Count: 111},
			},
			wantErr: service.ErrNotEnoughStock,
		},
		{
			name: "stock not found",
			setup: func(fields *fields, args *args) {
				fields.productMock.GetProductMock.Expect(ctx, uint32(args.item.SKU)).Return(&model.Product{}, nil)
				fields.lomsMock.GetStockInfoMock.Expect(ctx, args.item.SKU).Return(0, loms_app.ErrStockNotFound)
			},
			args: args{
				userID: 123,
				item:   model.Item{SKU: 123, Count: 111},
			},
			wantErr: service.ErrNotEnoughStock,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				productMock:      mock.NewProductServiceMock(ctrl),
				lomsMock:         mock.NewLOMSServiceMock(ctrl),
				cartProviderMock: mock.NewCartProviderMock(ctrl),
			}
			tt.setup(&f, &tt.args)

			cartService := NewService(f.cartProviderMock, f.productMock, f.lomsMock)

			err := cartService.AddItem(ctx, tt.args.userID, tt.args.item)

			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
