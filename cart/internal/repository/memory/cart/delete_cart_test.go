package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteCartMemoryRepo(t *testing.T) {
	ctx := context.Background()
	repo := NewRepository()

	t.Run("ExistUser", func(t *testing.T) {
		userID := model.UserID(1)
		items := []model.Item{
			{SKU: 123, Count: 2},
			{SKU: 456, Count: 1},
		}
		repo.storage[userID] = items

		err := repo.DeleteCart(ctx, userID)
		require.NoError(t, err)
		require.Empty(t, repo.storage[userID])

	})
	t.Run("NonExistUser", func(t *testing.T) {
		userID := model.UserID(2)

		err := repo.DeleteCart(ctx, userID)
		require.NoError(t, err)
		require.Contains(t, repo.storage, userID)
		require.Empty(t, repo.storage[userID])
	})
}
