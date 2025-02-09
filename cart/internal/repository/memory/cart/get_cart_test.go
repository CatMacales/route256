package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetCartMemoryRepo(t *testing.T) {
	ctx := context.Background()
	repo := NewRepository()

	t.Run("ExistUser", func(t *testing.T) {
		userID := model.UserID(1)
		expectedItems := []model.Item{
			{SKU: 1, Count: 2},
			{SKU: 2, Count: 3},
		}

		repo.storage[userID] = expectedItems

		items, err := repo.GetCart(ctx, userID)

		require.NoError(t, err)
		require.Equal(t, expectedItems, items)
	})

	t.Run("NonExistUser", func(t *testing.T) {
		userID := model.UserID(2)

		items, err := repo.GetCart(context.Background(), userID)

		require.ErrorIs(t, err, repository.ErrUserNotFound)
		require.Nil(t, items)
	})
}
