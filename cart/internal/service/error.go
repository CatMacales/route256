package service

import "errors"

var (
	ErrEmptyCart       = errors.New("cart is empty")
	ErrProductNotFound = errors.New("product not found")
)
