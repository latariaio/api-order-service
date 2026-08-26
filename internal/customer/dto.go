package customer

import (
	"time"

	"github.com/google/uuid"
)

// --- Entrada: Create ---

type AddressRequest struct {
	Street       string `json:"street" binding:"required"`
	Number       string `json:"number" binding:"required"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	ZipCode      string `json:"zipCode" binding:"required"`
}

type CreateCustomerRequest struct {
	Name         string         `json:"name" binding:"required"`
	Document     string         `json:"document" binding:"required"`
	DocumentType CustomerType   `json:"documentType" binding:"required"`
	Phone        string         `json:"phone" binding:"required"`
	Email        string         `json:"email" binding:"required,email"`
	Address      AddressRequest `json:"address" binding:"required"`
}

// --- Entrada: Update (parcial) ---

type UpdateAddressRequest struct {
	Street       *string `json:"street"`
	Number       *string `json:"number"`
	Complement   *string `json:"complement"`
	Neighborhood *string `json:"neighborhood"`
	City         *string `json:"city"`
	State        *string `json:"state"`
	ZipCode      *string `json:"zipCode"`
}

type UpdateCustomerRequest struct {
	Name    *string               `json:"name"`
	Phone   *string               `json:"phone"`
	Email   *string               `json:"email"`
	Address *UpdateAddressRequest `json:"address"`
}

// --- Saída: Response ---

type AddressResponse struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipCode"`
}

type CustomerResponse struct {
	ID           uuid.UUID       `json:"id"`
	Name         string          `json:"name"`
	Document     string          `json:"document"`
	DocumentType CustomerType    `json:"documentType"`
	Phone        string          `json:"phone"`
	Email        string          `json:"email"`
	Address      AddressResponse `json:"address"`
	CreatedAt    time.Time       `json:"createdAt"`
}
