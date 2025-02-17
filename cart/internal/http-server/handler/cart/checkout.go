package cart_handler

import (
	"encoding/json"
	"errors"
	"github.com/CatMacales/route256/cart/internal/http-server"
	"github.com/CatMacales/route256/cart/internal/lib/validation"
	"github.com/CatMacales/route256/cart/internal/repository"
	"github.com/google/uuid"
	"net/http"
)

const CHECKOUT = "POST /cart/checkout"

type CheckoutRequest struct {
	UserID int64 `json:"user_id" validate:"required,gte=0"`
}

type CheckoutResponse struct {
	OrderID uuid.UUID `json:"order_id"`
}

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	var checkoutRequest CheckoutRequest

	err := json.NewDecoder(r.Body).Decode(&checkoutRequest)
	if err != nil {
		http_server.GetErrorResponse(w, CHECKOUT, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = validation.BeautyStructValidate(checkoutRequest)
	if err != nil {
		http_server.GetErrorResponse(w, CHECKOUT, err, http.StatusBadRequest)
		return
	}

	orderID, err := h.cartService.Checkout(r.Context(), checkoutRequest.UserID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			http_server.GetErrorResponse(w, CHECKOUT, err, http.StatusNotFound)
			return
		}

		http_server.GetErrorResponse(w, CHECKOUT, err, http.StatusInternalServerError)
		return
	}

	checkoutResponse := CheckoutResponse{
		OrderID: orderID,
	}

	rawResponse, err := json.Marshal(checkoutResponse)
	if err != nil {
		http_server.GetErrorResponse(w, CHECKOUT, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(rawResponse); err != nil {
		http_server.GetErrorResponse(w, CHECKOUT, err, http.StatusInternalServerError)
		return
	}
}
