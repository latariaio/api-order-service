package customer

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerType string

const (
	CustomerTypeIndividual CustomerType = "PF" // CPF
	CustomerTypeCompany    CustomerType = "PJ" // CNPJ
)

type Address struct {
	Street       string `json:"street" gorm:"column:street;not null"`
	Number       string `json:"number" gorm:"column:number;not null"`
	Complement   string `json:"complement,omitempty" gorm:"column:complement"`
	Neighborhood string `json:"neighborhood" gorm:"column:neighborhood;not null"`
	City         string `json:"city" gorm:"column:city;not null"`
	State        string `json:"state" gorm:"column:state;not null"`
	ZipCode      string `json:"zipCode" gorm:"column:zip_code;not null"`
}

type Customer struct {
	ID           uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name         string       `json:"name" gorm:"column:name;not null"`
	Document     string       `json:"document" gorm:"column:document;not null"` // unique enquanto deleted_at IS NULL — ver migration 000002
	DocumentType CustomerType `json:"documentType" gorm:"column:document_type;type:varchar(20);not null"`
	Phone        string       `json:"phone" gorm:"column:phone;not null"`
	Email        string       `json:"email" gorm:"column:email"` // unique enquanto deleted_at IS NULL — ver migration 000002
	Address      Address      `json:"address" gorm:"embedded;embeddedPrefix:address_"`
	CreatedAt    time.Time    `json:"createdAt" gorm:"column:created_at;autoCreateTime"`

	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"column:deleted_at;index"`
}
