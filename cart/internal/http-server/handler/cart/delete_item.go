package cart_handler

import (
	"github.com/CatMacales/route256/cart/internal/http-server"
	"github.com/CatMacales/route256/cart/internal/lib/validation"
	"net/http"
)

const DELETE_ITEM = "DELETE /user/<user_id>/cart/<sku_id>"

type DeleteItemRequest struct {
	// path value
	UserID int64 `json:"user_id" validate:"required,gte=0"`
	SKU    int64 `json:"sku" validate:"required,gte=0"`
}

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	userID, err := parseIntPathValue(r, "user_id")
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_ITEM, err, http.StatusBadRequest)
		return
	}

	sku, err := parseIntPathValue(r, "sku_id")
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_ITEM, err, http.StatusBadRequest)
		return
	}

	deleteItemRequest := DeleteItemRequest{UserID: userID, SKU: sku}

	err = validation.BeautyStructValidate(deleteItemRequest)
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_ITEM, err, http.StatusBadRequest)
		return
	}

	err = h.cartService.DeleteItem(r.Context(), userID, sku)
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_ITEM, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
