//go:build integration

package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/http-server/handler/cart"
	"github.com/stretchr/testify/require"
	"net/http"
)

func (cs *CartSuite) TestGetCart() {
	ctx := context.Background()

	userID := model.UserID(123)

	item1 := model.Item{SKU: 1076963, Count: 2}
	item2 := model.Item{SKU: 773297411, Count: 1}

	product1, err := cs.productService.GetProduct(ctx, uint32(item1.SKU))
	require.NoError(cs.T(), err)
	product2, err := cs.productService.GetProduct(ctx, uint32(item2.SKU))
	require.NoError(cs.T(), err)

	err = cs.cartRepo.AddItem(ctx, userID, item1)
	require.NoError(cs.T(), err)

	err = cs.cartRepo.AddItem(ctx, userID, item2)
	require.NoError(cs.T(), err)

	res, err := cs.server.Client().Get(fmt.Sprintf("%s/user/%d/cart", cs.server.URL, userID))
	cs.Require().NoError(err)

	defer res.Body.Close()

	cs.Require().Equal(http.StatusOK, res.StatusCode)

	var body cart_handler.GetCartResponse
	err = json.NewDecoder(res.Body).Decode(&body)
	cs.Require().NoError(err)

	cs.Assert().ElementsMatch(body.Items, []model.CartItem{
		{Item: item1, Product: *product1},
		{Item: item2, Product: *product2},
	})
	cs.Assert().Equal(body.TotalPrice, product1.Price*uint32(item1.Count)+product2.Price*uint32(item2.Count))
}
