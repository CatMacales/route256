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
		cartProviderMock *mock.CartProviderMock
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
			name: "success delete cart",
			setup: func(fields *fields, args *args) {
				fields.cartProviderMock.DeleteCartMock.Expect(ctx, args.userID).Return(nil)
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
				cartProviderMock: mock.NewCartProviderMock(ctrl),
			}
			tt.setup(&f, &tt.args)

			cartService := NewService(f.cartProviderMock, nil, nil)

			err := cartService.DeleteCart(ctx, tt.args.userID)

			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
