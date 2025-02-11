package cart_handler

import (
	"encoding/json"
	"errors"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/http-server"
	"github.com/CatMacales/route256/cart/internal/lib/validation"
	"github.com/CatMacales/route256/cart/internal/service"
	"net/http"
)

const ADD_ITEM = "POST /user/<user_id>/cart/<sku_id>"

type AddItemRequest struct {
	// path value
	UserID int64 `json:"user_id" validate:"required,gte=0"`
	SKU    int64 `json:"sku" validate:"required,gte=0"`

	// body value
	Count uint16 `json:"count" validate:"required,gte=0"`
}

func (s *Handler) AddItem(w http.ResponseWriter, r *http.Request) {
	userID, err := parseIntPathValue(r, "user_id")
	if err != nil {
		http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusBadRequest)
		return
	}

	sku, err := parseIntPathValue(r, "sku_id")
	if err != nil {
		http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusBadRequest)
		return
	}

	addItemRequest := AddItemRequest{UserID: userID, SKU: sku}

	err = json.NewDecoder(r.Body).Decode(&addItemRequest)
	if err != nil {
		http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = validation.BeautyStructValidate(addItemRequest)
	if err != nil {
		http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusBadRequest)
		return
	}

	inputItem := model.Item{
		SKU:   addItemRequest.SKU,
		Count: addItemRequest.Count,
	}

	err = s.cartService.AddItem(r.Context(), addItemRequest.UserID, inputItem)
	if err != nil {
		if errors.Is(err, service.ErrEmptyCart) {
			http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusPreconditionFailed)
			return
		}
		http_server.GetErrorResponse(w, ADD_ITEM, err, http.StatusInternalServerError)
	}
}
