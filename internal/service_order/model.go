package service_order

import (
	"time"

	"github.com/google/uuid"
	"github.com/latariaio/api-order-service/internal/customer"
	"gorm.io/gorm"
)

type ServiceOrder struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Number          string
	CustomerID      string
	Customer        customer.Customer
	Status          string
	ReportedProblem string
	Diagnosis       string
	Notes           string
	OpenedAt        time.Time
	CompletedAt     time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
