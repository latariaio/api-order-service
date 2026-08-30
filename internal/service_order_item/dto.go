package service_order_item

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type AddServiceOrderItemRequest struct {
	ServiceID uuid.UUID `json:"serviceId" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,gt=0"`
}

// UpdateServiceOrderItemRequest — só quantidade é editável.
// Trocar de serviço deve ser DELETE + POST de um item novo.
type UpdateServiceOrderItemRequest struct {
	Quantity int `json:"quantity" binding:"required,gt=0"`
}

type ServiceOrderItemResponse struct {
	ID             uuid.UUID       `json:"id"`
	ServiceOrderID uuid.UUID       `json:"serviceOrderId"`
	ServiceID      uuid.UUID       `json:"serviceId"`
	Quantity       int             `json:"quantity"`
	UnitPrice      decimal.Decimal `json:"unitPrice"`
	TotalPrice     decimal.Decimal `json:"totalPrice"`
	CreatedAt      time.Time       `json:"createdAt"`
}
