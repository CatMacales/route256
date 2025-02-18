package cart

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/service/cart/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteItem(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)

	type fields struct {
		cartRepoMock *mock.CartProviderMock
	}

	type args struct {
		userID model.UserID
		sku    model.Sku
	}

	tests := []struct {
		name    string
		setup   func(*fields, *args)
		args    args
		wantErr error
	}{
		{
			name: "success delete item",
			setup: func(fields *fields, args *args) {
				fields.cartRepoMock.DeleteItemMock.Expect(ctx, args.userID, args.sku).Return(nil)
			},
			args: args{
				userID: 123,
				sku:    123,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				cartRepoMock: mock.NewCartProviderMock(ctrl),
			}
			tt.setup(&f, &tt.args)

			cartService := NewService(f.cartRepoMock, nil, nil)

			err := cartService.DeleteItem(ctx, tt.args.userID, tt.args.sku)

			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
