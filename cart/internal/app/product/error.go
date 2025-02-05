package product_app

import "fmt"

var (
	ErrProductNotFound = fmt.Errorf("product not found")
	ErrInvalidToken    = fmt.Errorf("token is invalid")
)
