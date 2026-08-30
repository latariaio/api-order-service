package service_order_item

func ToServiceOrderItemResponse(i ServiceOrderItem) ServiceOrderItemResponse {
	return ServiceOrderItemResponse{
		ID:             i.ID,
		ServiceOrderID: i.ServiceOrderID,
		ServiceID:      i.ServiceID,
		Quantity:       i.Quantity,
		UnitPrice:      i.UnitPrice,
		TotalPrice:     i.TotalPrice,
		CreatedAt:      i.CreatedAt,
	}
}

func ToServiceOrderItemResponseList(items []ServiceOrderItem) []ServiceOrderItemResponse {
	responses := make([]ServiceOrderItemResponse, len(items))
	for i, item := range items {
		responses[i] = ToServiceOrderItemResponse(item)
	}
	return responses
}
