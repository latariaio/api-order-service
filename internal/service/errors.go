package service

import "errors"

var (
	ErrServiceNotFound          = errors.New("service not found")
	ErrInvalidPrice             = errors.New("price must be greater than or equal to zero")
	ErrServiceNameAlreadyExists = errors.New("service name already exists")
	ErrServiceInUse             = errors.New("service is already used in an order and cannot be deleted")
)
