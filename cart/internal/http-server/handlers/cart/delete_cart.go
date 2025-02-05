package cart_http

import (
	"github.com/CatMacales/route256/cart/internal/http-server"
	"github.com/CatMacales/route256/cart/internal/lib/validation"
	"net/http"
)

const DELETE_CART = "DELETE /user/<user_id>/cart"

type DeleteCartRequest struct {
	// path value
	UserID int64 `json:"user_id" validate:"required,gte=0"`
}

func (s *Server) DeleteCart(w http.ResponseWriter, r *http.Request) {
	userID, err := parseIntPathValue(r, "user_id")
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_CART, err, http.StatusBadRequest)
		return
	}

	deleteCartRequest := DeleteCartRequest{UserID: userID}

	err = validation.BeautyStructValidate(deleteCartRequest)
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_CART, err, http.StatusBadRequest)
		return
	}

	err = s.cartService.DeleteCart(r.Context(), userID)
	if err != nil {
		http_server.GetErrorResponse(w, DELETE_CART, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
