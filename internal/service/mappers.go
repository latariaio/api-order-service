package service

func (r CreateServiceRequest) ToModel() Service {
	return Service{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
	}
}

func (r UpdateServiceRequest) ApplyTo(s *Service) {
	if r.Name != nil {
		s.Name = *r.Name
	}
	if r.Description != nil {
		s.Description = *r.Description
	}
	if r.Price != nil {
		s.Price = *r.Price
	}
}

func ToServiceResponse(s Service) ServiceResponse {
	return ServiceResponse{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		Price:       s.Price,
		CreatedAt:   s.CreatedAt,
	}
}

func ToServiceResponseList(services []Service) []ServiceResponse {
	responses := make([]ServiceResponse, len(services))
	for i, s := range services {
		responses[i] = ToServiceResponse(s)
	}
	return responses
}
