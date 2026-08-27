package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateServiceRequest struct {
	Name        string          `json:"name" binding:"required"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price" binding:"required"`
}

type UpdateServiceRequest struct {
	Name        *string          `json:"name"`
	Description *string          `json:"description"`
	Price       *decimal.Decimal `json:"price"`
}

type ServiceResponse struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
	CreatedAt   time.Time       `json:"createdAt"`
}
