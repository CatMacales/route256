package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteItemMemoryRepo(t *testing.T) {
	ctx := context.Background()
	repo := NewRepository()

	t.Run("ExistUserAndSKU", func(t *testing.T) {
		userID := model.UserID(1)
		sku := model.Sku(123)

		repo.storage[userID] = []model.Item{
			{SKU: sku, Count: 2},
			{SKU: 456, Count: 2},
		}

		err := repo.DeleteItem(ctx, userID, sku)
		require.NoError(t, err)
		require.Len(t, repo.storage[userID], 1)
		require.Equal(t, model.Sku(456), repo.storage[userID][0].SKU)
	})

	t.Run("NonExistUser", func(t *testing.T) {
		userID := model.UserID(999)
		sku := model.Sku(123)

		err := repo.DeleteItem(ctx, userID, sku)
		require.NoError(t, err)
		require.Empty(t, repo.storage[userID])
	})

	t.Run("NonExistSKU", func(t *testing.T) {
		userID := model.UserID(1)
		sku := model.Sku(123)

		repo.storage[userID] = []model.Item{
			{SKU: 111, Count: 2},
			{SKU: 456, Count: 2},
		}

		err := repo.DeleteItem(ctx, userID, sku)
		require.NoError(t, err)
		require.Len(t, repo.storage[userID], 2)
		require.Equal(t, model.Sku(111), repo.storage[userID][0].SKU)
		require.Equal(t, model.Sku(456), repo.storage[userID][1].SKU)
	})

	t.Run("EmptyCart", func(t *testing.T) {
		userID := model.UserID(1)
		sku := model.Sku(123)

		repo.storage[userID] = []model.Item{}

		err := repo.DeleteItem(ctx, userID, sku)
		require.NoError(t, err)
		require.Empty(t, repo.storage[userID])
	})
}
