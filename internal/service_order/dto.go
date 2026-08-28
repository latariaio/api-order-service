package service_order

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateServiceOrderRequest struct {
	CustomerID      uuid.UUID `json:"customerId" binding:"required"`
	ReportedProblem string    `json:"reportedProblem"`
}

// UpdateServiceOrderRequest — patch parcial, só reportedProblem/diagnosis/notes
// (customer_id e status nunca entram aqui, têm fluxos próprios)
type UpdateServiceOrderRequest struct {
	ReportedProblem *string `json:"reportedProblem"`
	Diagnosis       *string `json:"diagnosis"`
	Notes           *string `json:"notes"`
}

type UpdateServiceOrderStatusRequest struct {
	Status Status `json:"status" binding:"required"`
}

type ServiceOrderResponse struct {
	ID              uuid.UUID       `json:"id"`
	OrderNumber     int64           `json:"orderNumber"`
	CustomerID      uuid.UUID       `json:"customerId"`
	ReportedProblem string          `json:"reportedProblem"`
	Diagnosis       string          `json:"diagnosis,omitempty"`
	Notes           string          `json:"notes,omitempty"`
	Status          Status          `json:"status"`
	TotalPrice      decimal.Decimal `json:"totalPrice"`
	OpenedAt        time.Time       `json:"openedAt"`
	CompletedAt     *time.Time      `json:"completedAt,omitempty"`
	CreatedAt       time.Time       `json:"createdAt"`
}
