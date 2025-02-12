package cart_repository

import (
	"context"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"testing"
)

func BenchmarkDeleteCart(b *testing.B) {
	ctx := context.Background()
	repo := NewRepository()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		userID := model.UserID(1)
		items := []model.Item{
			{SKU: 123, Count: 2},
			{SKU: 456, Count: 1},
		}
		repo.storage[userID] = items
		b.StartTimer()
		err := repo.DeleteCart(ctx, userID)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
