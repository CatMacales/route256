package repository

import "errors"

var (
	ErrNotEnoughStock = errors.New("not enough stock")
	ErrStockNotFound  = errors.New("stock not found")
	ErrOrderNotFound  = errors.New("order not found")
)
