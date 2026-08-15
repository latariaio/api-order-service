package customer

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name     string    `gorm:"not null"`
	Document string    `gorm:"not null"`
	Phone    string    `gorm:"not null"`
	Email    string    `gorm:"not null"`
	Address  string    `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
