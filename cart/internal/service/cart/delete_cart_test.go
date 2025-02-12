package cart

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/service/cart/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteCart(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)

	type fields struct {
		cartRepoMock *mock.CartRepositoryMock
		productMock  *mock.ProductServiceMock
	}

	type args struct {
		userID model.UserID
	}

	tests := []struct {
		name    string
		setup   func(*fields, *args)
		args    args
		wantErr error
	}{
		{
			name: "Success delete cart",
			setup: func(fields *fields, args *args) {
				fields.cartRepoMock.DeleteCartMock.Expect(ctx, args.userID).Return(nil)
			},
			args: args{
				userID: 123,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				productMock:  nil,
				cartRepoMock: mock.NewCartRepositoryMock(ctrl),
			}
			tt.setup(&f, &tt.args)

			cartService := NewService(f.cartRepoMock, f.productMock)

			err := cartService.DeleteCart(ctx, tt.args.userID)

			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
