package service_order_item

import (
	"time"

	"github.com/google/uuid"
	"github.com/latariaio/api-order-service/internal/service"
	"github.com/latariaio/api-order-service/internal/service_order"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ServiceOrderItem struct {
	ID             uuid.UUID                  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ServiceOrderID uuid.UUID                  `gorm:"column:service_order_id;type:uuid;not null"`
	ServiceOrder   service_order.ServiceOrder `gorm:"foreignKey:ServiceOrderID"`
	ServiceID      uuid.UUID                  `gorm:"column:service_id;type:uuid;not null"`
	Service        service.Service            `gorm:"foreignKey:ServiceID"`
	Quantity       int                        `gorm:"column:quantity;not null"`
	UnitPrice      decimal.Decimal            `gorm:"column:unit_price;type:numeric(10,2);not null"`
	TotalPrice     decimal.Decimal            `gorm:"column:total_price;type:numeric(10,2);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
