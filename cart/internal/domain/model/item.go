package model

type Sku = int64

type Item struct {
	SKU   Sku    `json:"sku_id"`
	Count uint16 `json:"count"`
}

type CartItem struct {
	Item    Item
	Product Product
}
