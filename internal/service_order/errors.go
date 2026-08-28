package service_order

import "errors"

var (
	ErrServiceOrderNotFound     = errors.New("service order not found")
	ErrCustomerNotFound         = errors.New("customer not found")
	ErrInvalidStatusTransition  = errors.New("invalid status transition")
	ErrServiceOrderClosed       = errors.New("service order is closed and cannot be modified")
	ErrServiceOrderCannotDelete = errors.New("service order cannot be deleted in its current status")
)
