//go:build integration

package integration

import (
	"context"
	"fmt"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/stretchr/testify/require"
	"net/http"
)

func (cs *CartSuite) TestDeleteItem() {
	ctx := context.Background()

	userID := model.UserID(123)
	item1 := model.Item{SKU: 1076963, Count: 2}
	item2 := model.Item{SKU: 773297411, Count: 1}

	err := cs.cartRepo.AddItem(ctx, userID, item1)
	require.NoError(cs.T(), err)

	err = cs.cartRepo.AddItem(ctx, userID, item2)
	require.NoError(cs.T(), err)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/user/%d/cart/%d", cs.server.URL, userID, item1.SKU), nil)
	cs.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")

	res, err := cs.server.Client().Do(req)
	cs.Require().NoError(err)

	defer res.Body.Close()

	cs.Require().Equal(http.StatusNoContent, res.StatusCode)

	cart, err := cs.cartRepo.GetCart(context.Background(), userID)
	cs.Require().NoError(err)

	cs.Require().ElementsMatch(cart, []model.Item{item2})
}
