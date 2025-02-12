package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddItemMemoryRepo(t *testing.T) {
	ctx := context.Background()
	repo := NewRepository()

	userID := model.UserID(1)
	item := model.Item{
		SKU:   123,
		Count: 1,
	}

	// Add item to the user's cart
	err := repo.AddItem(ctx, userID, item)

	require.NoError(t, err)
	require.Len(t, repo.storage[userID], 1)
	require.Equal(t, item, repo.storage[userID][0])

	// Add the same item again with a different count to ensure the count is incremented
	item = model.Item{
		SKU:   123,
		Count: 2,
	}

	err = repo.AddItem(ctx, userID, item)
	require.NoError(t, err)
	require.Len(t, repo.storage[userID], 1)
	require.Equal(t, item.SKU, repo.storage[userID][0].SKU)
	require.Equal(t, uint16(3), repo.storage[userID][0].Count)

	// Add a different item to the user's cart
	item = model.Item{
		SKU:   456,
		Count: 2,
	}

	err = repo.AddItem(ctx, userID, item)
	require.NoError(t, err)
	require.Len(t, repo.storage[userID], 2)
	require.Equal(t, item, repo.storage[userID][1])
}
