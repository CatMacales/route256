package model

type Cart struct {
	Items      []CartItem `json:"items"`
	TotalPrice uint32     `json:"total_price"`
}
