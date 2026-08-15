package service_order_item

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceOrderItemHandler struct {
	service *ServiceOrderItemService
}

func NewServiceOrderItemHandler(service *ServiceOrderItemService) *ServiceOrderItemHandler {
	return &ServiceOrderItemHandler{
		service: service,
	}
}

type ServiceOrderItemInput struct {
	ServiceOrderID string
	ServiceID      string
	Quantity       int
	UnitPrice      float64
	TotalPrice     float64
}

func (h *ServiceOrderItemHandler) CreateServiceOrderItem(c *gin.Context) {
	var input ServiceOrderItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service order item created successfully"})
}
