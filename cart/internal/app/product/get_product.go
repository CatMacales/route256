package product_app

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"io"
	"net/http"
)

type getProductRequest struct {
	Token string `json:"token"`
	SKU   uint32 `json:"sku"`
}

type getProductResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

// GetProduct retrieves a product by its SKU from a remote service.
// It sends a POST request to the "/get_product" endpoint with the SKU and token.
// Returns a Product model containing the product's name and price, or an error if the request fails.
func (a *App) GetProduct(_ context.Context, sku uint32) (*model.Product, error) {
	const getProductEndpoint = "/get_product"

	rawRequest := getProductRequest{
		Token: a.token,
		SKU:   sku,
	}
	request, err := json.Marshal(rawRequest)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.Post(
		a.url+getProductEndpoint,
		"application/json",
		bytes.NewBuffer(request),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handleHTTPError(resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var productResp getProductResponse
	if err := json.Unmarshal(body, &productResp); err != nil {
		return nil, err
	}

	return &model.Product{
		Name:  productResp.Name,
		Price: productResp.Price,
	}, nil
}
