package service_order_item

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceOrderItemHandler struct {
	service *ServiceOrderItemService
}

func NewServiceOrderItemHandler(service *ServiceOrderItemService) *ServiceOrderItemHandler {
	return &ServiceOrderItemHandler{
		service: service,
	}
}

type ServiceOrderItemRequest struct {
	ServiceOrderID string
	ServiceID      string
	Quantity       int
	UnitPrice      float64
	TotalPrice     float64
}

func (h *ServiceOrderItemHandler) CreateServiceOrderItem(c *gin.Context) {
	var request ServiceOrderItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serviceOrderID, err := uuid.Parse(request.ServiceOrderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service_order_id: " + err.Error()})
		return
	}

	serviceID, err := uuid.Parse(request.ServiceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service_id: " + err.Error()})
		return
	}

	item := &ServiceOrderItem{
		ServiceOrderID: serviceOrderID,
		ServiceID:      serviceID,
		Quantity:       request.Quantity,
		UnitPrice:      request.UnitPrice,
		TotalPrice:     request.TotalPrice,
	}

	if err := h.service.Create(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *ServiceOrderItemHandler) GetServiceOrderItems(c *gin.Context) {
	items, err := h.service.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *ServiceOrderItemHandler) GetServiceOrderItem(c *gin.Context) {
	item, err := h.service.repo.FindById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *ServiceOrderItemHandler) UpdateServiceOrderItem(c *gin.Context) {
	var request ServiceOrderItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.repo.FindById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	item.Quantity = request.Quantity
	item.UnitPrice = request.UnitPrice
	item.TotalPrice = request.TotalPrice

	if err := h.service.repo.Update(item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *ServiceOrderItemHandler) DeleteServiceOrderItem(c *gin.Context) {
	if err := h.service.repo.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
}
