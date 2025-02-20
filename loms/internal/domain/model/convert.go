package model

import "github.com/CatMacales/route256/loms/pkg/api/loms/v1"

func ProtoToItem(item *loms.Item) Item {
	return Item{
		SKU:   item.Sku,
		Count: uint16(item.Count),
	}
}

func ItemToProto(item *Item) *loms.Item {
	return &loms.Item{
		Sku:   item.SKU,
		Count: uint32(item.Count),
	}
}

func ItemsToProto(items []Item) []*loms.Item {
	protoItems := make([]*loms.Item, 0, len(items))
	for _, item := range items {
		protoItems = append(protoItems, ItemToProto(&item))
	}
	return protoItems
}

func ProtoToItems(protoItems []*loms.Item) []Item {
	items := make([]Item, 0, len(protoItems))
	for _, protoItem := range protoItems {
		items = append(items, ProtoToItem(protoItem))
	}
	return items
}
