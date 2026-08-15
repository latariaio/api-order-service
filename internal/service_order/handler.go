package service_order

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceOrderHandler struct {
	service *ServiceOrderService
}

func NewServiceOrderHandler(service *ServiceOrderService) *ServiceOrderHandler {
	return &ServiceOrderHandler{service: service}
}

type CreateServiceOrderRequest struct {
	Number          string    `json:"number"`
	CustomerID      string    `json:"customer_id"`
	Status          string    `json:"status"`
	ReportedProblem string    `json:"reported_problem"`
	Diagnosis       string    `json:"diagnosis"`
	Notes           string    `json:"notes"`
	OpenedAt        time.Time `json:"opened_at"`
	CompletedAt     time.Time `json:"completed_at"`
}

type UpdateServiceOrderRequest struct {
	ID              string    `json:"id"`
	Number          string    `json:"number"`
	CustomerID      string    `json:"customer_id"`
	Status          string    `json:"status"`
	ReportedProblem string    `json:"reported_problem"`
	Diagnosis       string    `json:"diagnosis"`
	Notes           string    `json:"notes"`
	OpenedAt        time.Time `json:"opened_at"`
	CompletedAt     time.Time `json:"completed_at"`
}

func (h *ServiceOrderHandler) Create(ctx *gin.Context) {
	var request CreateServiceOrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order := &ServiceOrder{
		Number:          request.Number,
		CustomerID:      request.CustomerID,
		Status:          request.Status,
		ReportedProblem: request.ReportedProblem,
		Diagnosis:       request.Diagnosis,
		Notes:           request.Notes,
		OpenedAt:        request.OpenedAt,
		CompletedAt:     request.CompletedAt,
	}

	if err := h.service.Create(order); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, order)
}

func (h *ServiceOrderHandler) GetAll(ctx *gin.Context) {
	orders, err := h.service.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, orders)
}

func (h *ServiceOrderHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := h.service.FindByID(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, order)
}

func (h *ServiceOrderHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var request UpdateServiceOrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	idUp, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id: " + err.Error()})
		return
	}

	service_order := ServiceOrder{
		ID:              idUp,
		Number:          request.Number,
		CustomerID:      request.CustomerID,
		Status:          request.Status,
		ReportedProblem: request.ReportedProblem,
		Diagnosis:       request.Diagnosis,
		Notes:           request.Notes,
		OpenedAt:        request.OpenedAt,
		CompletedAt:     request.CompletedAt,
	}

	order := h.service.Update(id, &service_order)
	if order != nil {
		ctx.JSON(500, gin.H{"error": order.Error()})
		return
	}

	ctx.JSON(200, order)
}

func (h *ServiceOrderHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.service.Delete(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "deleted"})
}
