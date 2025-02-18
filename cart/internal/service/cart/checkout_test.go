package cart

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/repository"
	"github.com/CatMacales/route256/cart/internal/service/cart/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckout(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)

	type fields struct {
		lomsMock         *mock.LOMSServiceMock
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
			name: "success checkout",
			setup: func(fields *fields, args *args) {
				fields.cartProviderMock.GetCartMock.Expect(ctx, args.userID).Return([]model.Item{{SKU: 123, Count: 1}}, nil)
				fields.cartProviderMock.DeleteCartMock.Expect(ctx, args.userID).Return(nil)
				fields.lomsMock.CreateOrderMock.
					Expect(ctx, args.userID, []model.Item{{SKU: 123, Count: 1}}).
					Return(uuid.New(), nil)
			},
			args: args{
				userID: 123,
			},
			wantErr: nil,
		},
		{
			name: "user not found",
			setup: func(fields *fields, args *args) {
				fields.cartProviderMock.GetCartMock.Expect(ctx, args.userID).Return([]model.Item{}, repository.ErrUserNotFound)
			},
			args: args{
				userID: 123,
			},
			wantErr: repository.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				lomsMock:         mock.NewLOMSServiceMock(ctrl),
				cartProviderMock: mock.NewCartProviderMock(ctrl),
			}
			tt.setup(&f, &tt.args)

			cartService := NewService(f.cartProviderMock, nil, f.lomsMock)

			orderID, err := cartService.Checkout(ctx, tt.args.userID)

			assert.NoError(t, uuid.Validate(orderID.String()))
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
