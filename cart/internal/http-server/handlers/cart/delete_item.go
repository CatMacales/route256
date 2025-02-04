package cart_http

import (
	"github.com/CatMacales/route256/cart/internal/http-server"
	"github.com/CatMacales/route256/cart/internal/lib/validation"
	"net/http"
	"strconv"
)

const DELETE_ITEM = "DELETE /user/<user_id>/cart/<sku_id>"

type DeleteItemRequest struct {
	// path value
	UserID int64 `json:"user_id" validate:"required,gte=0"`
	SKU    int64 `json:"sku" validate:"required,gte=0"`
}

func (s *Server) DeleteItem(w http.ResponseWriter, r *http.Request) {
	rawUserID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawUserID, 10, 64)
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_ITEM, err, http.StatusBadRequest)
		return
	}

	rawSKU := r.PathValue("sku_id")
	sku, err := strconv.ParseInt(rawSKU, 10, 64)
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_ITEM, err, http.StatusBadRequest)
		return
	}

	deleteItemRequest := DeleteItemRequest{UserID: userID, SKU: sku}

	err = validation.BeautyStructValidate(deleteItemRequest)
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_ITEM, err, http.StatusBadRequest)
	}

	err = s.cartService.DeleteItem(r.Context(), userID, sku)
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_ITEM, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
