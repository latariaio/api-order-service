package customer

// --- Create → Model ---

func (r CreateCustomerRequest) ToModel() Customer {
	return Customer{
		Name:         r.Name,
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

// --- Update → Model (aplica só os campos enviados) ---

func (r UpdateCustomerRequest) ApplyTo(c *Customer) {
	if r.Name != nil {
		c.Name = *r.Name
	}
	if r.Phone != nil {
		c.Phone = *r.Phone
	}
	if r.Email != nil {
		c.Email = *r.Email
	}
	if r.Address != nil {
		if r.Address.Street != nil {
			c.Address.Street = *r.Address.Street
		}
		if r.Address.Number != nil {
			c.Address.Number = *r.Address.Number
		}
		if r.Address.Complement != nil {
			c.Address.Complement = *r.Address.Complement
		}
		if r.Address.Neighborhood != nil {
			c.Address.Neighborhood = *r.Address.Neighborhood
		}
		if r.Address.City != nil {
			c.Address.City = *r.Address.City
		}
		if r.Address.State != nil {
			c.Address.State = *r.Address.State
		}
		if r.Address.ZipCode != nil {
			c.Address.ZipCode = *r.Address.ZipCode
		}
	}
}

// --- Model → Response ---

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
