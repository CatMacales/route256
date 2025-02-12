package cart

import (
	"context"
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
		productMock  *mock.ProductServiceMock
		cartRepoMock *mock.CartRepositoryMock
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
			name: "Success add item",
			setup: func(fields *fields, args *args) {
				fields.productMock.GetProductMock.Return(&model.Product{}, nil)
				fields.cartRepoMock.AddItemMock.Expect(ctx, args.userID, args.item).Return(nil)
			},
			args: args{
				userID: 123,
				item:   model.Item{SKU: 123, Count: 1},
			},
			wantErr: nil,
		},
		{
			name: "Product not found",
			setup: func(fields *fields, args *args) {
				fields.productMock.GetProductMock.Return(&model.Product{}, product_app.ErrProductNotFound)
			},
			args: args{
				userID: 123,
				item:   model.Item{SKU: 123, Count: 1},
			},
			wantErr: service.ErrProductNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				productMock:  mock.NewProductServiceMock(ctrl),
				cartRepoMock: mock.NewCartRepositoryMock(ctrl),
			}
			tt.setup(&f, &tt.args)

			cartService := NewService(f.cartRepoMock, f.productMock)

			err := cartService.AddItem(ctx, tt.args.userID, tt.args.item)

			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
