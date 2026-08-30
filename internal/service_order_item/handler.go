package service_order_item

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/latariaio/api-order-service/internal/httputil"
)

type ServiceOrderItemHandler struct {
	service *ServiceOrderItemService
}

func NewServiceOrderItemHandler(service *ServiceOrderItemService) *ServiceOrderItemHandler {
	return &ServiceOrderItemHandler{service: service}
}

func (h *ServiceOrderItemHandler) AddItem(ctx *gin.Context) {
	orderID := ctx.Param("id")

	var request AddServiceOrderItemRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		if fieldErrors := httputil.FormatValidationError(err); fieldErrors != nil {
			ctx.JSON(400, gin.H{"errors": fieldErrors})
			return
		}
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	item, err := h.service.AddItem(orderID, request)
	if err != nil {
		switch {
		case errors.Is(err, ErrServiceOrderNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		case errors.Is(err, ErrServiceNotFound):
			ctx.JSON(400, gin.H{"error": err.Error()})
		case errors.Is(err, ErrServiceOrderClosed):
			ctx.JSON(409, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error adding item: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": ToServiceOrderItemResponse(*item)})
}

func (h *ServiceOrderItemHandler) ListItems(ctx *gin.Context) {
	orderID := ctx.Param("id")
	items, err := h.service.ListByOrder(orderID)
	if err != nil {
		log.Printf("unexpected error listing items: %v", err)
		ctx.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	ctx.JSON(200, gin.H{"data": ToServiceOrderItemResponseList(items)})
}

func (h *ServiceOrderItemHandler) UpdateItem(ctx *gin.Context) {
	orderID := ctx.Param("id")
	itemID := ctx.Param("itemId")

	var request UpdateServiceOrderItemRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		if fieldErrors := httputil.FormatValidationError(err); fieldErrors != nil {
			ctx.JSON(400, gin.H{"errors": fieldErrors})
			return
		}
		ctx.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	item, err := h.service.UpdateItem(orderID, itemID, request)
	if err != nil {
		switch {
		case errors.Is(err, ErrServiceOrderNotFound), errors.Is(err, ErrServiceOrderItemNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		case errors.Is(err, ErrItemDoesNotBelongToOrder):
			ctx.JSON(400, gin.H{"error": err.Error()})
		case errors.Is(err, ErrServiceOrderClosed):
			ctx.JSON(409, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error updating item: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.JSON(200, gin.H{"data": ToServiceOrderItemResponse(*item)})
}

func (h *ServiceOrderItemHandler) RemoveItem(ctx *gin.Context) {
	orderID := ctx.Param("id")
	itemID := ctx.Param("itemId")

	if err := h.service.RemoveItem(orderID, itemID); err != nil {
		switch {
		case errors.Is(err, ErrServiceOrderNotFound), errors.Is(err, ErrServiceOrderItemNotFound):
			ctx.JSON(404, gin.H{"error": err.Error()})
		case errors.Is(err, ErrItemDoesNotBelongToOrder):
			ctx.JSON(400, gin.H{"error": err.Error()})
		case errors.Is(err, ErrServiceOrderClosed):
			ctx.JSON(409, gin.H{"error": err.Error()})
		default:
			log.Printf("unexpected error removing item: %v", err)
			ctx.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.Status(204)
}
