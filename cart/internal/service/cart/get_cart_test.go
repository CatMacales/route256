package cart

import (
	"context"
	product_app "github.com/CatMacales/route256/cart/internal/app/product"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/repository"
	"github.com/CatMacales/route256/cart/internal/service"
	"github.com/CatMacales/route256/cart/internal/service/cart/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetCart(t *testing.T) {
	ctx := context.Background()
	ctrl := minimock.NewController(t)

	t.Run("SuccessGetCart", func(t *testing.T) {
		userID := model.UserID(1)
		item1 := model.Item{SKU: 100, Count: 2}
		item2 := model.Item{SKU: 200, Count: 1}
		product1 := model.Product{Name: "Product1", Price: 1000}
		product2 := model.Product{Name: "Product2", Price: 500}

		cartRepoMock := mock.NewCartRepositoryMock(ctrl).GetCartMock.Expect(ctx, userID).Return(
			[]model.Item{
				item1,
				item2,
			}, nil)

		productServiceMock := mock.NewProductServiceMock(ctrl).
			GetProductMock.When(ctx, 100).Then(&product1, nil).
			GetProductMock.When(ctx, 200).Then(&product2, nil)

		cartService := NewService(cartRepoMock, productServiceMock)

		cart, err := cartService.GetCart(ctx, userID)

		require.NoError(t, err)
		require.NotNil(t, cart)
		require.Len(t, cart.Items, 2)
		require.Equal(t, uint32(item1.Count)*product1.Price+uint32(item2.Count)*product2.Price, cart.TotalPrice)

		require.Equal(t, item1, cart.Items[0].Item)
		require.Equal(t, product1, cart.Items[0].Product)

		require.Equal(t, item2, cart.Items[1].Item)
		require.Equal(t, product2, cart.Items[1].Product)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		userID := model.UserID(1)

		cartRepoMock := mock.NewCartRepositoryMock(ctrl).GetCartMock.Return(nil, repository.ErrUserNotFound)

		cartService := NewService(cartRepoMock, nil)

		cart, err := cartService.GetCart(ctx, userID)

		require.ErrorIs(t, err, service.ErrEmptyCart)
		require.Nil(t, cart)
	})

	t.Run("ZeroCartLen", func(t *testing.T) {
		userID := model.UserID(1)

		cartRepoMock := mock.NewCartRepositoryMock(ctrl).GetCartMock.Return(make([]model.Item, 0), nil)

		cartService := NewService(cartRepoMock, nil)

		cart, err := cartService.GetCart(ctx, userID)

		require.ErrorIs(t, err, service.ErrEmptyCart)
		require.Empty(t, cart)
	})

	t.Run("ProductNotFound", func(t *testing.T) {
		userID := model.UserID(1)

		cartRepoMock := mock.NewCartRepositoryMock(ctrl).GetCartMock.Return([]model.Item{{SKU: 100, Count: 2}}, nil)

		productServiceMock := mock.NewProductServiceMock(ctrl).GetProductMock.Return(nil, product_app.ErrProductNotFound)

		cartService := NewService(cartRepoMock, productServiceMock)

		cart, err := cartService.GetCart(ctx, userID)

		require.Error(t, err)
		require.Nil(t, cart)
		require.ErrorIs(t, err, service.ErrProductNotFound)
	})
}
