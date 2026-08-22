package service_order_item

import (
	"time"

	"github.com/google/uuid"
	"github.com/latariaio/api-order-service/internal/service"
	"github.com/latariaio/api-order-service/internal/service_order"
	"gorm.io/gorm"
)

type ServiceOrderItem struct {
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ServiceOrderID uuid.UUID
	ServiceOrder   service_order.ServiceOrder `gorm:"foreignKey:ServiceOrderID"`
	ServiceID      uuid.UUID
	Service        service.Service `gorm:"foreignKey:ServiceID"`
	Quantity       int
	UnitPrice      float64
	TotalPrice     float64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
