package service_order

func (r CreateServiceOrderRequest) ToModel() ServiceOrder {
	return ServiceOrder{
		CustomerID:      r.CustomerID,
		ReportedProblem: r.ReportedProblem,
		Status:          StatusOpen,
	}
}

func (r UpdateServiceOrderRequest) ApplyTo(so *ServiceOrder) {
	if r.ReportedProblem != nil {
		so.ReportedProblem = *r.ReportedProblem
	}
	if r.Diagnosis != nil {
		so.Diagnosis = *r.Diagnosis
	}
	if r.Notes != nil {
		so.Notes = *r.Notes
	}
}

func ToServiceOrderResponse(so ServiceOrder) ServiceOrderResponse {
	return ServiceOrderResponse{
		ID:              so.ID,
		OrderNumber:     so.OrderNumber,
		CustomerID:      so.CustomerID,
		ReportedProblem: so.ReportedProblem,
		Diagnosis:       so.Diagnosis,
		Notes:           so.Notes,
		Status:          so.Status,
		TotalPrice:      so.TotalPrice,
		OpenedAt:        so.OpenedAt,
		CompletedAt:     so.CompletedAt,
		CreatedAt:       so.CreatedAt,
	}
}

func ToServiceOrderResponseList(list []ServiceOrder) []ServiceOrderResponse {
	responses := make([]ServiceOrderResponse, len(list))
	for i, so := range list {
		responses[i] = ToServiceOrderResponse(so)
	}
	return responses
}
