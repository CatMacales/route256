package model

import "github.com/google/uuid"

type OrderID = uuid.UUID

type OrderStatus uint8

const (
	StatusUnknown OrderStatus = iota
	StatusNew
	StatusAwaitingPayment
	StatusFailed
	StatusPayed
	StatusCancelled
)

var OrderStatuses = map[OrderStatus]string{
	StatusUnknown:         "unknown",
	StatusNew:             "new",
	StatusAwaitingPayment: "awaiting_payment",
	StatusFailed:          "failed",
	StatusPayed:           "payed",
	StatusCancelled:       "cancelled",
}

type Order struct {
	UserID UserID      `json:"user_id"`
	Items  []Item      `json:"items"`
	Status OrderStatus `json:"status"`
}

func (o OrderStatus) String() string {
	if s, ok := OrderStatuses[o]; ok {
		return s
	}
	return OrderStatuses[StatusUnknown]
}
