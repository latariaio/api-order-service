package service_order

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/latariaio/api-order-service/internal/httputil"
)

type ServiceOrderHandler struct {
	service *ServiceOrderService
}

func NewServiceOrderHandler(service *ServiceOrderService) *ServiceOrderHandler {
	return &ServiceOrderHandler{service: service}
}

func (h *ServiceOrderHandler) Create(ctx *gin.Context) {
	var request CreateServiceOrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		if fieldErrors := httputil.FormatValidationError(err); fieldErrors != nil {
			ctx.JSON(400, gin.H{"errors": fieldErrors})
			return
		}
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	order := request.ToModel()

	if err := h.service.Create(&order); err != nil {
		switch {
		case errors.Is(err, ErrCustomerNotFound):
			ctx.JSON(400, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error creating service order: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": ToServiceOrderResponse(order)})
}

func (h *ServiceOrderHandler) GetAll(ctx *gin.Context) {
	orders, err := h.service.GetAll()
	if err != nil {
		log.Printf("unexpected error listing service orders: %v", err)
		ctx.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	ctx.JSON(200, gin.H{"data": ToServiceOrderResponseList(orders)})
}

func (h *ServiceOrderHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := h.service.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrServiceOrderNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error getting service order: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}
	ctx.JSON(200, gin.H{"data": ToServiceOrderResponse(*order)})
}

type UpdateReportedProblemRequest struct {
	ReportedProblem string `json:"reportedProblem" binding:"required"`
}

func (h *ServiceOrderHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var request UpdateServiceOrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		if fieldErrors := httputil.FormatValidationError(err); fieldErrors != nil {
			ctx.JSON(400, gin.H{"errors": fieldErrors})
			return
		}
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	order, err := h.service.Update(id, request)
	if err != nil {
		switch {
		case errors.Is(err, ErrServiceOrderNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		case errors.Is(err, ErrServiceOrderClosed):
			ctx.JSON(409, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error updating service order: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.JSON(200, gin.H{"data": ToServiceOrderResponse(*order)})
}
func (h *ServiceOrderHandler) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")

	var request UpdateServiceOrderStatusRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		if fieldErrors := httputil.FormatValidationError(err); fieldErrors != nil {
			ctx.JSON(400, gin.H{"errors": fieldErrors})
			return
		}
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	order, err := h.service.UpdateStatus(id, request.Status)
	if err != nil {
		switch {
		case errors.Is(err, ErrServiceOrderNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		case errors.Is(err, ErrInvalidStatusTransition):
			ctx.JSON(400, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error updating service order status: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.JSON(200, gin.H{"data": ToServiceOrderResponse(*order)})
}

func (h *ServiceOrderHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	order, err := h.service.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrServiceOrderNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		case errors.Is(err, ErrServiceOrderCannotDelete):
			ctx.JSON(409, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error deleting service order: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	if order != nil {
		// não foi excluída, foi cancelada — devolve o estado atual
		ctx.JSON(200, gin.H{"data": ToServiceOrderResponse(*order)})
		return
	}

	ctx.Status(204)
}
