package product_app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (a *App) GetProduct(_ context.Context, sku uint32) (*model.Product, error) {
	rawRequest := getProductRequest{
		Token: a.token,
		SKU:   sku,
	}
	request, err := json.Marshal(rawRequest)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.Post(
		a.url+"/get_product",
		"application/json",
		bytes.NewBuffer(request),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, ErrProductNotFound
	case http.StatusUnauthorized:
		return nil, ErrInvalidToken
	case http.StatusOK:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var productResp getProductResponse
		err = json.Unmarshal(body, &productResp)
		if err != nil {
			return nil, err
		}

		return &model.Product{
			Name:  productResp.Name,
			Price: productResp.Price,
		}, nil
	default:
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}
