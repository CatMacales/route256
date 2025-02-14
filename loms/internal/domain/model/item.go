package model

type Sku = uint32

type Item struct {
	SKU   Sku    `json:"sku_id"`
	Count uint16 `json:"count"`
}
