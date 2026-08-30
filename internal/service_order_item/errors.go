package service_order_item

import "errors"

var (
	ErrServiceOrderItemNotFound = errors.New("service order item not found")
	ErrServiceOrderNotFound     = errors.New("service order not found")
	ErrServiceNotFound          = errors.New("service not found")
	ErrServiceOrderClosed       = errors.New("service order is closed and cannot be modified")
	ErrItemDoesNotBelongToOrder = errors.New("item does not belong to this service order")
)
