package cart

import (
	"context"
	product_app "github.com/CatMacales/route256/cart/internal/app/product"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/repository"
	"github.com/CatMacales/route256/cart/internal/service"
	"github.com/CatMacales/route256/cart/internal/service/cart/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetCart(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)

	type fields struct {
		productMock      *mock.ProductServiceMock
		cartProviderMock *mock.CartProviderMock
	}

	type args struct {
		userID model.UserID
	}

	tests := []struct {
		name     string
		setup    func(*fields, *args)
		args     args
		wantErr  error
		wantCart *model.Cart
	}{
		{
			name: "success get cart",
			setup: func(fields *fields, args *args) {
				fields.productMock.GetProductMock.
					Expect(ctx, uint32(100)).
					Return(&model.Product{Name: "Product1", Price: 1000}, nil)
				fields.cartProviderMock.GetCartMock.
					Expect(ctx, args.userID).
					Return([]model.Item{{SKU: 100, Count: 2}}, nil)
			},
			args: args{
				userID: 1,
			},
			wantErr: nil,
			wantCart: &model.Cart{
				Items:      []model.CartItem{{Item: model.Item{SKU: 100, Count: 2}, Product: model.Product{Name: "Product1", Price: 1000}}},
				TotalPrice: 2000,
			},
		},
		{
			name: "user not found",
			setup: func(fields *fields, args *args) {
				fields.cartProviderMock.GetCartMock.Return(nil, repository.ErrUserNotFound)
			},
			args: args{
				userID: 1,
			},
			wantErr:  service.ErrEmptyCart,
			wantCart: nil,
		},
		{
			name: "zero cart len",
			setup: func(fields *fields, args *args) {
				fields.cartProviderMock.GetCartMock.Return(make([]model.Item, 0), nil)
			},
			args: args{
				userID: 1,
			},
			wantErr:  service.ErrEmptyCart,
			wantCart: nil,
		},
		{
			name: "product not found",
			setup: func(fields *fields, args *args) {
				fields.productMock.GetProductMock.Return(nil, product_app.ErrProductNotFound)
				fields.cartProviderMock.GetCartMock.Return([]model.Item{{SKU: 100, Count: 2}}, nil)
			},
			args: args{
				userID: 1,
			},
			wantErr:  service.ErrProductNotFound,
			wantCart: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				productMock:      mock.NewProductServiceMock(ctrl),
				cartProviderMock: mock.NewCartProviderMock(ctrl),
			}
			tt.setup(&f, &tt.args)

			cartService := NewService(f.cartProviderMock, f.productMock, nil)

			cart, err := cartService.GetCart(ctx, tt.args.userID)

			assert.Equal(t, tt.wantCart, cart)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
