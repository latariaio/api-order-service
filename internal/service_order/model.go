package service_order

import (
	"time"

	"github.com/google/uuid"
	"github.com/latariaio/api-order-service/internal/customer"
	"github.com/shopspring/decimal"
)

type ServiceOrder struct {
	ID              uuid.UUID         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OrderNumber     int64             `gorm:"column:order_number;uniqueIndex;not null"`
	CustomerID      uuid.UUID         `gorm:"column:customer_id;type:uuid;not null"`
	Customer        customer.Customer `gorm:"foreignKey:CustomerID"`
	ReportedProblem string            `gorm:"column:reported_problem"`
	Diagnosis       string            `gorm:"column:diagnosis"`
	Notes           string            `gorm:"column:notes"`
	Status          Status            `gorm:"column:status;type:varchar(30);not null"`
	TotalPrice      decimal.Decimal   `gorm:"column:total_price;type:numeric(10,2);not null;default:0"`
	OpenedAt        time.Time         `gorm:"column:opened_at;not null"`
	CompletedAt     *time.Time        `gorm:"column:completed_at"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
