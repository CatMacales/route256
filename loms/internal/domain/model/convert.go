package model

import "github.com/CatMacales/route256/loms/pkg/api/loms/v1"

func ProtoToOrder(createOrderRequest *loms.CreateOrderRequest) Order {
	items := make([]Item, 0, len(createOrderRequest.Items))
	for _, item := range createOrderRequest.Items {
		items = append(items, Item{SKU: item.Sku, Count: uint16(item.Count)})
	}
	return Order{
		UserID: createOrderRequest.UserId,
		Items:  items,
		Status: StatusUnknown,
	}
}

func OrderToProto(order *Order) *loms.GetOrderInfoResponse {
	orderInfo := &loms.GetOrderInfoResponse{
		UserId: order.UserID,
		Status: loms.OrderStatus(order.Status),
		Items:  make([]*loms.Item, 0, len(order.Items)),
	}
	for _, item := range order.Items {
		orderInfo.Items = append(orderInfo.Items, &loms.Item{Sku: item.SKU, Count: uint32(item.Count)})
	}
	return orderInfo
}
