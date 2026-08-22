package customer

import (
	"time"

	"github.com/google/uuid"
)

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

func ToCustomerResponse(c Customer) CustomerResponse {
	return CustomerResponse{
		ID:           c.ID,
		Name:         c.Name,
		Document:     c.Document,
		DocumentType: c.DocumentType,
		Phone:        c.Phone,
		Email:        c.Email,
		Address: AddressResponse{
			Street:       c.Address.Street,
			Number:       c.Address.Number,
			Complement:   c.Address.Complement,
			Neighborhood: c.Address.Neighborhood,
			City:         c.Address.City,
			State:        c.Address.State,
			ZipCode:      c.Address.ZipCode,
		},
		CreatedAt: c.CreatedAt,
	}
}

func ToCustomerResponseList(customers []Customer) []CustomerResponse {
	responses := make([]CustomerResponse, len(customers))
	for i, c := range customers {
		responses[i] = ToCustomerResponse(c)
	}
	return responses
}

func ToModel(r CustomerResponse) Customer {
	return Customer{Name: r.Name,
		Document:     r.Document,
		DocumentType: r.DocumentType,
		Phone:        r.Phone,
		Email:        r.Email,
		Address: Address{
			Street:       r.Address.Street,
			Number:       r.Address.Number,
			Complement:   r.Address.Complement,
			Neighborhood: r.Address.Neighborhood,
			City:         r.Address.City,
			State:        r.Address.State,
			ZipCode:      r.Address.ZipCode,
		},
	}
}
