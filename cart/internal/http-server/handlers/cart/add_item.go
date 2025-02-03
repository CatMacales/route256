package cart_http

import (
	"encoding/json"
	"errors"
	"github.com/CatMacales/route256/cart/internal/app/product"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/http-server"
	"io"
	"net/http"
	"strconv"
)

const ADD_ITEM = "POST /user/<user_id>/cart/<sku_id>"

type AddItemRequest struct {
	// path value
	UserID int64 `json:"user_id"`
	SKU    int64 `json:"sku"`

	// body value
	Count uint16 `json:"count"`
}

func (s *Server) AddItem(w http.ResponseWriter, r *http.Request) {
	rawUserID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawUserID, 10, 64)
	if err != nil {
		http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusBadRequest)
		return
	}

	rawSKU := r.PathValue("sku_id")
	sku, err := strconv.ParseInt(rawSKU, 10, 64)
	if err != nil {
		http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusInternalServerError)
		return
	}

	addItemRequest := AddItemRequest{UserID: userID, SKU: sku}

	err = json.Unmarshal(body, &addItemRequest)
	if err != nil {
		http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusBadRequest)
		return
	}

	// TODO: add validation

	inputItem := model.Item{
		SKU:   addItemRequest.SKU,
		Count: addItemRequest.Count,
	}

	err = s.cartService.AddItem(r.Context(), addItemRequest.UserID, inputItem)
	if err != nil {
		if errors.Is(err, product_app.ErrProductNotFound) {
			http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusPreconditionFailed)
			return
		}
		http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusInternalServerError)
	}
}
