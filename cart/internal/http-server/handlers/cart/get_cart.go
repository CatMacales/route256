package cart_http

import (
	"encoding/json"
	"errors"
	"github.com/CatMacales/route256/cart/internal/domain/model"
	"github.com/CatMacales/route256/cart/internal/http-server"
	"github.com/CatMacales/route256/cart/internal/lib/validation"
	"github.com/CatMacales/route256/cart/internal/service"
	"net/http"
	"strconv"
)

const GET_CART = "GET /user/<user_id>/cart"

type GetCartRequest struct {
	// path value
	UserID int64 `json:"user_id" validate:"required,gte=0"`
}

type GetCartResponse struct {
	Items      []model.CartItem `json:"items,omitempty"`
	TotalPrice uint32           `json:"total_price,omitempty"`
}

func (s *Server) GetCart(w http.ResponseWriter, r *http.Request) {
	rawUserID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawUserID, 10, 64)
	if err != nil {
		http_server.GetErrorResponse(w, GET_CART, err, http.StatusBadRequest)
		return
	}

	getCartRequest := GetCartRequest{UserID: userID}

	err = validation.BeautyStructValidate(getCartRequest)
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_ITEM, err, http.StatusBadRequest)
	}

	cart, err := s.cartService.GetCart(r.Context(), userID)
	if err != nil {
		if errors.Is(err, service.ErrEmptyCart) {
			http_server.GetErrorResponse(w, GET_CART, err, http.StatusNotFound)
			return
		}
		http_server.GetErrorResponse(w, GET_CART, err, http.StatusInternalServerError)
		return
	}

	getCartResponse := GetCartResponse{
		Items:      (*cart).Items,
		TotalPrice: (*cart).TotalPrice,
	}

	rawResponse, err := json.Marshal(getCartResponse)
	if err != nil {
		http_server.GetErrorResponse(w, GET_CART, err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(rawResponse)
}
